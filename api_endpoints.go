package notion

/*
 * blocks.children.append()
 */

type BlocksChildrenAppendPathParameters struct {
	BlockID string `json:"block_id"`
}

type BlocksChildrenAppendQueryParameters struct{}

type BlocksChildrenAppendBodyParameters struct {
	Children []Block `json:"children"`
}

type BlocksChildrenAppendParameters struct {
	BlocksChildrenAppendPathParameters
	BlocksChildrenAppendQueryParameters
	BlocksChildrenAppendBodyParameters
}

type BlocksChildrenAppendResponse struct {
	BlockBase
}

var blocksChildrenAppend = struct {
	method      string
	pathParams  []string
	queryParams []string
	bodyParams  []string
	path        func(p *BlocksChildrenAppendPathParameters) string
}{
	method:      "patch",
	pathParams:  []string{"block_id"},
	queryParams: nil,
	bodyParams:  []string{"children"},
	path: func(p *BlocksChildrenAppendPathParameters) string {
		return "blocks/" + p.BlockID + "/children"
	},
}

/*
 * blocks.children.list()
 */

type BlocksChildrenListPathParameters struct {
	BlockID string `json:"block_id"`
}

type BlocksChildrenListQueryParameters struct {
	PaginationParameters
}

type BlocksChildrenListBodyParameters struct{}

type BlocksChildrenListParameters struct {
	BlocksChildrenListPathParameters
	BlocksChildrenListQueryParameters
	BlocksChildrenListBodyParameters
}

type BlocksChildrenListResponse struct {
	PaginatedList // with type Block
}

var blocksChildrenList = struct {
	method      string
	pathParams  []string
	queryParams []string
	bodyParams  []string
	path        func(p *BlocksChildrenAppendPathParameters) string
}{
	method:      "get",
	pathParams:  []string{"block_id"},
	queryParams: []string{"start_cursor", "page_size"},
	bodyParams:  nil,
	path: func(p *BlocksChildrenAppendPathParameters) string {
		return "blocks/" + p.BlockID + "/children"
	},
}

/*
 * databases.list()
 */

/*
interface DatabasesListPathParameters {}
interface DatabasesListQueryParameters extends PaginationParameters {}
interface DatabasesListBodyParameters {}

export interface DatabasesListParameters extends DatabasesListPathParameters, DatabasesListQueryParameters, DatabasesListBodyParameters {}
export interface DatabasesListResponse extends PaginatedList<Database> {}

export const databasesList = {
  method: 'get',
  pathParams: [],
  queryParams: ['start_cursor', 'page_size'],
  bodyParams: [],
  path: () => `databases`,
} as const;
*/

/*
 * databases.query()
 */

/*
interface DatabasesQueryPathParameters {
  database_id: string;
}
interface DatabasesQueryQueryParameters {}
interface DatabasesQueryBodyParameters extends PaginationParameters {
  filter?: Filter;
  sorts?: Sort[];
}

export interface DatabasesQueryParameters extends DatabasesQueryPathParameters, DatabasesQueryQueryParameters, DatabasesQueryBodyParameters {}
export interface DatabasesQueryResponse extends PaginatedList<Page> {}

export const databasesQuery = {
  method: 'post',
  pathParams: ['database_id'],
  queryParams: [],
  bodyParams: ['filter', 'sorts', 'start_cursor', 'page_size'],
  path: (p: DatabasesQueryPathParameters) => `databases/${p.database_id}/query`,
} as const;
*/

/*
 * databases.retrieve()
 */

/*
interface DatabasesRetrievePathParameters {
  database_id: string;
}
interface DatabasesRetrieveQueryParameters {}
interface DatabasesRetrieveBodyParameters {}

export interface DatabasesRetrieveParameters extends DatabasesRetrievePathParameters, DatabasesRetrieveQueryParameters, DatabasesRetrieveBodyParameters {}
export interface DatabasesRetrieveResponse extends Database {}

export const databasesRetrieve = {
  method: 'get',
  pathParams: ['database_id'],
  queryParams: [],
  bodyParams: [],
  path: (p: DatabasesRetrievePathParameters) => `databases/${p.database_id}`,
} as const;
*/

/*
 * pages.create()
 */

/*
interface PagesCreatePathParameters {}
interface PagesCreateQueryParameters {}
interface PagesCreateBodyParameters {
  parent: ParentInput;
  properties: { [propertyName: string]: PropertyValue; };
  children?: Block[];
}

export interface PagesCreateParameters extends PagesCreatePathParameters, PagesCreateQueryParameters, PagesCreateBodyParameters {}
export interface PagesCreateResponse extends BlockBase {}

export const pagesCreate = {
  method: 'post',
  pathParams: [],
  queryParams: [],
  bodyParams: ['parent', 'properties', 'children'],
  path: () => `pages`,
} as const
*/

/*
 * pages.retrieve()
 */

/*
interface PagesRetrievePathParameters {
  page_id: string;
}
interface PagesRetrieveQueryParameters {}
interface PagesRetrieveBodyParameters {}

export interface PagesRetrieveParameters extends PagesRetrievePathParameters, PagesRetrieveQueryParameters, PagesRetrieveBodyParameters {}
export interface PagesRetrieveResponse extends Page {}

export const pagesRetrieve = {
  method: 'get',
  pathParams: ['page_id'],
  queryParams: [],
  bodyParams: [],
  path: (p: PagesRetrievePathParameters) => `pages/${p.page_id}`,
} as const;
*/

/*
 * pages.update()
 */

/*
interface PagesUpdatePathParameters {
  page_id: string;
}
interface PagesUpdateQueryParameters {}
interface PagesUpdateBodyParameters {
  properties: { [propertyNameOrId: string]: PropertyValue; };
}

export interface PagesUpdateParameters extends PagesUpdatePathParameters, PagesUpdateQueryParameters, PagesUpdateBodyParameters {}
export interface PagesUpdateResponse extends Page {}

export const pagesUpdate = {
  method: 'patch',
  pathParams: ['page_id'],
  queryParams: [],
  bodyParams: ['properties'],
  path: (p: PagesUpdatePathParameters) => `pages/${p.page_id}`,
} as const;
*/

/*
 * users.retrieve()
 */

/*
interface UsersRetrievePathParameters {
  user_id: string;
}
interface UsersRetrieveQueryParameters {}
interface UsersRetrieveBodyParameters {}

export interface UsersRetrieveParameters extends UsersRetrievePathParameters, UsersRetrieveQueryParameters, UsersRetrieveBodyParameters {}
export interface UsersRetrieveResponse extends UserBase {}

export const usersRetrieve = {
  method: 'get',
  pathParams: ['user_id'],
  queryParams: [],
  bodyParams: [],
  path: (p: UsersRetrievePathParameters) => `users/${p.user_id}`,
} as const;
*/

/*
 * users.list()
 */

/*
interface UsersListPathParameters {}
interface UsersListQueryParameters extends PaginationParameters {}
interface UsersListBodyParameters {}

export interface UsersListParameters extends UsersListPathParameters, UsersListQueryParameters, UsersListBodyParameters {}
export interface UsersListResponse extends PaginatedList<User> {}

export const usersList = {
  method: 'get',
  pathParams: [],
  queryParams: ['start_cursor', 'page_size'],
  bodyParams: [],
  path: () => `users`,
} as const;
*/

/*
 * search()
 */

/*
interface SearchPathParameters {}
interface SearchQueryParameters {}
interface SearchBodyParameters extends PaginationParameters {
  query?: string;
  sort?: SearchSort;
  filter?: SearchFilter;
}

export interface SearchParameters extends SearchPathParameters, SearchQueryParameters, SearchBodyParameters {}
export interface SearchResponse extends PaginatedList<Page | Database> {}

export const search = {
  method: 'post',
  pathParams: [],
  queryParams: [],
  bodyParams: ['query', 'sort', 'filter', 'start_cursor', 'page_size'],
  path: () => `search`,
} as const;
*/
