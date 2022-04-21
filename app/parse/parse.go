package parse

import (
	"fmt"
	"github.com/asstart/english-words/app/ewords"
)

type TermSourceParser interface {
	ParseTermSource(source *string) ([]ewords.TermSource, error)
}

func Source2TermSource(source *string, tsp TermSourceParser) ([]ewords.TermSource, error) {
	res, err := tsp.ParseTermSource(source)
	if err != nil {
		return nil, fmt.Errorf("error while parsing source %v to []TermSource: %v", source, err)
	}
	return res, nil
}
