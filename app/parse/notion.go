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
	err = np.markHandled(pages)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (np *NotionParser) downloadSource(source *string) ([]notion.Page, error) {
	nc := notion.CreateNotionClient(*np.ApiKey)
	ctx := context.Background()

	var t = false
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
		term, errTerm := extractTitleProperty(p, termKey)
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
	text, ok := p.Properties.(notion.PageProperties)[property]
	if !ok {
		return "", fmt.Errorf("property: %v not found in page: %v", property, p)
	}
	if len(text.RichText) > 1 {
		return "", fmt.Errorf("property: %v is ambigious: %v", property, text.RichText)
	}
	if len(text.RichText) == 0 {
		return "", nil
	}
	return *text.RichText[0].PlainText, nil
}

func extractTitleProperty(p notion.Page, property string) (string, error) {
	title, ok := p.Properties.(notion.PageProperties)[property]
	if !ok {
		return "", fmt.Errorf("property: %v not found in page: %v", property, p)
	}
	if len(title.Title) != 1 {
		return "", fmt.Errorf("property: %v is ambigious: %v", property, title.Title)
	}
	return *title.Title[0].PlainText, nil
}

func (np *NotionParser) markHandled(pages []notion.Page) error {
	updated := []string{}
	for _, p := range pages {
		nc := notion.CreateNotionClient(*np.ApiKey)
		ctx := context.Background()

		t := true
		up := notion.UpdatePage{
			Properties: notion.PageProperties{
				handledKey: notion.PageProperty{
					Checkbox: &t,
				},
			},
		}

		_, err := nc.UpdatePage(ctx, up, *p.ID)
		if err == nil {
			updated = append(updated, *p.ID)
		} else {
			rerr := np.rollback(updated)
			if rerr != nil {
				return fmt.Errorf("failed to update %v, err: %v, rollback failed: %v", p, err, rerr)
			}
			return fmt.Errorf("failed to update %v, err: %v, rollback successful", p, err)
		}
	}
	return nil
}

func (np *NotionParser) rollback(ids []string) error {
	rolbacked := []string{}
	for _, id := range ids {
		nc := notion.CreateNotionClient(*np.ApiKey)
		ctx := context.Background()

		f := false
		up := notion.UpdatePage{
			Properties: notion.PageProperties{
				handledKey: notion.PageProperty{
					Checkbox: &f,
				},
			},
		}

		_, err := nc.UpdatePage(ctx, up, id)
		if err != nil {
			return fmt.Errorf("rollback unsucsessfull. Rolbacked: %v", rolbacked)
		} else {
			rolbacked = append(rolbacked, id)
		}
	}
	return nil
}
