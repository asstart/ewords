// +build !integration

package ewords

import (
	"encoding/json"
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

	var text TextCondition
	json.Unmarshal([]byte(source), &text)

	AssertEqualsString(t, "equals_string", *text.Equals)
	AssertEqualsString(t, "not_equal_string", *text.DoesntEqual)
	AssertEqualsString(t, "contains_string", *text.Contains)
	AssertEqualsString(t, "does_not_contain_string", *text.DoesntContain)
	AssertEqualsString(t, "starts_with_string", *text.StartsWith)
	AssertEqualsString(t, "ends_with_string", *text.EndsWith)
	AssertEqualsBool(t, true, *text.IsEmpty)
	AssertEqualsBool(t, true, *text.IsNotEmpty)
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

	var pages DatabasePages
	json.Unmarshal([]byte(source), &pages)

	AssertEqualsBool(t, false, *pages.HasMore)
	AssertEqualsInt(t, 1, len(pages.Results))
	page := pages.Results[0]
	AssertEqualsString(t, "4f5cb747-e1d5-46a8-9bcd-48df3a21d675", *page.ID)
	// AssertEqualsTime(t, time.)
	AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *page.CreatedBy.ID)
	AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *page.LastEditedBy.ID)
	AssertEqualsString(t, "external", *page.Cover.Type)
	AssertEqualsString(t, "https://www.notion.so/images/page-cover/rijksmuseum_vermeer_the_milkmaid.jpg", *page.Cover.ExternalFile.URL)
	AssertEqualsString(t, "emoji", *page.Icon.Type)
	AssertEqualsString(t, "😁", *page.Icon.Emoji)
	AssertEqualsString(t, "database_id", *page.Parent.Type)
	AssertEqualsString(t, "ffd97167-8029-47f8-8623-c2347dd9c563", *page.Parent.DatabaseID)
	AssertEqualsBool(t, false, *page.Archived)
	props := page.Properties.(PageProperties)
	v, ok := props["Property 2"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "%3A%60gb", *v.ID)
	AssertEqualsString(t, "date", *v.Type)
	AssertEqualsString(t, "2022-04-04", *v.Date.Start)
	AssertNil(t, v.Date.End)
	AssertNil(t, v.Date.TimeZone)
	v, ok = props["Property 4"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "files", *v.Type)
	AssertEqualsInt(t, 1, len(v.Files))
	AssertEqualsString(t, "2.jpg", *v.Files[0].Name)
	AssertEqualsString(t, "file", *v.Files[0].Type)
	AssertEqualsString(
		t,
		"https://s3.us-west-2.amazonaws.com/secure.notion-static.com/8a85ee6e-dfa5-42b8-8900-9cb92d490d54/2.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220418%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220418T223419Z&X-Amz-Expires=3600&X-Amz-Signature=326c1cca24e649c1fb2861644212546f26ad2f11b7c4e72426a5d648ad868700&X-Amz-SignedHeaders=host&x-id=GetObject",
		*v.Files[0].NotionFile.URL)
	// AssertEqualsTime(t, time.Parse("2022-04-18T23:34:19.557Z"), *v.Files[0].NotionFile.ExpireTime)
	v, ok = props["Tags"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "%40Lx%7C", *v.ID)
	AssertEqualsString(t, "multi_select", *v.Type)
	AssertEqualsInt(t, 2, len(v.MultiSelect))
	AssertEqualsString(t, "a0f30ec1-f592-4ec2-a69a-71fbd2eafbf3", *v.MultiSelect[0].ID)
	AssertEqualsString(t, "word", *v.MultiSelect[0].Name)
	AssertEqualsString(t, "blue", *v.MultiSelect[0].Color)
	AssertEqualsString(t, "28d91a14-b38b-4d3c-afa6-738b3864ae0f", *v.MultiSelect[1].ID)
	AssertEqualsString(t, "drow", *v.MultiSelect[1].Name)
	AssertEqualsString(t, "purple", *v.MultiSelect[1].Color)
	v, ok = props["updated_at"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "Az%3Cf", *v.ID)
	AssertEqualsString(t, "last_edited_time", *v.Type)
	// AssertEqualsTime(t, , v.LastEditedTime)
	v, ok = props["Property 6"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "KP%7CT", *v.ID)
	AssertEqualsString(t, "url", *v.Type)
	AssertEqualsString(t, "https://pkg.go.dev/github.com/asstart/english-words/app/ewords/notion#Filter.Property", *v.URL)
	v, ok = props["Defenition"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "OIab", *v.ID)
	AssertEqualsString(t, "rich_text", *v.Type)
	AssertEqualsInt(t, 0, len(v.RichText))
	v, ok = props["Property 13"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "Rl%3F%3D", *v.ID)
	AssertEqualsString(t, "last_edited_by", *v.Type)
	AssertEqualsString(t, "user", *v.LastEditedBy.Object)
	AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *v.LastEditedBy.ID)
	v, ok = props["Property"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "SUv~", *v.ID)
	AssertEqualsString(t, "number", *v.Type)
	AssertEqualsInt(t, 1234, int(*v.Number))

	v, ok = props["Transcription"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "TmLV", *v.ID)
	AssertEqualsString(t, "rich_text", *v.Type)
	AssertEqualsInt(t, 1, len(v.RichText))
	AssertEqualsString(t, "text", *v.RichText[0].Type)
	AssertEqualsString(t, "hello", *v.RichText[0].Text.Content)
	AssertNil(t, v.RichText[0].Text.Link)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Bold)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Italic)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Strikethrough)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Underline)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Code)
	AssertEqualsString(t, "default", *v.RichText[0].Annotations.Color)
	AssertEqualsString(t, "hello", *v.RichText[0].PlainText)
	AssertNil(t, v.RichText[0].Href)

	v, ok = props["Property 7"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "Trvl", *v.ID)
	AssertEqualsString(t, "email", *v.Type)
	AssertEqualsString(t, "some@gmail.com", *v.Email)

	v, ok = props["Property 3"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "%5EbRD", *v.ID)
	AssertEqualsString(t, "people", *v.Type)
	AssertEqualsInt(t, 1, len(v.People))
	AssertEqualsString(t, "user", *v.People[0].Object)
	AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *v.People[0].ID)

	v, ok = props["Property 12"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "_cww", *v.ID)
	AssertEqualsString(t, "created_by", *v.Type)
	AssertEqualsString(t, "user", *v.CreatedBy.Object)
	AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *v.CreatedBy.ID)

	v, ok = props["Example"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "cjRY", *v.ID)
	AssertEqualsString(t, "rich_text", *v.Type)
	AssertEqualsInt(t, 1, len(v.RichText))
	AssertEqualsString(t, "text", *v.RichText[0].Type)
	AssertEqualsString(t, "Hello there", *v.RichText[0].Text.Content)
	AssertNil(t, v.RichText[0].Text.Link)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Bold)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Italic)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Strikethrough)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Underline)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Code)
	AssertEqualsString(t, "default", *v.RichText[0].Annotations.Color)
	AssertEqualsString(t, "Hello there", *v.RichText[0].PlainText)
	AssertNil(t, v.RichText[0].Href)

	v, ok = props["handled"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "il~C", *v.ID)
	AssertEqualsString(t, "checkbox", *v.Type)
	AssertEqualsBool(t, true, *v.Checkbox)

	v, ok = props["Property 5"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "m~bE", *v.ID)
	AssertEqualsString(t, "checkbox", *v.Type)
	AssertEqualsBool(t, false, *v.Checkbox)

	v, ok = props["created_at"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "o~QP", *v.ID)
	AssertEqualsString(t, "created_time", *v.Type)
	//TODO check time

	v, ok = props["Property 9"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "u%3Bcf", *v.ID)
	AssertEqualsString(t, "formula", *v.Type)
	AssertEqualsString(t, "boolean", *v.Formula.Type)
	AssertEqualsBool(t, false, *v.Formula.BooleanFormula)

	v, ok = props["Property 8"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "z%5D%3AD", *v.ID)
	AssertEqualsString(t, "phone_number", *v.Type)
	AssertEqualsString(t, "+79998887766", *v.PhoneNumber)

	v, ok = props["Word"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "title", *v.ID)
	AssertEqualsString(t, "title", *v.Type)
	AssertEqualsInt(t, 1, len(v.Title))
	AssertEqualsString(t, "text", *v.Title[0].Type)
	AssertEqualsString(t, "Hello", *v.Title[0].Text.Content)
	AssertNil(t, v.Title[0].Text.Link)
	AssertEqualsBool(t, false, *v.Title[0].Annotations.Bold)
	AssertEqualsBool(t, false, *v.Title[0].Annotations.Italic)
	AssertEqualsBool(t, false, *v.Title[0].Annotations.Strikethrough)
	AssertEqualsBool(t, false, *v.Title[0].Annotations.Underline)
	AssertEqualsBool(t, false, *v.Title[0].Annotations.Code)
	AssertEqualsString(t, "default", *v.Title[0].Annotations.Color)
	AssertEqualsString(t, "Hello", *v.Title[0].PlainText)
	AssertNil(t, v.Title[0].Href)
}

func TestMarshalRichTextContainsFilter(t *testing.T) {
	var f = Filter{
		Property: StrPtr("Landmark"),
		RichText: &TextCondition{

			Contains: StrPtr("Bridge"),
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

	AssertEqualsString(t, Minimise(expected), Minimise(string(marshaled)))
}

func TestMarshalCompoundFilter(t *testing.T) {
	var f = Filter{
		And: []Filter{
			Filter{
				Property: StrPtr("Seen"),
				CheckBox: &CheckboxCondition{
					Equals: BoolPtr(false),
				},
			},
			Filter{
				Property: StrPtr("Yearly visitor count"),
				Number: &NumberCondition{
					GreaterThan: FloatPtr(1000000),
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

	AssertEqualsString(t, Minimise(expected), Minimise(string(marshaled)))
}

func TestMarshalMiltililevelCompoundFilter(t *testing.T) {
	var f = Filter{
		Or: []Filter{
			Filter{
				Property: StrPtr("Description"),
				RichText: &TextCondition{
					Contains: StrPtr("fish"),
				},
			},
			Filter{
				And: []Filter{
					Filter{
						Property: StrPtr("Food group"),
						Select: &SelectCondition{
							Equals: StrPtr("🥦Vegetable"),
						},
					},
					Filter{
						Property: StrPtr("Is protein rich?"),
						CheckBox: &CheckboxCondition{
							Equals: BoolPtr(true),
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

	AssertEqualsString(t, Minimise(expected), Minimise(string(marshaled)))
}

func TestSortMarshal(t *testing.T) {
	s := Sort{
		Property:  StrPtr("Food group"),
		Direction: StrPtr("descending"),
	}

	expected := `
	{
		"property": "Food group",
		"direction": "descending"
	}
	`

	marshaled, _ := json.Marshal(s)

	AssertEqualsString(t, Minimise(expected), Minimise(string(marshaled)))
}

func TestDatabaseQuery(t *testing.T) {
	query := DatabaseQuery{
		Filter: &Filter{
			Property: StrPtr("Landmark"),
			RichText: &TextCondition{

				Contains: StrPtr("Bridge"),
			},
		},
		Sorts: []Sort{
			Sort{
				Property:  StrPtr("Food group"),
				Direction: StrPtr("descending"),
			},
			Sort{
				Property:  StrPtr("Name"),
				Direction: StrPtr("ascending"),
			},
		},
		StartCursor: StrPtr("3-295-0235"),
		PageSize:    IntPtr(50),
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

	AssertEqualsString(t, Minimise(expected), Minimise(string(marshaled)))
}

func TestUpdatePage(t *testing.T) {
	up := UpdatePage{
		Properties: PageProperties{
			"handled": PageProperty{
				Checkbox: BoolPtr(true),
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

	AssertEqualsString(t, Minimise(expected), Minimise(string(marshaled)))
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

	var page Page
	json.Unmarshal([]byte(jsn), &page)

	AssertEqualsString(t, "4f5cb747-e1d5-46a8-9bcd-48df3a21d675", *page.ID)
	// AssertEqualsTime(t, time.)
	AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *page.CreatedBy.ID)
	AssertEqualsString(t, "61b898c0-7772-48d5-886e-2a222669737c", *page.LastEditedBy.ID)
	AssertEqualsString(t, "external", *page.Cover.Type)
	AssertEqualsString(t, "https://www.notion.so/images/page-cover/rijksmuseum_vermeer_the_milkmaid.jpg", *page.Cover.ExternalFile.URL)
	AssertEqualsString(t, "emoji", *page.Icon.Type)
	AssertEqualsString(t, "😁", *page.Icon.Emoji)
	AssertEqualsString(t, "database_id", *page.Parent.Type)
	AssertEqualsString(t, "ffd97167-8029-47f8-8623-c2347dd9c563", *page.Parent.DatabaseID)
	AssertEqualsBool(t, false, *page.Archived)
	props := page.Properties.(PageProperties)
	v, ok := props["Property 2"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "%3A%60gb", *v.ID)
	AssertEqualsString(t, "date", *v.Type)
	AssertEqualsString(t, "2022-04-04", *v.Date.Start)
	AssertNil(t, v.Date.End)
	AssertNil(t, v.Date.TimeZone)
	v, ok = props["Tags"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "%40Lx%7C", *v.ID)
	AssertEqualsString(t, "multi_select", *v.Type)
	AssertEqualsInt(t, 2, len(v.MultiSelect))
	AssertEqualsString(t, "a0f30ec1-f592-4ec2-a69a-71fbd2eafbf3", *v.MultiSelect[0].ID)
	AssertEqualsString(t, "word", *v.MultiSelect[0].Name)
	AssertEqualsString(t, "blue", *v.MultiSelect[0].Color)
	AssertEqualsString(t, "28d91a14-b38b-4d3c-afa6-738b3864ae0f", *v.MultiSelect[1].ID)
	AssertEqualsString(t, "drow", *v.MultiSelect[1].Name)
	AssertEqualsString(t, "purple", *v.MultiSelect[1].Color)
	v, ok = props["updated_at"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "Az%3Cf", *v.ID)
	AssertEqualsString(t, "last_edited_time", *v.Type)
	// AssertEqualsTime(t, , v.LastEditedTime)
	v, ok = props["Property 6"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "KP%7CT", *v.ID)
	AssertEqualsString(t, "url", *v.Type)
	AssertEqualsString(t, "https://pkg.go.dev/github.com/asstart/english-words/app/ewords/notion#Filter.Property", *v.URL)
	v, ok = props["Defenition"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "OIab", *v.ID)
	AssertEqualsString(t, "rich_text", *v.Type)
	AssertEqualsInt(t, 0, len(v.RichText))
	v, ok = props["Property 13"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "Rl%3F%3D", *v.ID)
	AssertEqualsString(t, "last_edited_by", *v.Type)
	AssertEqualsString(t, "user", *v.LastEditedBy.Object)
	AssertEqualsString(t, "61b898c0-7772-48d5-886e-2a222669737c", *v.LastEditedBy.ID)
	v, ok = props["Property"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "SUv~", *v.ID)
	AssertEqualsString(t, "number", *v.Type)
	AssertEqualsInt(t, 1234, int(*v.Number))

	v, ok = props["Transcription"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "TmLV", *v.ID)
	AssertEqualsString(t, "rich_text", *v.Type)
	AssertEqualsInt(t, 1, len(v.RichText))
	AssertEqualsString(t, "text", *v.RichText[0].Type)
	AssertEqualsString(t, "hello", *v.RichText[0].Text.Content)
	AssertNil(t, v.RichText[0].Text.Link)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Bold)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Italic)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Strikethrough)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Underline)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Code)
	AssertEqualsString(t, "default", *v.RichText[0].Annotations.Color)
	AssertEqualsString(t, "hello", *v.RichText[0].PlainText)
	AssertNil(t, v.RichText[0].Href)

	v, ok = props["Property 7"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "Trvl", *v.ID)
	AssertEqualsString(t, "email", *v.Type)
	AssertEqualsString(t, "some@gmail.com", *v.Email)

	v, ok = props["Property 3"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "%5EbRD", *v.ID)
	AssertEqualsString(t, "people", *v.Type)
	AssertEqualsInt(t, 1, len(v.People))
	AssertEqualsString(t, "user", *v.People[0].Object)
	AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *v.People[0].ID)

	v, ok = props["Property 12"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "_cww", *v.ID)
	AssertEqualsString(t, "created_by", *v.Type)
	AssertEqualsString(t, "user", *v.CreatedBy.Object)
	AssertEqualsString(t, "c7f2ae70-6b98-438f-8564-c59a71d7b3a4", *v.CreatedBy.ID)

	v, ok = props["Example"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "cjRY", *v.ID)
	AssertEqualsString(t, "rich_text", *v.Type)
	AssertEqualsInt(t, 1, len(v.RichText))
	AssertEqualsString(t, "text", *v.RichText[0].Type)
	AssertEqualsString(t, "Hello there", *v.RichText[0].Text.Content)
	AssertNil(t, v.RichText[0].Text.Link)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Bold)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Italic)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Strikethrough)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Underline)
	AssertEqualsBool(t, false, *v.RichText[0].Annotations.Code)
	AssertEqualsString(t, "default", *v.RichText[0].Annotations.Color)
	AssertEqualsString(t, "Hello there", *v.RichText[0].PlainText)
	AssertNil(t, v.RichText[0].Href)

	v, ok = props["handled"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "il~C", *v.ID)
	AssertEqualsString(t, "checkbox", *v.Type)
	AssertEqualsBool(t, true, *v.Checkbox)

	v, ok = props["Property 5"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "m~bE", *v.ID)
	AssertEqualsString(t, "checkbox", *v.Type)
	AssertEqualsBool(t, false, *v.Checkbox)

	v, ok = props["created_at"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "o~QP", *v.ID)
	AssertEqualsString(t, "created_time", *v.Type)
	//TODO check time

	v, ok = props["Property 9"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "u%3Bcf", *v.ID)
	AssertEqualsString(t, "formula", *v.Type)
	AssertEqualsString(t, "boolean", *v.Formula.Type)
	AssertEqualsBool(t, false, *v.Formula.BooleanFormula)

	v, ok = props["Property 8"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "z%5D%3AD", *v.ID)
	AssertEqualsString(t, "phone_number", *v.Type)
	AssertEqualsString(t, "+79998887766", *v.PhoneNumber)

	v, ok = props["Word"]
	AssertEqualsBool(t, true, ok)
	AssertEqualsString(t, "title", *v.ID)
	AssertEqualsString(t, "title", *v.Type)
	AssertEqualsInt(t, 1, len(v.Title))
	AssertEqualsString(t, "text", *v.Title[0].Type)
	AssertEqualsString(t, "Hello", *v.Title[0].Text.Content)
	AssertNil(t, v.Title[0].Text.Link)
	AssertEqualsBool(t, false, *v.Title[0].Annotations.Bold)
	AssertEqualsBool(t, false, *v.Title[0].Annotations.Italic)
	AssertEqualsBool(t, false, *v.Title[0].Annotations.Strikethrough)
	AssertEqualsBool(t, false, *v.Title[0].Annotations.Underline)
	AssertEqualsBool(t, false, *v.Title[0].Annotations.Code)
	AssertEqualsString(t, "default", *v.Title[0].Annotations.Color)
	AssertEqualsString(t, "Hello", *v.Title[0].PlainText)
	AssertNil(t, v.Title[0].Href)
}