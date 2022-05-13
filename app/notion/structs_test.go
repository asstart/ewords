//go:build !integration
// +build !integration

package notion_test

import (
	"encoding/json"
	"github.com/asstart/english-words/app/notion"
	"github.com/asstart/english-words/app/utils"
	"testing"
)

func TestUnmarshalTextCondition(t *testing.T) {
	source := `
	{
		"equals":"equals_string",
		"does_not_equal":"not_equal_string",
		"contains":"contains_string",
		"does_not_contain":"does_not_contain_string",
		"starts_with":"starts_with_string",
		"ends_with":"ends_with_string",
		"is_empty":true,
		"is_not_empty":true
	}
	`

	var text notion.TextCondition
	json.Unmarshal([]byte(source), &text)

	utils.AssertEqualsString(t, "equals_string", *text.Equals)
	utils.AssertEqualsString(t, "not_equal_string", *text.DoesntEqual)
	utils.AssertEqualsString(t, "contains_string", *text.Contains)
	utils.AssertEqualsString(t, "does_not_contain_string", *text.DoesntContain)
	utils.AssertEqualsString(t, "starts_with_string", *text.StartsWith)
	utils.AssertEqualsString(t, "ends_with_string", *text.EndsWith)
	utils.AssertEqualsBool(t, true, *text.IsEmpty)
	utils.AssertEqualsBool(t, true, *text.IsNotEmpty)
}

func TestUnmarshalDatabasePage(t *testing.T) {
	source := `
	{
		"object": "list",
		"results": [
		  {
			"object": "page",
			"id": "4f5cb747-e1d5-46a8-9bcd-48df3a21d675",
			"created_time": "2022-04-17T12:36:00.000Z",
			"last_edited_time": "2022-04-18T22:33:00.000Z",
			"created_by": {
			  "object": "user",
			  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
			},
			"last_edited_by": {
			  "object": "user",
			  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
			},
			"cover": {
			  "type": "external",
			  "external": {
				"url": "https://www.notion.so/images/page-cover/rijksmuseum_vermeer_the_milkmaid.jpg"
			  }
			},
			"icon": {
			  "type": "emoji",
			  "emoji": "😁"
			},
			"parent": {
			  "type": "database_id",
			  "database_id": "ffd97167-8029-47f8-8623-c2347dd9c563"
			},
			"archived": false,
			"properties": {
			  "Property 2": {
				"id": "%3A%60gb",
				"type": "date",
				"date": {
				  "start": "2022-04-04",
				  "end": null,
				  "time_zone": null
				}
			  },
			  "Property 4": {
				"id": "%40LUH",
				"type": "files",
				"files": [
				  {
					"name": "2.jpg",
					"type": "file",
					"file": {
					  "url": "https://s3.us-west-2.amazonaws.com/secure.notion-static.com/8a85ee6e-dfa5-42b8-8900-9cb92d490d54/2.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220418%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220418T223419Z&X-Amz-Expires=3600&X-Amz-Signature=326c1cca24e649c1fb2861644212546f26ad2f11b7c4e72426a5d648ad868700&X-Amz-SignedHeaders=host&x-id=GetObject",
					  "expiry_time": "2022-04-18T23:34:19.557Z"
					}
				  }
				]
			  },
			  "Tags": {
				"id": "%40Lx%7C",
				"type": "multi_select",
				"multi_select": [
				  {
					"id": "a0f30ec1-f592-4ec2-a69a-71fbd2eafbf3",
					"name": "word",
					"color": "blue"
				  },
				  {
					"id": "28d91a14-b38b-4d3c-afa6-738b3864ae0f",
					"name": "drow",
					"color": "purple"
				  }
				]
			  },
			  "updated_at": {
				"id": "Az%3Cf",
				"type": "last_edited_time",
				"last_edited_time": "2022-04-18T22:33:00.000Z"
			  },
			  "Property 6": {
				"id": "KP%7CT",
				"type": "url",
				"url": "https://pkg.go.dev/github.com/asstart/english-words/app/ewords/notion#Filter.Property"
			  },
			  "Defenition": {
				"id": "OIab",
				"type": "rich_text",
				"rich_text": []
			  },
			  "Property 13": {
				"id": "Rl%3F%3D",
				"type": "last_edited_by",
				"last_edited_by": {
				  "object": "user",
				  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
				}
			  },
			  "Property": {
				"id": "SUv~",
				"type": "number",
				"number": 1234
			  },
			  "Property 1": {
				"id": "TF%5CC",
				"type": "select",
				"select": {
				  "id": "3d34ca02-5de7-491c-a306-1bab98285a72",
				  "name": "prope",
				  "color": "red"
				}
			  },
			  "Transcription": {
				"id": "TmLV",
				"type": "rich_text",
				"rich_text": [
				  {
					"type": "text",
					"text": {
					  "content": "hello",
					  "link": null
					},
					"annotations": {
					  "bold": false,
					  "italic": false,
					  "strikethrough": false,
					  "underline": false,
					  "code": false,
					  "color": "default"
					},
					"plain_text": "hello",
					"href": null
				  }
				]
			  },
			  "Property 7": {
				"id": "Trvl",
				"type": "email",
				"email": "some@gmail.com"
			  },
			  "Property 3": {
				"id": "%5EbRD",
				"type": "people",
				"people": [
				  {
					"object": "user",
					"id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
				  }
				]
			  },
			  "Property 12": {
				"id": "_cww",
				"type": "created_by",
				"created_by": {
				  "object": "user",
				  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
				}
			  },
			  "Example": {
				"id": "cjRY",
				"type": "rich_text",
				"rich_text": [
				  {
					"type": "text",
					"text": {
					  "content": "Hello there",
					  "link": null
					},
					"annotations": {
					  "bold": false,
					  "italic": false,
					  "strikethrough": false,
					  "underline": false,
					  "code": false,
					  "color": "default"
					},
					"plain_text": "Hello there",
					"href": null
				  }
				]
			  },
			  "handled": {
				"id": "il~C",
				"type": "checkbox",
				"checkbox": true
			  },
			  "Property 5": {
				"id": "m~bE",
				"type": "checkbox",
				"checkbox": false
			  },
			  "created_at": {
				"id": "o~QP",
				"type": "created_time",
				"created_time": "2022-04-17T12:36:00.000Z"
			  },
			  "Property 9": {
				"id": "u%3Bcf",
				"type": "formula",
				"formula": {
				  "type": "boolean",
				  "boolean": false
				}
			  },
			  "Property 8": {
				"id": "z%5D%3AD",
				"type": "phone_number",
				"phone_number": "+79998887766"
			  },
			  "Word": {
				"id": "title",
				"type": "title",
				"title": [
				  {
					"type": "text",
					"text": {
					  "content": "Hello",
					  "link": null
					},
					"annotations": {
					  "bold": false,
					  "italic": false,
					  "strikethrough": false,
					  "underline": false,
					  "code": false,
					  "color": "default"
					},
					"plain_text": "Hello",
					"href": null
				  }
				]
			  }
			},
			"url": "https://www.notion.so/Hello-4f5cb747e1d546a89bcd48df3a21d675"
		  }
		],
		"next_cursor": null,
		"has_more": false,
		"type": "page",
		"page": {}
	  }
	`

	var pages notion.DatabasePages
	json.Unmarshal([]byte(source), &pages)

	utils.AssertEqualsBool(t, false, *pages.HasMore)
	utils.AssertEqualsInt(t, 1, len(pages.Results))
	page := pages.Results[0]
	utils.AssertEqualsString(t, "4f5cb747-e1d5-46a8-9bcd-48df3a21d675", *page.ID)
	// AssertEqualsTime(t, time.)
	utils.AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *page.CreatedBy.ID)
	utils.AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *page.LastEditedBy.ID)
	utils.AssertEqualsString(t, "external", *page.Cover.Type)
	utils.AssertEqualsString(t, "https://www.notion.so/images/page-cover/rijksmuseum_vermeer_the_milkmaid.jpg", *page.Cover.ExternalFile.URL)
	utils.AssertEqualsString(t, "emoji", *page.Icon.Type)
	utils.AssertEqualsString(t, "😁", *page.Icon.Emoji)
	utils.AssertEqualsString(t, "database_id", *page.Parent.Type)
	utils.AssertEqualsString(t, "ffd97167-8029-47f8-8623-c2347dd9c563", *page.Parent.DatabaseID)
	utils.AssertEqualsBool(t, false, *page.Archived)
	props := page.Properties.(notion.PageProperties)
	v, ok := props["Property 2"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "%3A%60gb", *v.ID)
	utils.AssertEqualsString(t, "date", *v.Type)
	utils.AssertEqualsString(t, "2022-04-04", *v.Date.Start)
	utils.AssertNil(t, v.Date.End)
	utils.AssertNil(t, v.Date.TimeZone)
	v, ok = props["Property 4"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "files", *v.Type)
	utils.AssertEqualsInt(t, 1, len(v.Files))
	utils.AssertEqualsString(t, "2.jpg", *v.Files[0].Name)
	utils.AssertEqualsString(t, "file", *v.Files[0].Type)
	utils.AssertEqualsString(
		t,
		"https://s3.us-west-2.amazonaws.com/secure.notion-static.com/8a85ee6e-dfa5-42b8-8900-9cb92d490d54/2.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220418%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220418T223419Z&X-Amz-Expires=3600&X-Amz-Signature=326c1cca24e649c1fb2861644212546f26ad2f11b7c4e72426a5d648ad868700&X-Amz-SignedHeaders=host&x-id=GetObject",
		*v.Files[0].NotionFile.URL)
	// AssertEqualsTime(t, time.Parse("2022-04-18T23:34:19.557Z"), *v.Files[0].NotionFile.ExpireTime)
	v, ok = props["Tags"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "%40Lx%7C", *v.ID)
	utils.AssertEqualsString(t, "multi_select", *v.Type)
	utils.AssertEqualsInt(t, 2, len(v.MultiSelect))
	utils.AssertEqualsString(t, "a0f30ec1-f592-4ec2-a69a-71fbd2eafbf3", *v.MultiSelect[0].ID)
	utils.AssertEqualsString(t, "word", *v.MultiSelect[0].Name)
	utils.AssertEqualsString(t, "blue", *v.MultiSelect[0].Color)
	utils.AssertEqualsString(t, "28d91a14-b38b-4d3c-afa6-738b3864ae0f", *v.MultiSelect[1].ID)
	utils.AssertEqualsString(t, "drow", *v.MultiSelect[1].Name)
	utils.AssertEqualsString(t, "purple", *v.MultiSelect[1].Color)
	v, ok = props["updated_at"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "Az%3Cf", *v.ID)
	utils.AssertEqualsString(t, "last_edited_time", *v.Type)
	// AssertEqualsTime(t, , v.LastEditedTime)
	v, ok = props["Property 6"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "KP%7CT", *v.ID)
	utils.AssertEqualsString(t, "url", *v.Type)
	utils.AssertEqualsString(t, "https://pkg.go.dev/github.com/asstart/english-words/app/ewords/notion#Filter.Property", *v.URL)
	v, ok = props["Defenition"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "OIab", *v.ID)
	utils.AssertEqualsString(t, "rich_text", *v.Type)
	utils.AssertEqualsInt(t, 0, len(v.RichText))
	v, ok = props["Property 13"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "Rl%3F%3D", *v.ID)
	utils.AssertEqualsString(t, "last_edited_by", *v.Type)
	utils.AssertEqualsString(t, "user", *v.LastEditedBy.Object)
	utils.AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *v.LastEditedBy.ID)
	v, ok = props["Property"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "SUv~", *v.ID)
	utils.AssertEqualsString(t, "number", *v.Type)
	utils.AssertEqualsInt(t, 1234, int(*v.Number))

	v, ok = props["Transcription"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "TmLV", *v.ID)
	utils.AssertEqualsString(t, "rich_text", *v.Type)
	utils.AssertEqualsInt(t, 1, len(v.RichText))
	utils.AssertEqualsString(t, "text", *v.RichText[0].Type)
	utils.AssertEqualsString(t, "hello", *v.RichText[0].Text.Content)
	utils.AssertNil(t, v.RichText[0].Text.Link)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Bold)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Italic)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Strikethrough)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Underline)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Code)
	utils.AssertEqualsString(t, "default", *v.RichText[0].Annotations.Color)
	utils.AssertEqualsString(t, "hello", *v.RichText[0].PlainText)
	utils.AssertNil(t, v.RichText[0].Href)

	v, ok = props["Property 7"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "Trvl", *v.ID)
	utils.AssertEqualsString(t, "email", *v.Type)
	utils.AssertEqualsString(t, "some@gmail.com", *v.Email)

	v, ok = props["Property 3"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "%5EbRD", *v.ID)
	utils.AssertEqualsString(t, "people", *v.Type)
	utils.AssertEqualsInt(t, 1, len(v.People))
	utils.AssertEqualsString(t, "user", *v.People[0].Object)
	utils.AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *v.People[0].ID)

	v, ok = props["Property 12"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "_cww", *v.ID)
	utils.AssertEqualsString(t, "created_by", *v.Type)
	utils.AssertEqualsString(t, "user", *v.CreatedBy.Object)
	utils.AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *v.CreatedBy.ID)

	v, ok = props["Example"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "cjRY", *v.ID)
	utils.AssertEqualsString(t, "rich_text", *v.Type)
	utils.AssertEqualsInt(t, 1, len(v.RichText))
	utils.AssertEqualsString(t, "text", *v.RichText[0].Type)
	utils.AssertEqualsString(t, "Hello there", *v.RichText[0].Text.Content)
	utils.AssertNil(t, v.RichText[0].Text.Link)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Bold)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Italic)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Strikethrough)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Underline)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Code)
	utils.AssertEqualsString(t, "default", *v.RichText[0].Annotations.Color)
	utils.AssertEqualsString(t, "Hello there", *v.RichText[0].PlainText)
	utils.AssertNil(t, v.RichText[0].Href)

	v, ok = props["handled"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "il~C", *v.ID)
	utils.AssertEqualsString(t, "checkbox", *v.Type)
	utils.AssertEqualsBool(t, true, *v.Checkbox)

	v, ok = props["Property 5"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "m~bE", *v.ID)
	utils.AssertEqualsString(t, "checkbox", *v.Type)
	utils.AssertEqualsBool(t, false, *v.Checkbox)

	v, ok = props["created_at"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "o~QP", *v.ID)
	utils.AssertEqualsString(t, "created_time", *v.Type)
	//TODO check time

	v, ok = props["Property 9"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "u%3Bcf", *v.ID)
	utils.AssertEqualsString(t, "formula", *v.Type)
	utils.AssertEqualsString(t, "boolean", *v.Formula.Type)
	utils.AssertEqualsBool(t, false, *v.Formula.BooleanFormula)

	v, ok = props["Property 8"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "z%5D%3AD", *v.ID)
	utils.AssertEqualsString(t, "phone_number", *v.Type)
	utils.AssertEqualsString(t, "+79998887766", *v.PhoneNumber)

	v, ok = props["Word"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "title", *v.ID)
	utils.AssertEqualsString(t, "title", *v.Type)
	utils.AssertEqualsInt(t, 1, len(v.Title))
	utils.AssertEqualsString(t, "text", *v.Title[0].Type)
	utils.AssertEqualsString(t, "Hello", *v.Title[0].Text.Content)
	utils.AssertNil(t, v.Title[0].Text.Link)
	utils.AssertEqualsBool(t, false, *v.Title[0].Annotations.Bold)
	utils.AssertEqualsBool(t, false, *v.Title[0].Annotations.Italic)
	utils.AssertEqualsBool(t, false, *v.Title[0].Annotations.Strikethrough)
	utils.AssertEqualsBool(t, false, *v.Title[0].Annotations.Underline)
	utils.AssertEqualsBool(t, false, *v.Title[0].Annotations.Code)
	utils.AssertEqualsString(t, "default", *v.Title[0].Annotations.Color)
	utils.AssertEqualsString(t, "Hello", *v.Title[0].PlainText)
	utils.AssertNil(t, v.Title[0].Href)
}

func TestMarshalRichTextContainsFilter(t *testing.T) {
	var f = notion.Filter{
		Property: utils.StrPtr("Landmark"),
		RichText: &notion.TextCondition{

			Contains: utils.StrPtr("Bridge"),
		},
	}

	expected := `
	{
		"property": "Landmark",
		"rich_text": {
			"contains": "Bridge"
		}
	}
	`

	marshaled, _ := json.Marshal(f)

	utils.AssertEqualsString(t, utils.Minimise(expected), utils.Minimise(string(marshaled)))
}

func TestMarshalCompoundFilter(t *testing.T) {
	var f = notion.Filter{
		And: []notion.Filter{
			notion.Filter{
				Property: utils.StrPtr("Seen"),
				CheckBox: &notion.CheckboxCondition{
					Equals: utils.BoolPtr(false),
				},
			},
			notion.Filter{
				Property: utils.StrPtr("Yearly visitor count"),
				Number: &notion.NumberCondition{
					GreaterThan: utils.FloatPtr(1000000),
				},
			},
		},
	}

	expected := `
	{
		"and": [
			{
				"property": "Seen",
				"checkbox": {
					"equals": false
				}
			},
			{
				"property": "Yearly visitor count",
				"number": {
					"greater_than": 1000000
				}
			}
		]
	}
	`

	marshaled, _ := json.Marshal(f)

	utils.AssertEqualsString(t, utils.Minimise(expected), utils.Minimise(string(marshaled)))
}

func TestMarshalMiltililevelCompoundFilter(t *testing.T) {
	var f = notion.Filter{
		Or: []notion.Filter{
			notion.Filter{
				Property: utils.StrPtr("Description"),
				RichText: &notion.TextCondition{
					Contains: utils.StrPtr("fish"),
				},
			},
			notion.Filter{
				And: []notion.Filter{
					notion.Filter{
						Property: utils.StrPtr("Food group"),
						Select: &notion.SelectCondition{
							Equals: utils.StrPtr("🥦Vegetable"),
						},
					},
					notion.Filter{
						Property: utils.StrPtr("Is protein rich?"),
						CheckBox: &notion.CheckboxCondition{
							Equals: utils.BoolPtr(true),
						},
					},
				},
			},
		},
	}

	expected := `
	{
		"or": [
			{
				"property": "Description",
				"rich_text": {
					"contains": "fish"
				}
			},
			{
				"and": [
					{
						"property": "Food group",
						"select": {
							"equals": "🥦Vegetable"
						}
					},
					{
						"property": "Is protein rich?",
						"checkbox": {
							"equals": true
						}
					}
				]
			}
		]
	}
	`

	marshaled, _ := json.Marshal(f)

	utils.AssertEqualsString(t, utils.Minimise(expected), utils.Minimise(string(marshaled)))
}

func TestSortMarshal(t *testing.T) {
	s := notion.Sort{
		Property:  utils.StrPtr("Food group"),
		Direction: utils.StrPtr("descending"),
	}

	expected := `
	{
		"property": "Food group",
		"direction": "descending"
	}
	`

	marshaled, _ := json.Marshal(s)

	utils.AssertEqualsString(t, utils.Minimise(expected), utils.Minimise(string(marshaled)))
}

func TestDatabaseQuery(t *testing.T) {
	query := notion.DatabaseQuery{
		Filter: &notion.Filter{
			Property: utils.StrPtr("Landmark"),
			RichText: &notion.TextCondition{

				Contains: utils.StrPtr("Bridge"),
			},
		},
		Sorts: []notion.Sort{
			notion.Sort{
				Property:  utils.StrPtr("Food group"),
				Direction: utils.StrPtr("descending"),
			},
			notion.Sort{
				Property:  utils.StrPtr("Name"),
				Direction: utils.StrPtr("ascending"),
			},
		},
		StartCursor: utils.StrPtr("3-295-0235"),
		PageSize:    utils.IntPtr(50),
	}

	marshaled, _ := json.Marshal(query)

	expected := `
	{
		"filter":{
			"property": "Landmark",
			"rich_text": {
				"contains": "Bridge"
			}
		},
		"sorts":[
			{
				"property": "Food group",
				"direction": "descending"
			},
			{
				"property": "Name",
				"direction": "ascending"
			}
		],
		"start_cursor":"3-295-0235",
		"page_size":50
	}
	`

	utils.AssertEqualsString(t, utils.Minimise(expected), utils.Minimise(string(marshaled)))
}

func TestUpdatePage(t *testing.T) {
	up := notion.UpdatePage{
		Properties: notion.PageProperties{
			"handled": notion.PageProperty{
				Checkbox: utils.BoolPtr(true),
			},
		},
	}

	expected := `
	{
		"properties":
		{
			"handled":
			{
				"checkbox":true
			}
		}
	}
	`

	marshaled, _ := json.Marshal(up)

	utils.AssertEqualsString(t, utils.Minimise(expected), utils.Minimise(string(marshaled)))
}

func TestUnmarshalPage(t *testing.T) {
	jsn := `
	{
		"object": "page",
		"id": "4f5cb747-e1d5-46a8-9bcd-48df3a21d675",
		"created_time": "2022-04-17T12:36:00.000Z",
		"last_edited_time": "2022-04-20T22:26:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "61b898c0-7772-48d5-886e-2a222669737c"
		},
		"cover": {
		  "type": "external",
		  "external": {
			"url": "https://www.notion.so/images/page-cover/rijksmuseum_vermeer_the_milkmaid.jpg"
		  }
		},
		"icon": {
		  "type": "emoji",
		  "emoji": "😁"
		},
		"parent": {
		  "type": "database_id",
		  "database_id": "ffd97167-8029-47f8-8623-c2347dd9c563"
		},
		"archived": false,
		"properties": {
		  "Property 2": {
			"id": "%3A%60gb",
			"type": "date",
			"date": {
			  "start": "2022-04-04",
			  "end": null,
			  "time_zone": null
			}
		  },
		  "Tags": {
			"id": "%40Lx%7C",
			"type": "multi_select",
			"multi_select": [
			  {
				"id": "a0f30ec1-f592-4ec2-a69a-71fbd2eafbf3",
				"name": "word",
				"color": "blue"
			  },
			  {
				"id": "28d91a14-b38b-4d3c-afa6-738b3864ae0f",
				"name": "drow",
				"color": "purple"
			  }
			]
		  },
		  "updated_at": {
			"id": "Az%3Cf",
			"type": "last_edited_time",
			"last_edited_time": "2022-04-20T22:26:00.000Z"
		  },
		  "Property 6": {
			"id": "KP%7CT",
			"type": "url",
			"url": "https://pkg.go.dev/github.com/asstart/english-words/app/ewords/notion#Filter.Property"
		  },
		  "Defenition": {
			"id": "OIab",
			"type": "rich_text",
			"rich_text": []
		  },
		  "Property 13": {
			"id": "Rl%3F%3D",
			"type": "last_edited_by",
			"last_edited_by": {
			  "object": "user",
			  "id": "61b898c0-7772-48d5-886e-2a222669737c",
			  "name": "ewords",
			  "avatar_url": null,
			  "type": "bot",
			  "bot": {}
			}
		  },
		  "Property": {
			"id": "SUv~",
			"type": "number",
			"number": 1234
		  },
		  "Property 1": {
			"id": "TF%5CC",
			"type": "select",
			"select": {
			  "id": "3d34ca02-5de7-491c-a306-1bab98285a72",
			  "name": "prope",
			  "color": "red"
			}
		  },
		  "Transcription": {
			"id": "TmLV",
			"type": "rich_text",
			"rich_text": [
			  {
				"type": "text",
				"text": {
				  "content": "hello",
				  "link": null
				},
				"annotations": {
				  "bold": false,
				  "italic": false,
				  "strikethrough": false,
				  "underline": false,
				  "code": false,
				  "color": "default"
				},
				"plain_text": "hello",
				"href": null
			  }
			]
		  },
		  "Property 7": {
			"id": "Trvl",
			"type": "email",
			"email": "some@gmail.com"
		  },
		  "Property 3": {
			"id": "%5EbRD",
			"type": "people",
			"people": [
			  {
				"object": "user",
				"id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
			  }
			]
		  },
		  "Property 12": {
			"id": "_cww",
			"type": "created_by",
			"created_by": {
			  "object": "user",
			  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
			}
		  },
		  "Example": {
			"id": "cjRY",
			"type": "rich_text",
			"rich_text": [
			  {
				"type": "text",
				"text": {
				  "content": "Hello there",
				  "link": null
				},
				"annotations": {
				  "bold": false,
				  "italic": false,
				  "strikethrough": false,
				  "underline": false,
				  "code": false,
				  "color": "default"
				},
				"plain_text": "Hello there",
				"href": null
			  }
			]
		  },
		  "handled": {
			"id": "il~C",
			"type": "checkbox",
			"checkbox": true
		  },
		  "Property 5": {
			"id": "m~bE",
			"type": "checkbox",
			"checkbox": false
		  },
		  "created_at": {
			"id": "o~QP",
			"type": "created_time",
			"created_time": "2022-04-17T12:36:00.000Z"
		  },
		  "Property 9": {
			"id": "u%3Bcf",
			"type": "formula",
			"formula": {
			  "type": "boolean",
			  "boolean": false
			}
		  },
		  "Property 8": {
			"id": "z%5D%3AD",
			"type": "phone_number",
			"phone_number": "+79998887766"
		  },
		  "Word": {
			"id": "title",
			"type": "title",
			"title": [
			  {
				"type": "text",
				"text": {
				  "content": "Hello",
				  "link": null
				},
				"annotations": {
				  "bold": false,
				  "italic": false,
				  "strikethrough": false,
				  "underline": false,
				  "code": false,
				  "color": "default"
				},
				"plain_text": "Hello",
				"href": null
			  }
			]
		  }
		},
		"url": "https://www.notion.so/Hello-4f5cb747e1d546a89bcd48df3a21d675"
	  }
	`

	var page notion.Page
	json.Unmarshal([]byte(jsn), &page)

	utils.AssertEqualsString(t, "4f5cb747-e1d5-46a8-9bcd-48df3a21d675", *page.ID)
	// AssertEqualsTime(t, time.)
	utils.AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *page.CreatedBy.ID)
	utils.AssertEqualsString(t, "61b898c0-7772-48d5-886e-2a222669737c", *page.LastEditedBy.ID)
	utils.AssertEqualsString(t, "external", *page.Cover.Type)
	utils.AssertEqualsString(t, "https://www.notion.so/images/page-cover/rijksmuseum_vermeer_the_milkmaid.jpg", *page.Cover.ExternalFile.URL)
	utils.AssertEqualsString(t, "emoji", *page.Icon.Type)
	utils.AssertEqualsString(t, "😁", *page.Icon.Emoji)
	utils.AssertEqualsString(t, "database_id", *page.Parent.Type)
	utils.AssertEqualsString(t, "ffd97167-8029-47f8-8623-c2347dd9c563", *page.Parent.DatabaseID)
	utils.AssertEqualsBool(t, false, *page.Archived)
	props := page.Properties.(notion.PageProperties)
	v, ok := props["Property 2"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "%3A%60gb", *v.ID)
	utils.AssertEqualsString(t, "date", *v.Type)
	utils.AssertEqualsString(t, "2022-04-04", *v.Date.Start)
	utils.AssertNil(t, v.Date.End)
	utils.AssertNil(t, v.Date.TimeZone)
	v, ok = props["Tags"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "%40Lx%7C", *v.ID)
	utils.AssertEqualsString(t, "multi_select", *v.Type)
	utils.AssertEqualsInt(t, 2, len(v.MultiSelect))
	utils.AssertEqualsString(t, "a0f30ec1-f592-4ec2-a69a-71fbd2eafbf3", *v.MultiSelect[0].ID)
	utils.AssertEqualsString(t, "word", *v.MultiSelect[0].Name)
	utils.AssertEqualsString(t, "blue", *v.MultiSelect[0].Color)
	utils.AssertEqualsString(t, "28d91a14-b38b-4d3c-afa6-738b3864ae0f", *v.MultiSelect[1].ID)
	utils.AssertEqualsString(t, "drow", *v.MultiSelect[1].Name)
	utils.AssertEqualsString(t, "purple", *v.MultiSelect[1].Color)
	v, ok = props["updated_at"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "Az%3Cf", *v.ID)
	utils.AssertEqualsString(t, "last_edited_time", *v.Type)
	// AssertEqualsTime(t, , v.LastEditedTime)
	v, ok = props["Property 6"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "KP%7CT", *v.ID)
	utils.AssertEqualsString(t, "url", *v.Type)
	utils.AssertEqualsString(t, "https://pkg.go.dev/github.com/asstart/english-words/app/ewords/notion#Filter.Property", *v.URL)
	v, ok = props["Defenition"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "OIab", *v.ID)
	utils.AssertEqualsString(t, "rich_text", *v.Type)
	utils.AssertEqualsInt(t, 0, len(v.RichText))
	v, ok = props["Property 13"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "Rl%3F%3D", *v.ID)
	utils.AssertEqualsString(t, "last_edited_by", *v.Type)
	utils.AssertEqualsString(t, "user", *v.LastEditedBy.Object)
	utils.AssertEqualsString(t, "61b898c0-7772-48d5-886e-2a222669737c", *v.LastEditedBy.ID)
	v, ok = props["Property"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "SUv~", *v.ID)
	utils.AssertEqualsString(t, "number", *v.Type)
	utils.AssertEqualsInt(t, 1234, int(*v.Number))

	v, ok = props["Transcription"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "TmLV", *v.ID)
	utils.AssertEqualsString(t, "rich_text", *v.Type)
	utils.AssertEqualsInt(t, 1, len(v.RichText))
	utils.AssertEqualsString(t, "text", *v.RichText[0].Type)
	utils.AssertEqualsString(t, "hello", *v.RichText[0].Text.Content)
	utils.AssertNil(t, v.RichText[0].Text.Link)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Bold)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Italic)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Strikethrough)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Underline)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Code)
	utils.AssertEqualsString(t, "default", *v.RichText[0].Annotations.Color)
	utils.AssertEqualsString(t, "hello", *v.RichText[0].PlainText)
	utils.AssertNil(t, v.RichText[0].Href)

	v, ok = props["Property 7"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "Trvl", *v.ID)
	utils.AssertEqualsString(t, "email", *v.Type)
	utils.AssertEqualsString(t, "some@gmail.com", *v.Email)

	v, ok = props["Property 3"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "%5EbRD", *v.ID)
	utils.AssertEqualsString(t, "people", *v.Type)
	utils.AssertEqualsInt(t, 1, len(v.People))
	utils.AssertEqualsString(t, "user", *v.People[0].Object)
	utils.AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *v.People[0].ID)

	v, ok = props["Property 12"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "_cww", *v.ID)
	utils.AssertEqualsString(t, "created_by", *v.Type)
	utils.AssertEqualsString(t, "user", *v.CreatedBy.Object)
	utils.AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *v.CreatedBy.ID)

	v, ok = props["Example"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "cjRY", *v.ID)
	utils.AssertEqualsString(t, "rich_text", *v.Type)
	utils.AssertEqualsInt(t, 1, len(v.RichText))
	utils.AssertEqualsString(t, "text", *v.RichText[0].Type)
	utils.AssertEqualsString(t, "Hello there", *v.RichText[0].Text.Content)
	utils.AssertNil(t, v.RichText[0].Text.Link)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Bold)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Italic)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Strikethrough)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Underline)
	utils.AssertEqualsBool(t, false, *v.RichText[0].Annotations.Code)
	utils.AssertEqualsString(t, "default", *v.RichText[0].Annotations.Color)
	utils.AssertEqualsString(t, "Hello there", *v.RichText[0].PlainText)
	utils.AssertNil(t, v.RichText[0].Href)

	v, ok = props["handled"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "il~C", *v.ID)
	utils.AssertEqualsString(t, "checkbox", *v.Type)
	utils.AssertEqualsBool(t, true, *v.Checkbox)

	v, ok = props["Property 5"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "m~bE", *v.ID)
	utils.AssertEqualsString(t, "checkbox", *v.Type)
	utils.AssertEqualsBool(t, false, *v.Checkbox)

	v, ok = props["created_at"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "o~QP", *v.ID)
	utils.AssertEqualsString(t, "created_time", *v.Type)
	//TODO check time

	v, ok = props["Property 9"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "u%3Bcf", *v.ID)
	utils.AssertEqualsString(t, "formula", *v.Type)
	utils.AssertEqualsString(t, "boolean", *v.Formula.Type)
	utils.AssertEqualsBool(t, false, *v.Formula.BooleanFormula)

	v, ok = props["Property 8"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "z%5D%3AD", *v.ID)
	utils.AssertEqualsString(t, "phone_number", *v.Type)
	utils.AssertEqualsString(t, "+79998887766", *v.PhoneNumber)

	v, ok = props["Word"]
	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, "title", *v.ID)
	utils.AssertEqualsString(t, "title", *v.Type)
	utils.AssertEqualsInt(t, 1, len(v.Title))
	utils.AssertEqualsString(t, "text", *v.Title[0].Type)
	utils.AssertEqualsString(t, "Hello", *v.Title[0].Text.Content)
	utils.AssertNil(t, v.Title[0].Text.Link)
	utils.AssertEqualsBool(t, false, *v.Title[0].Annotations.Bold)
	utils.AssertEqualsBool(t, false, *v.Title[0].Annotations.Italic)
	utils.AssertEqualsBool(t, false, *v.Title[0].Annotations.Strikethrough)
	utils.AssertEqualsBool(t, false, *v.Title[0].Annotations.Underline)
	utils.AssertEqualsBool(t, false, *v.Title[0].Annotations.Code)
	utils.AssertEqualsString(t, "default", *v.Title[0].Annotations.Color)
	utils.AssertEqualsString(t, "Hello", *v.Title[0].PlainText)
	utils.AssertNil(t, v.Title[0].Href)
}
