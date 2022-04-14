package ewords

import (
	"fmt"
	"regexp"
	"strings"
)

type MemriseFormatter struct{}

type MemriseExamplePublisher struct {
	FilePath        string
	Formatter       TermExampleFormatter
	OutputFormatter OutputFormatter
}

type MemriseDefenitionPublisher struct {
	FilePath        string
	Formatter       TermDefenitionFormatter
	OutputFormatter OutputFormatter
}

func (*MemriseFormatter) FormatExample(te *TermExample) *string {
	pattern := regexp.MustCompile(te.Term)
	pluggedEx := pattern.ReplaceAllLiteralString(te.Example, plug)
	res := fmt.Sprintf("%v\t``\t%v", te.Term, pluggedEx)
	return &res
}

func (*MemriseFormatter) FormatDefenition(td *TermDefenition) *string {
	res := fmt.Sprintf("%v\t%v\t%v", td.Term, td.Transcription, td.Defenition)
	return &res
}

func (*MemriseFormatter) FormatOutput(source []string) *string {
	res := strings.Join(source, "\n")
	return &res
}

func (fp *MemriseExamplePublisher) PublishExample(te []TermExample) error {
	formatted := []string{}
	for _, exm := range te {
		formatted = append(formatted, *fp.Formatter.FormatExample(&exm))
	}
	output := fp.OutputFormatter.FormatOutput(formatted)

	return WriteFile(fp.FilePath, output)
}

func (fp *MemriseDefenitionPublisher) PublishDefenition(te []TermDefenition) error {
	formatted := []string{}
	for _, exm := range te {
		formatted = append(formatted, *fp.Formatter.FormatDefenition(&exm))
	}
	output := fp.OutputFormatter.FormatOutput(formatted)

	return WriteFile(fp.FilePath, output)
}
