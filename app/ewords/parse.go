package ewords

import (
	"fmt"
	"strings"
)

type TermSourceParser interface {
	ParseTermSource(source *string) (*TermSource, error)
}

type InputParser interface {
	ParseInput(input *string) []string
}

type TsvParser struct{}

func Str2TermSource(source *string, ip InputParser, tsp TermSourceParser) ([]TermSource, error) {
	lines := ip.ParseInput(source)
	res := []TermSource{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		ts, err := tsp.ParseTermSource(&line)
		if err != nil {
			return nil, fmt.Errorf("error while parsing source %v to TermSource: %v", line, err)
		}
		res = append(res, *ts)
	}
	return res, nil
}

func (*TsvParser) ParseTermSource(source *string) (*TermSource, error) {
	parts := strings.Split(*source, "\t")
	if len(parts) != 4 {
		return nil, fmt.Errorf("source string %v - doesn't match with template: term transcription defenition example", parts)
	}
	return &TermSource{
		Term:          parts[0],
		Transcription: parts[1],
		Definition:    parts[2],
		Example:       parts[3],
	}, nil
}

func (*TsvParser) ParseInput(source *string) []string {
	return strings.Split(*source, "\n")
}
