// +build integration

package notion

import (
	"context"
	"testing"
	"os"
)

func TestQueryDatabase(t *testing.T) {
	nc := CreateNotionClient(os.Getenv("NOTION_API_KEY"))
	ctx := context.Background()

	query := DatabaseQuery{}
	database := "fffa4d9f5e4640e6815290991ad32919"
	resp, err := nc.QueryDatabase(ctx, query, database)
	if err != nil {
		t.Fatalf("Error while executing request: %v", err)
	}
	AssertEqualsInt(t, 1, len(resp.Results))
	page := resp.Results[0]
	AssertEqualsString(t, "0d7a66fd-bf30-49e5-a397-37721038b0d5", *page.ID)
	AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *page.CreatedBy.ID)
	AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *page.LastEditedBy.ID)
	AssertEqualsString(t, "database_id", *page.Parent.Type)
	AssertEqualsString(t, "fffa4d9f-5e46-40e6-8152-90991ad32919", *page.Parent.DatabaseID)
	AssertEqualsBool(t, false, *page.Archived)
	props := page.Properties.(PageProperties)
	v, ok := props["Data"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "value", *v.RichText[0].PlainText)
	v, ok = props["Name"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "key", *v.Title[0].PlainText)
}
