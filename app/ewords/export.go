package ewords

import (
	"fmt"

	"github.com/spf13/afero"
)

type ExampleFileExporter struct {
	FS              afero.Fs
	FilePath        string
	Formatter       TermExampleFormatter
	OutputFormatter OutputFormatter
}

type DefenitionFileExporter struct {
	FS              afero.Fs
	FilePath        string
	Formatter       TermDefenitionFormatter
	OutputFormatter OutputFormatter
}

func (fp *ExampleFileExporter) ExportExample(te []TermExample) error {
	formatted := []string{}
	for _, exm := range te {
		f, err := fp.Formatter.FormatExample(&exm)
		if err != nil {
			return fmt.Errorf("error while exporting example, error: %v", err)
		}
		formatted = append(formatted, *f)

	}
	output := fp.OutputFormatter.FormatOutput(formatted)

	err := WriteFile(fp.FilePath, output, fp.FS)
	if err != nil {
		return fmt.Errorf("error while exporting examples. source: %v, error: %v", te, err)
	}
	return nil
}

func (fp *DefenitionFileExporter) ExportDefenition(td []TermDefenition) error {
	formatted := []string{}
	for _, df := range td {
		f, err := fp.Formatter.FormatDefenition(&df)
		if err != nil {
			return fmt.Errorf("error while exporting defenition, error: %v", err)
		}
		formatted = append(formatted, *f)
	}
	output := fp.OutputFormatter.FormatOutput(formatted)

	err := WriteFile(fp.FilePath, output, fp.FS)
	if err != nil {
		return fmt.Errorf("error while exporting defenitions. source: %v, error: %v", td, err)
	}
	return nil
}
