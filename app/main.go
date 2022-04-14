package main

import (
	"fmt"
	"path"
	"reflect"
	"regexp"

	"github.com/asstart/english-words/app/ewords"
)

type options struct {
	FilePath      string
	DirPath       string
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

	if op.FilePath == "" && op.DirPath == "" {
		panic("neither s nor d were defined")
	}

	if op.FilePath != "" && op.DirPath != "" {
		panic("both s and d were defined")
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

	source, err := ewords.ReadFile(&file)
	if err != nil {
		panic(fmt.Sprintf("Error while reading %v - %v", file, err))
	}
	ts, err := ewords.Str2TermSource(&source, &ewords.TsvParser{}, &ewords.TsvParser{})
	if err != nil {
		panic(fmt.Sprintf("Errow while parsing source to TermSource- %v", err))

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
