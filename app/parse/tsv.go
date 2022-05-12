package parse

import (
	"fmt"
	"strings"

	"github.com/asstart/english-words/app/ewords"
	"github.com/spf13/afero"
)

type TsvParser struct {
	FS afero.Fs
}

func (tp *TsvParser) ParseTermSource(source *string) ([]ewords.TermSource, error) {
	filedata, err := ewords.ReadFile(source, tp.FS)
	if err != nil {
		panic(fmt.Sprintf("Error while reading %v - %v", source, err))
	}

	lines := parseInput(&filedata)
	res := []ewords.TermSource{}
	for _, line := range lines {
		if empty(&line) {
			continue
		}
		ts, err := parseSingle(&line)
		if err != nil {
			return nil, fmt.Errorf("error while parsing line %v to TermSource: %v", line, err)
		}
		res = append(res, *ts)
	}
	return res, nil
}

func empty(str *string) bool {
	return strings.TrimSpace(*str) == ""
}

func parseSingle(source *string) (*ewords.TermSource, error) {
	parts := strings.Split(*source, "\t")
	if len(parts) != 4 {
		return nil, fmt.Errorf("source string %v - doesn't match with template: term transcription defenition example", parts)
	}
	return &ewords.TermSource{
		Term:          parts[0],
		Transcription: parts[1],
		Definition:    parts[2],
		Example:       parts[3],
	}, nil
}

func parseInput(source *string) []string {
	return strings.Split(*source, "\n")
}
