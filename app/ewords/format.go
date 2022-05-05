package ewords

import (
	"fmt"
	"regexp"
	"strings"
)

const plug string = "_____"
const tsvTemplate string = "%v\t%v"

type TermExampleFormatter interface {
	FormatExample(te *TermExample) (*string, error)
}

type TermDefenitionFormatter interface {
	FormatDefenition(td *TermDefenition) (*string, error)
}

type OutputFormatter interface {
	FormatOutput(lines []string) *string
}

type SimpleTsv struct{}

func (*SimpleTsv) FormatExample(te *TermExample) (*string, error) {
	pluggedEx, err := dumbPlug(te.Term, te.Example)
	if err != nil {
		return nil, err
	}
	res := fmt.Sprintf(tsvTemplate, te.Term, pluggedEx)
	return &res, nil
}

func dumbPlug(term string, phrase string) (string, error) {
	rx := fmt.Sprintf(`(?i)%v`, term)
	pattern, err := regexp.Compile(rx)
	if err != nil {
		return "", fmt.Errorf("error while making plug for example, err: %v", err)
	}
	return pattern.ReplaceAllLiteralString(phrase, plug), nil
}

func (*SimpleTsv) FormatDefenition(te *TermDefenition) (*string, error) {
	res := fmt.Sprintf(tsvTemplate, te.Term, te.Defenition)
	return &res, nil
}

func (*SimpleTsv) FormatOutput(source []string) *string {
	res := strings.Join(source, "\n")
	return &res
}
