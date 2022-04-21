package parse

import (
	"context"
	"fmt"
	"strings"

	"github.com/asstart/english-words/app/ewords"
	"github.com/asstart/english-words/app/notion"
)

type NotionParser struct {
	ApiKey *string
}

const (
	termKey    = "Word"
	transKey   = "Transcription"
	defKey     = "Defenition"
	exmpKey    = "Example"
	handledKey = "Handled"
)

func (np *NotionParser) ParseTermSource(source *string) ([]ewords.TermSource, error) {
	pages, err := np.downloadSource(source)
	if err != nil {
		return nil, err
	}
	ts, err := page2term(pages)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (np *NotionParser) downloadSource(source *string) ([]notion.Page, error) {
	nc := notion.CreateNotionClient(*np.ApiKey)
	ctx := context.Background()

	var t = true
	var h = handledKey
	query := notion.DatabaseQuery{
		Filter: &notion.Filter{
			Property: &h,
			CheckBox: &notion.CheckboxCondition{
				Equals: &t,
			},
		},
	}

	var res []notion.Page
	finished := false
	for !finished {
		pages, err := nc.QueryDatabase(ctx, query, *source)
		if err != nil {
			return nil, err
		}
		res = append(res, pages.Results...)
		if *pages.HasMore {
			query = notion.DatabaseQuery{
				StartCursor: pages.NextCursor,
			}
		} else {
			finished = true
		}
	}

	return res, nil
}

func page2term(pages []notion.Page) ([]ewords.TermSource, error) {
	var res []ewords.TermSource
	for _, p := range pages {
		term, errTerm := extractTextProperty(p, termKey)
		transcription, errTrans := extractTextProperty(p, transKey)
		defenition, errDef := extractTextProperty(p, defKey)
		example, errExmpl := extractTextProperty(p, exmpKey)

		err := compose(errTerm, errTrans, errDef, errExmpl)
		if err != nil {
			return nil, err
		}

		ts := ewords.TermSource{
			Term:          term,
			Transcription: transcription,
			Definition:    defenition,
			Example:       example,
		}
		res = append(res, ts)
	}
	return res, nil
}

func compose(errors ...error) error {
	var res []string
	for _, e := range errors {
		if e != nil {
			res = append(res, e.Error())
		}
	}
	if len(res) == 0 {
		return nil
	} else {
		return fmt.Errorf(strings.Join(res, ";"))
	}
}

func extractTextProperty(p notion.Page, property string) (string, error) {
	term, ok := p.Properties.(notion.PageProperties)[property]
	if !ok {
		return "", fmt.Errorf("property: %v not found in page: %v", property, p)
	}
	if len(term.RichText) != 1 {
		return "", fmt.Errorf("property: %v is ambigious: %v", property, term.RichText)
	}
	return *term.RichText[0].PlainText, nil
}
