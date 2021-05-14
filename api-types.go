package notion

import "encoding/json"

// tries to follow https://github.com/makenotion/notion-sdk-js/blob/main/src/api-types.ts

type PaginationParameters struct {
	StartCursor string
	PageSize    int
}

type PaginatedList struct {
	Object     string                 `json:"object"` // 'list',
	Results    []APISingularObject    `json:"-"`
	ResultsRaw json.RawMessage        `json:"results"`
	HasMore    bool                   `json:"has_more"`
	NextCursor string                 `json:"next_cursor"`
	RawJSON    map[string]interface{} `json:"-"`
}

// can be: Database, Page, User, Block
type APISingularObject interface {
}

// can be: APISingularObject, PaginatedList
type APIObject interface {
}

// can be: ParagraphBlock, HeadingOneBlock,  HeadingTwoBlock, HeadingThreeBlock
// BulletedListItemBlock, NumberedListItemBlock, ToDoBlock, ToggleBlock, ChildPageBlock
// UnsupportedBlock
type Block interface {
	// TODO: possibly add Type() or BlockType() method that returnsType
	// (instead of having Type in BlockBase)
}

type BlockBase struct {
	Object         string `json:"object"` // 'block';
	ID             string `json:"id"`
	Type           string `json:"type"`
	CreatedTime    string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
	HasChildren    bool   `json:"has_children"`
}

type ParagraphBlock struct {
	BlockBase
	Type     string // 'paragraph';
	Text     []RichText
	Children []BlockBase
}

type HeadingOneBlock struct {
	BlockBase
	Type        string //: 'heading_1'
	Text        []RichText
	HasChildren bool //false;
}

type HeadingTwoBlock struct {
	BlockBase
	Type        string //'heading_2';
	Text        []RichText
	HasChildren bool //false;
}

type HeadingThreeBlock struct {
	BlockBase
	Type        string //'heading_3';
	Text        []RichText
	HasChildren bool //false;
}

type BulletedListItemBlock struct {
	BlockBase
	Type     string //'bulleted_list_item';
	Text     []RichText
	Children []Block
}

type NumberedListItemBlock struct {
	BlockBase
	Type     string //'numbered_list_item';
	Text     []RichText
	Children []Block
}

type ToDoBlock struct {
	BlockBase
	Type     string //'to_do';
	Text     []RichText
	Checked  bool
	Children []Block
}

type ToggleBlock struct {
	BlockBase
	Type     string //'toggle';
	Text     []RichText
	Children []Block
}

type ChildPageBlock struct {
	BlockBase
	Type  string //'child_page';
	Title string
}

type UnsupportedBlock struct {
	BlockBase
	Type string //'unsupported';
}

/*
export type Property =
  | TitleProperty
  | RichTextProperty
  | NumberProperty
  | SelectProperty
  | MultiSelectProperty
  | DateProperty
  | PeopleProperty
  | FileProperty
  | CheckboxProperty
  | URLProperty
  | EmailProperty
  | PhoneNumberProperty
  | FormulaProperty
  | RelationProperty
  | RollupProperty
  | CreatedTimeProperty
  | CreatedByProperty
  | LastEditedTimeProperty
  | LastEditedByProperty;
*/
type Property interface {
}

type PropertyBase struct {
	ID   string
	Type string
}

// TODO: this is meant to represent Record<string, never>
// which is { 'key': value }, where 'value' is of type 'never'
type RecordString string

type TitleProperty struct {
	PropertyBase
	Type  string //'title';
	Title RecordString
}

type RichTextProperty struct {
	PropertyBase
	Type     string //'rich_text';
	RichText RecordString
}

type NumberProperty struct {
	PropertyBase
	Type   string //'number';
	Number struct {
		Format string // 'number' | 'number_with_commas' | 'percent' | 'dollar' | 'euro' | 'pound' | 'yen' | 'ruble' | 'rupee' | 'won' | 'yuan';
	}
}

type SelectProperty struct {
	PropertyBase
	Type   string //'select';
	Select []SelectOption
}

type MultiSelectProperty struct {
	PropertyBase
	Type        string //'multi_select';
	MultiSelect []MultiSelectOption
}

type DateProperty struct {
	PropertyBase
	Type string //'date';
	Date RecordString
}

type PeopleProperty struct {
	PropertyBase
	Type   string //'people';
	People RecordString
}

type FileProperty struct {
	PropertyBase
	Type string //'file';
	File RecordString
}

type CheckboxProperty struct {
	PropertyBase
	Type     string //'checkbox';
	Checkbox RecordString
}

type URLProperty struct {
	PropertyBase
	Type string //'url';
	URL  RecordString
}

type EmailProperty struct {
	PropertyBase
	Type  string //'email';
	Email RecordString
}

type PhoneNumberProperty struct {
	PropertyBase
	Type        string //'phone_number';
	PhoneNumber RecordString
}

type FormulaProperty struct {
	PropertyBase
	Type    string //'formula';
	Formula struct {
		Expression string
	}
}

type RelationProperty struct {
	PropertyBase
	Type     string //'relation';
	Relation struct {
		DatabaseID         string
		SyncedPropertyName string
		SyncedPropertyID   string
	}
}

type RollupProperty struct {
	PropertyBase
	Type   string //'rollup';
	Rollup struct {
		Relation_property_name string
		Relation_property_id   string
		Rollup_property_name   string
		Rollup_property_id     string
		Function               string // 'count_all' | 'count_values' | 'count_unique_values' | 'count_empty' | 'count_not_empty' | 'percent_empty' | 'percent_not_empty' | 'sum' | 'average' | 'median' | 'min' | 'max' | 'range';
	}
}

type CreatedTimeProperty struct {
	PropertyBase
	Type        string //'created_time';
	CreatedTime RecordString
}

type CreatedByProperty struct {
	PropertyBase
	Type      string //'created_by';
	CreatedBy RecordString
}

type LastEditedTimeProperty struct {
	PropertyBase
	Type           string //'last_edited_time';
	LastEditedTime RecordString
}

type LastEditedByProperty struct {
	PropertyBase
	Type         string //'last_edited_by';
	LastEditedBy RecordString
}

/*
 * User (output)
 */

// PersonUser | BotUser
type User interface {
	Type() string
}

type UserBase struct {
	Object     string // 'user';
	ID         string
	Type       string
	Name       string
	Avatar_url string
}

type PersonUser struct {
	UserBase
	//Type 'person';
	Person struct {
		Email string
	}
}

type BotUser struct {
	UserBase
	//Type 'bot';
}

/*
* Misc (output)
 */

type SelectOption struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Color Color  `json:"color"`
}

type MultiSelectOption struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Color Color  `json:"color"`
}

type SearchSort struct {
	Direction string `json:"direction"` // 'ascending' | 'descending';
	Timestamp string `json:"timestamp"` // 'last_edited_time';
}

type SearchFilter struct {
	Value    string `json:"value"`    // 'object',
	Property string `json:"property"` // 'object';
}

type Color string // 'default' | 'gray' | 'brown' | 'orange' | 'yellow' | 'green' | 'blue' | 'purple' | 'pink' | 'red';

type BackgroundColor string // 'gray_background' | 'brown_background' | 'orange_background' | 'yellow_background' | 'green_background' | 'blue_background' | 'purple_background' | 'pink_background' | 'red_background';

/*
* Filter (input)
 */

//  SinglePropertyFilter | CompoundFilter
type Filter interface{}

/*
 export type SinglePropertyFilter =
	 | TextFilter
	 | NumberFilter
	 | CheckboxFilter
	 | SelectFilter
	 | MultiSelectFilter
	 | DateFilter
	 | PeopleFilter
	 | FilesFilter
	 | RelationFilter
	 | FormulaFilter
*/
type SinglePropertyFilter interface{}

/*
 export interface CompoundFilter {
	 or?: Filter[];
	 and?: Filter[];
 }

 export interface SinglePropertyFilterBase {
	 property: string;
 }

 export interface TextFilter extends SinglePropertyFilterBase {
	 equals?: string;
	 does_not_equal?: string;
	 contains?: string;
	 does_not_contain?: string;
	 starts_with?: string;
	 ends_with?: string;
	 is_empty?: true;
	 is_not_empty?: true;
 }

 export interface NumberFilter extends SinglePropertyFilterBase {
	 equals?: number;
	 does_not_equal?: number;
	 greater_than?: number;
	 less_than?: number;
	 greater_than_or_equal_to?: number;
	 less_than_or_equal_to?: number;
	 is_empty?: true;
	 is_not_empty?: true;
 }

 export interface CheckboxFilter extends SinglePropertyFilterBase {
	 equals?: bool
	 does_not_equal?: bool
 }

 export interface SelectFilter extends SinglePropertyFilterBase {
	 equals?: string;
	 does_not_equal?: string;
	 is_empty?: true;
	 is_not_empty?: true;
 }

 export interface MultiSelectFilter extends SinglePropertyFilterBase {
	 contains?: string;
	 does_not_contain?: string;
	 is_empty?: true;
	 is_not_empty?: true;
 }

 export interface DateFilter extends SinglePropertyFilterBase {
	 equals?: string;
	 before?: string;
	 after?: string;
	 on_or_before?: string;
	 is_empty?: true;
	 is_not_empty?: true;
	 on_or_after?: string;
	 past_week: Record<string, never>;
	 past_month: Record<string, never>;
	 past_year: Record<string, never>;
	 next_week: Record<string, never>;
	 next_month: Record<string, never>;
	 next_year: Record<string, never>;
 }

 export interface PeopleFilter extends SinglePropertyFilterBase {
	 contains?: string;
	 does_not_contain?: string;
	 is_empty?: true;
	 is_not_empty?: true;
 }

 export interface FilesFilter extends SinglePropertyFilterBase {
	 is_empty?: true;
	 is_not_empty?: true;
 }

 export interface RelationFilter extends SinglePropertyFilterBase {
	 contains?: string;
	 does_not_contain?: string;
	 is_empty?: true;
	 is_not_empty?: true;
 }

 export interface FormulaFilter extends SinglePropertyFilterBase {
	 text?: Omit<TextFilter, 'property'>;
	 checkbox?: Omit<CheckboxFilter, 'property'>;
	 number?: Omit<NumberFilter, 'property'>;
	 date?: Omit<DateFilter, 'property'>;
 }

*/

/*
* Sort (input)
 */

/*
 export interface Sort {
	 property?: string;
	 timestamp?: 'created_time' | 'last_edited_time';
	 direction: 'ascending' | 'descending';
 }
*/

/*
* Page
 */

type Page struct {
	Object string `json:"object"` // 'page',
	ID     string `json:"id"`
	Parent Parent `json:"parent"`
	// TODO: fix me
	//properties: { [propertyName: string]: PropertyValue };
	Properties map[string]PropertyValue `json:"properties"`
}

/*
* Parent
 */

// DatabaseParent | PageParent | WorkspaceParent
type Parent interface {
}

type DatabaseParent struct {
	Type       string `json:"type"` // 'database_id';
	DatabaseID string `json:"database_id"`
}

type PageParent struct {
	Type   string `json:"type"` //: 'page_id';
	PageID string `json:"page_id"`
}

type WorkspaceParent struct {
	Type string `json:"type"` //'workspace';
}

/*
 export type ParentInput = Omit<DatabaseParent, 'type'> |  Omit<PageParent, 'type'>;
*/

/*
* PropertyValue
 */

/*
 export type PropertyValue =
	| TitlePropertyValue
	| RichTextPropertyValue
	| NumberPropertyValue
	| SelectPropertyValue
	| MultiSelectPropertyValue
	| DatePropertyValue
	| FormulaPropertyValue
	| RollupPropertyValue
	| PeoplePropertyValue
	| FilesPropertyValue
	| CheckboxPropertyValue
	| URLPropertyValue
	| EmailPropertyValue
	| PhoneNumberPropertyValue
	| CreatedTimePropertyValue
	| CreatedByPropertyValue
	| LastEditedTimePropertyValue
	| LastEditedByPropertyValue;
*/
type PropertyValue interface{}

type PropertyValueBase struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

/*
 export interface TitlePropertyValue extends PropertyValueBase {
	 type: 'title';
	 title: RichText[];
 }

 export interface RichTextPropertyValue extends PropertyValueBase {
	 type: 'rich_text';
	 rich_text: RichText[];
 }

 export interface NumberPropertyValue extends PropertyValueBase {
	 type: 'number';
	 number: number;
 }

 export interface SelectPropertyValue extends PropertyValueBase {
	 type: 'select';
	 select: {
		 id: string;
		 name: string;
		 color: Color;
	 };
 }

 export interface MultiSelectPropertyValue extends PropertyValueBase {
	 type: 'multi_select';
	 multi_select: {
		 id: string;
		 name: string;
		 color: Color;
	 };
 }

 export interface DatePropertyValue extends PropertyValueBase {
	 type: 'date';
	 date: {
		 start: string;
		 end?: string;
	 };
 }

 export interface FormulaPropertyValue extends PropertyValueBase {
	 type: 'formula';
	 formula: StringFormulaValue | NumberFormulaValue | BooleanFormulaValue | DateFormulaValue;
 }

 export interface StringFormulaValue {
	 type: 'string';
	 string?: string;
 }
 export interface NumberFormulaValue {
	 type: 'number';
	 number?: number;
 }
 export interface BooleanFormulaValue {
	 type: 'boolean';
	 boolean: bool
 }
 export interface DateFormulaValue {
	 type: 'date';
	 date: DatePropertyValue;
 }

 export interface RollupPropertyValue extends PropertyValueBase {
	 type: 'rollup';
	 rollup: NumberRollupValue | DateRollupValue | ArrayRollupValue;
 }

 export interface NumberRollupValue {
	 type: 'number';
	 number: number;
 }
 export interface DateRollupValue {
	 type: 'date';
	 date: DatePropertyValue;
 }
 export interface ArrayRollupValue {
	 type: 'array';
	 array: unknown[]; // TODO: this is exactly like PropertyValue but without the `id`
 }

 export interface PeoplePropertyValue extends PropertyValueBase {
	 type: 'people';
	 people: User[];
 }

 export interface FilesPropertyValue extends PropertyValueBase {
	 type: 'files';
	 files: { name: string; }[];
 }

 export interface CheckboxPropertyValue extends PropertyValueBase {
	 type: 'checkbox';
	 checkbox: bool
 }

 export interface URLPropertyValue extends PropertyValueBase {
	 type: 'url';
	 url: string;
 }

 export interface EmailPropertyValue extends PropertyValueBase {
	 type: 'email';
	 email: string;
 }

 export interface PhoneNumberPropertyValue extends PropertyValueBase {
	 type: 'phone_number';
	 phone_number: string;
 }

 export interface CreatedTimePropertyValue extends PropertyValueBase {
	 type: 'created_time';
	 created_time: string;
 }

 export interface CreatedByPropertyValue extends PropertyValueBase {
	 type: 'created_by';
	 created_by: User;
 }

 export interface LastEditedTimePropertyValue extends PropertyValueBase {
	 type: 'last_edited_time';
	 last_edited_time: string;
 }

 export interface LastEditedByPropertyValue extends PropertyValueBase {
	 type: 'last_edited_by';
	 last_edited_by: User;
 }
*/

/*
* Rich text object (output)
 */

// RichTextText | RichTextMention | RichTextEquation;
type RichText interface{}

type RichTextBase struct {
	PlainText   string       `json:"plain_text"`
	Href        string       `json:"href"`
	Annotations *Annotations `json:"annotations"`
	Type        string       `json:"type"`
}

type RichTextText struct {
	RichTextBase
	//type: 'text';
	Text *struct {
		Content string `json:"content"`
		Link    *struct {
			Type string `json:"type"` // 'url'
			URL  string `json:"url"`
		} `json:"link"`
	} `json:"text"`
}

/*
 export interface RichTextMention extends RichTextBase {
	 type: 'mention';
	 mention: UserMention | PageMention | DatabaseMention | DateMention;
 }

 export interface UserMention {
	 type: 'user';
	 user: User;
 }

 export interface PageMention {
	 type: 'page';
	 page: { id: string; };
 }

 export interface DatabaseMention {
	 type: 'database';
	 database: { id: string };
 }

 export interface DateMention {
	 type: 'date';
	 date: DatePropertyValue;
 }

 export interface RichTextEquation extends RichTextBase {
	 type: 'equation';
	 equation: {
		 expression: string;
	 };
 }
*/

type Annotations struct {
	Bold          bool  `json:"bold"`
	Italic        bool  `json:"italic"`
	Strikethrough bool  `json:"strikethrough"`
	Underline     bool  `json:"underline"`
	Code          bool  `json:"code"`
	Color         Color `json:"color"`
}
