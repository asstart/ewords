//go:build integration
// +build integration

package notion_test

import (
	"context"
	"github.com/asstart/english-words/app/notion"
	"github.com/asstart/english-words/app/utils"
	"os"
	"testing"
)

func TestQueryDatabase(t *testing.T) {
	nc := notion.CreateNotionClient(os.Getenv("NOTION_API_KEY"))
	ctx := context.Background()

	query := notion.DatabaseQuery{}
	database := "fffa4d9f5e4640e6815290991ad32919"
	resp, err := nc.QueryDatabase(ctx, query, database)
	if err != nil {
		t.Fatalf("Error while executing request: %v", err)
	}
	utils.AssertEqualsInt(t, 1, len(resp.Results))
	page := resp.Results[0]
	utils.AssertEqualsString(t, "8cb8ae71-33b5-4802-a3ed-be7dec66b25c", *page.ID)
	utils.AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *page.CreatedBy.ID)
	utils.AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *page.LastEditedBy.ID)
	utils.AssertEqualsString(t, "database_id", *page.Parent.Type)
	utils.AssertEqualsString(t, "fffa4d9f-5e46-40e6-8152-90991ad32919", *page.Parent.DatabaseID)
	utils.AssertEqualsBool(t, false, *page.Archived)
	props := page.Properties.(notion.PageProperties)
	v, ok := props["Data"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "value", *v.RichText[0].PlainText)
	v, ok = props["Name"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "key", *v.Title[0].PlainText)
}
