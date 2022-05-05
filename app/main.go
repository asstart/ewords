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
	FilePath         string
	DirPath          string
	NotionDB         string
	ExportExample    bool
	ExportDefenition bool
	ExampleDir       string
	DefenitionDir    string
	ExportToFile     bool
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
	{Flag: "sf", Field: "FilePath", Help: "path to a source file", Value: ""},
	{Flag: "sd", Field: "DirPath", Help: "path to a source directory", Value: ""},
	{Flag: "sn", Field: "NotionDB", Help: "id of notion database", Value: ""},
	{Flag: "f", Field: "ExportToFile", Help: "export to a file", Value: false},
	{Flag: "ee", Field: "ExportExample", Help: "export examples", Value: false},
	{Flag: "ed", Field: "ExportDefenition", Help: "export defenitions", Value: false},
	{Flag: "ted", Field: "ExampleDir", Help: "where store example output", Value: "ewords_example"},
	{Flag: "tdd", Field: "DefenitionDir", Help: "where store defenition output", Value: "ewords_defenition"},
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
		ts := parseFileCmd(op.FilePath)
		files[op.FilePath] = ts
	} else if op.DirPath != "" {
		lf := parseDirCmd(op.DirPath)
		for _, file := range lf {
			ts := parseFileCmd(file)
			files[file] = ts
		}
	} else if op.NotionDB != "" {
		ts := parseNotionCmd(op.NotionDB)
		if len(ts) > 0 {
			filename := fmt.Sprintf("notion_%v", time.Now().Format(time.Stamp))
			files[filename] = ts
		}

	}

	if op.ExportToFile {
		for file, ts := range files {
			if op.ExportExample {
				exportExamplesCmd(file, ts)
			}
			if op.ExportDefenition {
				exportDefenitionsCmd(file, ts)
			}
		}

	}

}

func exportExamplesCmd(file string, ts []ewords.TermSource) {
	filename := path.Base(file)
	pattern, _ := regexp.Compile("[.]")
	eout := pattern.ReplaceAllString(filename, "-examples.")
	exportExamples(ts, path.Join(op.ExampleDir, eout))
}

func exportDefenitionsCmd(file string, ts []ewords.TermSource) {
	filename := path.Base(file)
	pattern, _ := regexp.Compile("[.]")
	dout := pattern.ReplaceAllString(filename, "-defenitions.")
	exportDefenitions(ts, path.Join(op.DefenitionDir, dout))
}

func exportExamples(ts []ewords.TermSource, exmplFile string) {
	ep := ewords.ExampleFileExporter{
		FilePath:        exmplFile,
		Formatter:       &ewords.SimpleTsv{},
		OutputFormatter: &ewords.SimpleTsv{},
	}

	err := ep.ExportExample(ewords.ToExamples(&ewords.DefaultTerm{}, ts))
	if err != nil {
		panic(err)
	}
}

func exportDefenitions(ts []ewords.TermSource, defFile string) {
	dp := ewords.DefenitionFileExporter{
		FilePath:        defFile,
		Formatter:       &ewords.SimpleTsv{},
		OutputFormatter: &ewords.SimpleTsv{},
	}
	err := dp.ExportDefenition(ewords.ToDefenitions(&ewords.DefaultTerm{}, ts))
	if err != nil {
		panic(err)
	}
}

func parseFileCmd(file string) []ewords.TermSource {
	ts, err := parse.Source2TermSource(&file, &parse.TsvParser{})
	if err != nil {
		panic(fmt.Sprintf("Errow while parsing TSV source to TermSource- %v", err))

	}
	return ts
}

func parseNotionCmd(db string) []ewords.TermSource {
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

func parseDirCmd(dir string) []string {
	paths, err := ewords.ListFiles(&dir)
	if err != nil {
		panic(fmt.Sprintf("can't list files in: %v, %v", dir, err))
	}
	return paths
}
