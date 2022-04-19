package ewords

import (
	"encoding/json"
	"time"
)

type DatabasePages struct {
	Results    []Page  `json:"results"`
	NextCursor *string `json:"next_cursor"`
	HasMore    *bool   `json:"has_more"`
}

type Page struct {
	ID             *string     `json:"id"`
	CreatedTime    *time.Time  `json:"created_time"`
	CreatedBy      *User       `json:"created_by"`
	LastEditedTime *time.Time  `json:"last_edited_time"`
	LastEditedBy   *User       `json:"last_edited_by"`
	Parent         *Parent     `json:"parent"`
	Icon           *Emoji      `json:"icon,omitempty"`
	Cover          *File       `json:"cover,omitempty"`
	Archived       *bool       `json:"archived"`
	Properties     interface{} `json:"properties"`
	URL            *string     `json:"url"`
}

type Parent struct {
	Type       *string `json:"type"`
	DatabaseID *string `json:"database_id,omitempty"`
	PageID     *string `json:"page_id,omitempty"`
	Workspace  *bool   `json:"workspace,omitempty"`
}

type PageProperty struct {
	Title *string `json:"tile,omitempty"`
}

type DatabaseProperties map[string]DatabaseProperty

type DatabaseProperty struct {
	ID   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`

	Title          []RichText     `json:"title,omitempty"`
	RichText       []RichText     `json:"rich_text,omitempty"`
	Number         *float64       `json:"number,omitempty"`
	Select         *SelectOption  `json:"select,omitempty"`
	MultiSelect    []SelectOption `json:"multi_select,omitempty"`
	Date           *DateProperty  `json:"date,omitempty"`
	People         []User         `json:"people,omitempty"`
	Files          []File         `json:"files,omitempty"`
	Checkbox       *bool          `json:"checkbox,omitempty"`
	URL            *string        `json:"url,omitempty"`
	Email          *string        `json:"email,omitempty"`
	PhoneNumber    *string        `json:"phone_number,omitempty"`
	Formula        *Formula       `json:"formula,omitempty"`
	Relation       *Relation      `json:"relation,omitempty"`
	Rollup         *Rollup        `json:"rollup,omitempty"`
	CreatedTime    *time.Time     `json:"created_time,omitempty"`
	CreatedBy      *User          `json:"created_by,omitempty"`
	LastEditedTime *time.Time     `json:"last_edited_time,omitempty"`
	LastEditedBy   *User          `json:"last_edited_by,omitempty"`
}

//TODO Check Start,End correctness
type DateProperty struct {
	Start    *string `json:"start"`
	End      *string `json:"end"`
	TimeZone *string `json:"time_zone"`
}

type SelectOption struct {
	Name  *string `json:"name,omitempty"`
	ID    *string `json:"id,omitempty"`
	Color *string `json:"color,omitempty"`
}

type Formula struct {
	Type           *string    `json:"type"`
	StringFormula  *string    `json:"string,omitempty"`
	BooleanFormula *bool      `json:"boolean,omitempty"`
	NumberFormula  *float64   `json:"number,omitempty"`
	DateFormula    *time.Time `json:"date,omitempty"`
	// Expression *string `json:"expression"`
}

type Relation struct {
	DatabaseID         *string `json:"database_id"`
	SyncedPropertyName *string `json:"synced_property_name"`
	SyncedPropertyID   *string `json:"synced_property_id"`
}

type Rollup struct {
	RelationPropertyName string `json:"relation_property_name"`
	RelationPropertyID   string `json:"relation_property_id"`
	RollupPropertyName   string `json:"rollup_property_name"`
	RollupPropertyID     string `json:"rollup_property_id"`
	Function             string `json:"function"`
}

type RichText struct {
	Type        *string      `json:"type"`
	PlainText   *string      `json:"plain_text"`
	Href        *string      `json:"href"`
	Annotations *Annotations `json:"annotations"`
	Text        *Text        `json:"text,omitempty"`
	Mention     *Mention     `json:"mention,omitempty"`
	Equation    *Equation    `json:"equation,omitempty"`
}

type Text struct {
	Content *string `json:"content"`
	Link    *Link   `json:"link,omitempty"`
}

type Mention struct {
	MentionType *string `json:"type"`

	User     *string    `json:"user,omitempty"`
	Page     *string    `json:"page,omitempty"`
	Database *string    `json:"database,omitempty"`
	Date     *time.Time `json:"date,omitempty"`
	Template *Template  `json:"template_mention,omitempty"`
}

type Equation struct {
	Expression *string `json:"expression"`
}

type Template struct {
	TemplateMentionDate *string `json:"template_mention_date"`
	TemplateMentionUser *string `json:"template_mention_user"`
}

type Link struct {
	URL *string `json:"url"`
}

type Annotations struct {
	Bold          *bool   `json:"bold,omitempty"`
	Italic        *bool   `json:"italic,omitempty"`
	Strikethrough *bool   `json:"strikethrough,omitempty"`
	Underline     *bool   `json:"underline,omitempty"`
	Code          *bool   `json:"code,omitempty"`
	Color         *string `json:"color,omitempty"`
}

type User struct {
	ID        *string `json:"id"`
	Object    *string `json:"object,omitempty"`
	Type      *string `json:"type,omitempty"`
	Name      *string `json:"name,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Person    *Person `json:"person,omitempty"`
	Bot       *Bot    `json:"bor,omitempty"`
}

type Person struct {
	Email *string `json:"email"`
}

type Bot struct {
	Owner *BotOwner `json:"owner"`
}

type BotOwner struct {
	Type      *string `json:"type"`
	Workspace *bool   `json:"workspace"`
	Owner     *User   `json:"user"`
}

type Emoji struct {
	Type  *string `json:"type"`
	Emoji *string `json:"emoji"`
}

type File struct {
	Type         *string       `json:"type,omitempty"`
	Name         *string       `json:"name,omitempty"`
	ExternalFile *ExternalFile `json:"external,omitempty"`
	NotionFile   *NotionFile   `json:"file,omitempty"`
}

type ExternalFile struct {
	URL *string `json:"url"`
}

type NotionFile struct {
	URL        *string    `json:"url"`
	ExpireTime *time.Time `json:"expire_time,omitempty"`
}

type DatabaseQuery struct {
	Filter      *Filter `json:"filter,omitempty"`
	Sorts       []Sort  `json:"sorts,omitempty"`
	StartCursor *string `json:"start_cursor,omitempty"`
	PageSize    *int    `json:"page_size,omitempty"`
}

type Sort struct {
	Property  *string `json:"property"`
	Direction *string `json:"direction,omitempty"`
}

type Filter struct {
	Property *string `json:"property,omitempty"`

	RichText    *TextCondition     `json:"rich_text,omitempty"`
	Title       *TextCondition     `json:"title,omitempty"`
	URL         *TextCondition     `json:"url,omitempty"`
	Email       *TextCondition     `json:"email,omitempty"`
	PhoneNumber *TextCondition     `json:"phone_number,omitempty"`
	Number      *NumberCondition   `json:"number,omitempty"`
	CheckBox    *CheckboxCondition `json:"checkbox,omitempty"`
	Select      *SelectCondition   `json:"select,omitempty"`
	MultiSelect *SelectCondition   `json:"multi_select,omitempty"`
	File        *FileCondition     `json:"files,omitempty"`
	Relation    *RelationCondition `json:"relation,omitempty"`
	Date        *DateCondition     `json:"date,omitempty"`

	And []Filter `json:"and,omitempty"`
	Or  []Filter `json:"or,omitempty"`
}

type TextCondition struct {
	Equals        *string `json:"equals,omitempty"`
	DoesntEqual   *string `json:"does_not_equal,omitempty"`
	Contains      *string `json:"contains,omitempty"`
	DoesntContain *string `json:"does_not_contain,omitempty"`
	StartsWith    *string `json:"starts_with,omitempty"`
	EndsWith      *string `json:"ends_with,omitempty"`
	IsEmpty       *bool   `json:"is_empty,omitempty"`
	IsNotEmpty    *bool   `json:"is_not_empty,omitempty"`
}

type NumberCondition struct {
	Equals               *float64 `json:"equals,omitempty"`
	DoesntEqual          *float64 `json:"does_not_equal,omitempty"`
	GreaterThan          *float64 `json:"greater_than,omitempty"`
	LessThan             *float64 `json:"less_than,omitempty"`
	GreaterThanOrEqualTo *float64 `json:"greater_than_or_equal_to,omitempty"`
	LessThanOrEqualTo    *float64 `json:"less_than_or_equals_to,omitempty"`
	IsEmpty              *bool    `json:"is_empty,omitempty"`
	IsNotEmpty           *bool    `json:"is_not_empty,omitempty"`
}

type CheckboxCondition struct {
	Equals      *bool `json:"equals,omitempty"`
	DoesntEqual *bool `json:"does_not_equal,omitempty"`
}

type SelectCondition struct {
	Equals      *string `json:"equals,omitempty"`
	DoesntEqual *string `json:"does_not_equal,omitempty"`
	IsEmpty     *bool   `json:"is_empty,omitempty"`
	IsNotEmpty  *bool   `json:"is_not_empty,omitempty"`
}

type PeopleCondition struct {
	IsEmpty        *bool   `json:"is_empty,omitempty"`
	IsNotEmpty     *bool   `json:"is_not_empty,omitempty"`
	Contains       *string `json:"contains,omitempty"`
	DoesntContains *string `json:"does_not_contain,omitempty"`
}

type FileCondition struct {
	IsEmpty    *bool `json:"is_empty,omitempty"`
	IsNotEmpty *bool `json:"is_not_empty,omitempty"`
}

type RelationCondition struct {
	IsEmpty        *bool   `json:"is_empty,omitempty"`
	IsNotEmpty     *bool   `json:"is_not_empty,omitempty"`
	Contains       *string `json:"contains,omitempty"`
	DoesntContains *string `json:"does_not_contain,omitempty"`
}

type TimestampFilterCondition struct {
	Timestamp *time.Time `json:"timestamp"`

	Date *DateCondition `json:"date,omitempty"`
}

type DateCondition struct {
	Equals     *time.Time  `json:"equals,omitempty"`
	Before     *time.Time  `json:"before,omitempty"`
	After      *time.Time  `json:"after,omitempty"`
	OnOrBefore *time.Time  `json:"on_or_before,omitempty"`
	OnOrAfter  *time.Time  `json:"on_or_after,omitempty"`
	PastWeek   interface{} `json:"past_week,omitempty"`
	PastMonth  interface{} `json:"past_month,omitempty"`
	PastYear   interface{} `json:"past_year,omitempty"`
	NextWeek   interface{} `json:"next_week,omitempty"`
	NextMonth  interface{} `json:"next_month,omitempty"`
	NextYear   interface{} `json:"next_year,omitempty"`
	IsEmpty    *bool       `json:"is_empty,omitempty"`
	IsNotEmpty *bool       `json:"is_not_empty,omitempty"`
}

func (p *Page) UnmarshalJSON(b []byte) error {
	type (
		P   Page
		DTO struct {
			P
			Properties json.RawMessage `json:"properties"`
		}
	)

	var dto DTO

	if err := json.Unmarshal(b, &dto); err != nil {
		return err
	}

	page := dto.P

	var props DatabaseProperties

	if err := json.Unmarshal(dto.Properties, &props); err != nil {
		return err
	}

	page.Properties = props

	*p = Page(page)
	return nil
}
