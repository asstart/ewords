package ewords

import ()

const plug string = "_____"

type TermExampleFormatter interface {
	FormatExample(te *TermExample) *string
}

type TermDefenitionFormatter interface {
	FormatDefenition(td *TermDefenition) *string
}

type OutputFormatter interface {
	FormatOutput(lines []string) *string
}

func TermExamples2Str(te []TermExample, tef TermExampleFormatter, of OutputFormatter) *string {
	formatted := []string{}
	for _, exmp := range te {
		formatted = append(formatted, *tef.FormatExample(&exmp))
	}
	return of.FormatOutput(formatted)
}

func TermDefenitions2Str(td []TermDefenition, tdf TermDefenitionFormatter, of OutputFormatter) *string {
	formatted := []string{}
	for _, exmp := range td {
		formatted = append(formatted, *tdf.FormatDefenition(&exmp))
	}
	return of.FormatOutput(formatted)
}
