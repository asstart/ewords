package main

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"regexp"
	"time"

	"github.com/asstart/english-words/app/ewords"
	"github.com/asstart/english-words/app/parse"
)

const APIKeyEnv = "NOTION_API_KEY"

type options struct {
	FilePath      string
	DirPath       string
	NotionDB      string
	ExampleDir    string
	DefenitionDir string
	M             bool
}

func (op *options) SetString(field string, value *string) {
	f := reflect.ValueOf(op).Elem().FieldByName(field)
	f.SetString(*value)
}

func (op *options) SetBool(field string, value *bool) {
	f := reflect.ValueOf(op).Elem().FieldByName(field)
	f.SetBool(*value)
}

func (op *options) SetInt(field string, value *int) {
}

var flags = []*ewords.FlagDef{
	{Flag: "s", Field: "FilePath", Help: "path to the source file", Value: ""},
	{Flag: "d", Field: "DirPath", Help: "path to the source directory", Value: ""},
	{Flag: "n", Field: "NotionDB", Help: "Id of notion database", Value: ""},
	{Flag: "m", Field: "M", Help: "export to memrise format", Value: false},
	{Flag: "ex", Field: "ExampleDir", Help: "where store example output", Value: "ewords_example"},
	{Flag: "df", Field: "DefenitionDir", Help: "where store defenition output", Value: "ewords_defenition"},
}

var op = options{}

func main() {
	err := ewords.ParseArgs(flags, &op)
	if err != nil {
		panic(err)
	}

	if op.FilePath == "" && op.DirPath == "" && op.NotionDB == "" {
		panic("neither s nor d nor n were defined")
	}

	if (op.FilePath != "" && op.DirPath != "") ||
		(op.FilePath != "" && op.NotionDB != "") ||
		(op.NotionDB != "" && op.DirPath != "") {
		panic("only one option of (s, d, n) should be defined")
	}

	files := map[string][]ewords.TermSource{}

	if op.FilePath != "" {
		ts := parseFile(op.FilePath)
		files[op.FilePath] = ts
	} else if op.DirPath != "" {
		lf := parseDir(op.DirPath)
		for _, file := range lf {
			ts := parseFile(file)
			files[file] = ts
		}
	} else if op.NotionDB != "" {
		ts := parseNotion(op.NotionDB)
		filename := fmt.Sprintf("notion_%v", time.Now().Format(time.Stamp))
		files[filename] = ts
	}

	if op.M {
		for file, ts := range files {
			publishMemriseCmd(file, ts)
		}

	}

}

func publishMemriseCmd(file string, ts []ewords.TermSource) {
	filename := path.Base(file)
	pattern, _ := regexp.Compile("[.]")
	eout := pattern.ReplaceAllString(filename, "-examples.")
	dout := pattern.ReplaceAllString(filename, "-defenitions.")
	publishMemrise(ts, path.Join(op.ExampleDir, eout), path.Join(op.DefenitionDir, dout))
}

func publishMemrise(ts []ewords.TermSource, exmplFile string, defFile string) {
	ep := ewords.MemriseExamplePublisher{
		FilePath:        exmplFile,
		Formatter:       &ewords.MemriseFormatter{},
		OutputFormatter: &ewords.MemriseFormatter{},
	}
	dp := ewords.MemriseDefenitionPublisher{
		FilePath:        defFile,
		Formatter:       &ewords.MemriseFormatter{},
		OutputFormatter: &ewords.MemriseFormatter{},
	}
	ep.PublishExample(ewords.ToExamples(&ewords.DefaultTerm{}, ts))
	dp.PublishDefenition(ewords.ToDefenitions(&ewords.DefaultTerm{}, ts))
}

func parseFile(file string) []ewords.TermSource {
	ts, err := parse.Source2TermSource(&file, &parse.TsvParser{})
	if err != nil {
		panic(fmt.Sprintf("Errow while parsing TSV source to TermSource- %v", err))

	}
	return ts
}

func parseNotion(db string) []ewords.TermSource {
	var key = os.Getenv(APIKeyEnv)
	ts, err := parse.Source2TermSource(
		&db,
		&parse.NotionParser{
			ApiKey: &key,
		},
	)
	if err != nil {
		panic(fmt.Sprintf("Errow while parsing Notion source to TermSource - %v", err))
	}
	return ts
}

func parseDir(dir string) []string {
	paths, err := ewords.ListFiles(&dir)
	if err != nil {
		panic(fmt.Sprintf("can't list files in: %v, %v", dir, err))
	}
	return paths

}
