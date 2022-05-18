package main

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/asstart/english-words/app/ewords"
	"github.com/asstart/english-words/app/parse"
	"github.com/asstart/english-words/app/utils"
	"github.com/spf13/afero"
)

const APIKeyEnv = "NOTION_API_KEY"

var AppFS = afero.NewOsFs()

type options struct {
	FilePath         string
	DirPath          string
	NotionDB         string
	ExportExample    bool
	ExportDefenition bool
	ExampleDir       string
	DefenitionDir    string
}

func (op *options) SetString(field string, value *string) {
	f := reflect.ValueOf(op).Elem().FieldByName(field)
	f.SetString(strings.TrimSpace(*value))
}

func (op *options) SetBool(field string, value *bool) {
	f := reflect.ValueOf(op).Elem().FieldByName(field)
	f.SetBool(*value)
}

func (op *options) SetInt(field string, value *int) {
	f := reflect.ValueOf(op).Elem().FieldByName(field)
	f.SetInt(int64(*value))
}

var flags = []*ewords.FlagDef{
	{Flag: "file", Field: "FilePath", Help: "path to a source file", Value: ""},
	{Flag: "dir", Field: "DirPath", Help: "path to a source directory", Value: ""},
	{Flag: "notion", Field: "NotionDB", Help: "id of source notion database", Value: ""},
	{Flag: "example", Field: "ExportExample", Help: "do export examples", Value: false},
	{Flag: "defenition", Field: "ExportDefenition", Help: "do export defenitions", Value: false},
	{Flag: "exdir", Field: "ExampleDir", Help: "where store examples output", Value: "ewords_example"},
	{Flag: "defdir", Field: "DefenitionDir", Help: "where store defenitions output", Value: "ewords_defenition"},
}

var op = options{}

func main() {
	err := ewords.ParseArgs(flags, &op)
	if err != nil {
		panic(err)
	}

	validateOptions()

	runMessage()

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

	for file, ts := range files {
		if op.ExportExample {
			exportExamplesCmd(file, ts)
		}
		if op.ExportDefenition {
			exportDefenitionsCmd(file, ts)
		}
	}

}

func validateOptions() {

	if !needExport() {
		panic("at least one option should be set: example, defenition")
	}
	if !sourceSet() {
		panic("source should be set: file, notion, dir")
	}
	if !oneSource() {
		panic("only one source should be set: file, notion, dir")
	}
	if !dirCanBeCreated(op.ExampleDir) {
		panic(fmt.Sprintf("dir %v can't be created", op.ExampleDir))
	}
	if !dirCanBeCreated(op.DefenitionDir) {
		panic(fmt.Sprintf("dir %v can't be created", op.DefenitionDir))
	}
}

func needExport() bool {
	return op.ExportDefenition || op.ExportExample
}

func sourceSet() bool {
	return op.FilePath != "" || op.DirPath != "" || op.NotionDB != ""
}

func oneSource() bool {
	return (op.NotionDB != "" && op.FilePath == "" && op.DirPath == "") ||
		(op.NotionDB == "" && op.FilePath != "" && op.DirPath == "") ||
		(op.NotionDB == "" && op.FilePath == "" && op.DirPath != "")
}

func runMessage() {
	var source string
	if op.NotionDB != "" {
		source = fmt.Sprintf("notion db %v", op.NotionDB)
	} else if op.FilePath != "" {
		source = fmt.Sprintf("file %v", op.FilePath)
	} else if op.DirPath != "" {
		source = fmt.Sprintf("directory %v", op.DirPath)
	} else {
		panic(fmt.Sprintf("You shouldn't be here. Check opts: %v", utils.PrettyPrint(op)))
	}

	var result string
	if op.ExportExample {
		path, _ := ewords.RealPath(op.ExampleDir)
		result += fmt.Sprintf("examples will be saved to %v\n", path)
	}
	if op.ExportDefenition {
		path, _ := ewords.RealPath(op.DefenitionDir)
		result += fmt.Sprintf("defenitions will be saved to %v\n", path)
	}

	fmt.Printf("Source: %v\nTargets:\n%v", source, result)
}

func dirCanBeCreated(dir string) bool {
	res, _ := ewords.DirCanBeCreated(&dir, AppFS)
	return res == ewords.DirExists || res == ewords.CanCreate
}

func exportExamplesCmd(file string, ts []ewords.TermSource) {
	filename := fmt.Sprintf("%v-examples", path.Base(file))
	exportExamples(ts, path.Join(op.ExampleDir, filename))
}

func exportDefenitionsCmd(file string, ts []ewords.TermSource) {
	filename := fmt.Sprintf("%v-defenitions", path.Base(file))
	exportDefenitions(ts, path.Join(op.DefenitionDir, filename))
}

func exportExamples(ts []ewords.TermSource, exmplFile string) {
	ep := ewords.ExampleFileExporter{
		FS:              AppFS,
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
		FS:              AppFS,
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
	ts, err := parse.Source2TermSource(&file, &parse.TsvParser{FS: AppFS})
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
		panic(fmt.Sprintf("error while parsing Notion source to TermSource - %v", err))
	}
	return ts
}

func parseDirCmd(dir string) []string {
	paths, err := ewords.ListFiles(&dir, AppFS)
	if err != nil {
		panic(fmt.Sprintf("can't list files in: %v, %v", dir, err))
	}
	return paths
}
