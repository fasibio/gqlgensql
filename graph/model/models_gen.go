// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AddCatInput struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Owner       *UserRef `json:"owner"`
	Partner     *UserRef `json:"partner"`
	Description string   `json:"description"`
}

type AddCatPayload struct {
	Cat []*Cat `json:"Cat"`
}

type AddTodoInput struct {
	ID    string   `json:"id"`
	Text  string   `json:"text"`
	Done  bool     `json:"done"`
	Owner *UserRef `json:"owner"`
}

type AddTodoPayload struct {
	Todo []*Todo `json:"Todo"`
}

type AddUserInput struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Test  string     `json:"test"`
	Todos []*TodoRef `json:"todos"`
}

type AddUserPayload struct {
	User []*User `json:"User"`
}

type Cat struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Owner       *User  `json:"owner"`
	Partner     *User  `json:"partner"`
	Description string `json:"description"`
}

type CatAggregateResult struct {
	Count          *int    `json:"count"`
	NameMin        *string `json:"nameMin"`
	NameMax        *string `json:"nameMax"`
	DescriptionMin *string `json:"descriptionMin"`
	DescriptionMax *string `json:"descriptionMax"`
}

type CatFilter struct {
	ID  []string        `json:"id"`
	Has []*CatHasFilter `json:"has"`
	And []*CatFilter    `json:"and"`
	Or  []*CatFilter    `json:"or"`
	Not []*CatFilter    `json:"not"`
}

type CatOrder struct {
	Asc  *CatOrderable `json:"asc"`
	Desc *CatOrderable `json:"desc"`
	Then *CatOrder     `json:"then"`
}

type CatPatch struct {
	ID          *string  `json:"id"`
	Name        *string  `json:"name"`
	Owner       *UserRef `json:"owner"`
	Partner     *UserRef `json:"partner"`
	Description *string  `json:"description"`
}

type DeleteCatPayload struct {
	Cat    []*Cat  `json:"Cat"`
	NumIds *int    `json:"numIds"`
	Msg    *string `json:"msg"`
}

type DeleteTodoPayload struct {
	Todo   []*Todo `json:"Todo"`
	NumIds *int    `json:"numIds"`
	Msg    *string `json:"msg"`
}

type DeleteUserPayload struct {
	User   []*User `json:"User"`
	NumIds *int    `json:"numIds"`
	Msg    *string `json:"msg"`
}

type SQLMutationParams struct {
	Add          *bool    `json:"add"`
	Update       *bool    `json:"update"`
	Delete       *bool    `json:"delete"`
	DirectiveEtx []string `json:"directiveEtx"`
}

type SQLQueryParams struct {
	Get          *bool    `json:"get"`
	Query        *bool    `json:"query"`
	Aggregate    *bool    `json:"aggregate"`
	DirectiveEtx []string `json:"directiveEtx"`
}

type Todo struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Done  bool   `json:"done"`
	Owner *User  `json:"owner"`
}

type TodoAggregateResult struct {
	Count   *int    `json:"count"`
	TextMin *string `json:"textMin"`
	TextMax *string `json:"textMax"`
}

type TodoFilter struct {
	ID  []string         `json:"id"`
	Has []*TodoHasFilter `json:"has"`
	And []*TodoFilter    `json:"and"`
	Or  []*TodoFilter    `json:"or"`
	Not []*TodoFilter    `json:"not"`
}

type TodoOrder struct {
	Asc  *TodoOrderable `json:"asc"`
	Desc *TodoOrderable `json:"desc"`
	Then *TodoOrder     `json:"then"`
}

type TodoPatch struct {
	ID    *string  `json:"id"`
	Text  *string  `json:"text"`
	Done  *bool    `json:"done"`
	Owner *UserRef `json:"owner"`
}

type TodoRef struct {
	ID    *string  `json:"id"`
	Text  *string  `json:"text"`
	Done  *bool    `json:"done"`
	Owner *UserRef `json:"owner"`
}

type UpdateCatInput struct {
	Filter *CatFilter `json:"filter"`
	Set    *CatPatch  `json:"set"`
	Remove *CatPatch  `json:"remove"`
}

type UpdateCatPayload struct {
	Cat    []*Cat `json:"Cat"`
	NumIds *int   `json:"numIds"`
}

type UpdateTodoInput struct {
	Filter *TodoFilter `json:"filter"`
	Set    *TodoPatch  `json:"set"`
	Remove *TodoPatch  `json:"remove"`
}

type UpdateTodoPayload struct {
	Todo   []*Todo `json:"Todo"`
	NumIds *int    `json:"numIds"`
}

type UpdateUserInput struct {
	Filter *UserFilter `json:"filter"`
	Set    *UserPatch  `json:"set"`
	Remove *UserPatch  `json:"remove"`
}

type UpdateUserPayload struct {
	User   []*User `json:"User"`
	NumIds *int    `json:"numIds"`
}

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Test  string  `json:"test"`
	Todos []*Todo `json:"todos"`
}

type UserAggregateResult struct {
	Count   *int    `json:"count"`
	NameMin *string `json:"nameMin"`
	NameMax *string `json:"nameMax"`
	TestMin *string `json:"testMin"`
	TestMax *string `json:"testMax"`
}

type UserFilter struct {
	ID  []string         `json:"id"`
	Has []*UserHasFilter `json:"has"`
	And []*UserFilter    `json:"and"`
	Or  []*UserFilter    `json:"or"`
	Not []*UserFilter    `json:"not"`
}

type UserOrder struct {
	Asc  *UserOrderable `json:"asc"`
	Desc *UserOrderable `json:"desc"`
	Then *UserOrder     `json:"then"`
}

type UserPatch struct {
	ID    *string    `json:"id"`
	Name  *string    `json:"name"`
	Test  *string    `json:"test"`
	Todos []*TodoRef `json:"todos"`
}

type UserRef struct {
	ID    *string    `json:"id"`
	Name  *string    `json:"name"`
	Test  *string    `json:"test"`
	Todos []*TodoRef `json:"todos"`
}

type CatHasFilter string

const (
	CatHasFilterID          CatHasFilter = "id"
	CatHasFilterName        CatHasFilter = "name"
	CatHasFilterOwner       CatHasFilter = "owner"
	CatHasFilterPartner     CatHasFilter = "partner"
	CatHasFilterDescription CatHasFilter = "description"
)

var AllCatHasFilter = []CatHasFilter{
	CatHasFilterID,
	CatHasFilterName,
	CatHasFilterOwner,
	CatHasFilterPartner,
	CatHasFilterDescription,
}

func (e CatHasFilter) IsValid() bool {
	switch e {
	case CatHasFilterID, CatHasFilterName, CatHasFilterOwner, CatHasFilterPartner, CatHasFilterDescription:
		return true
	}
	return false
}

func (e CatHasFilter) String() string {
	return string(e)
}

func (e *CatHasFilter) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CatHasFilter(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CatHasFilter", str)
	}
	return nil
}

func (e CatHasFilter) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CatOrderable string

const (
	CatOrderableName        CatOrderable = "name"
	CatOrderableDescription CatOrderable = "description"
)

var AllCatOrderable = []CatOrderable{
	CatOrderableName,
	CatOrderableDescription,
}

func (e CatOrderable) IsValid() bool {
	switch e {
	case CatOrderableName, CatOrderableDescription:
		return true
	}
	return false
}

func (e CatOrderable) String() string {
	return string(e)
}

func (e *CatOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CatOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CatOrderable", str)
	}
	return nil
}

func (e CatOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TodoHasFilter string

const (
	TodoHasFilterID    TodoHasFilter = "id"
	TodoHasFilterText  TodoHasFilter = "text"
	TodoHasFilterDone  TodoHasFilter = "done"
	TodoHasFilterOwner TodoHasFilter = "owner"
)

var AllTodoHasFilter = []TodoHasFilter{
	TodoHasFilterID,
	TodoHasFilterText,
	TodoHasFilterDone,
	TodoHasFilterOwner,
}

func (e TodoHasFilter) IsValid() bool {
	switch e {
	case TodoHasFilterID, TodoHasFilterText, TodoHasFilterDone, TodoHasFilterOwner:
		return true
	}
	return false
}

func (e TodoHasFilter) String() string {
	return string(e)
}

func (e *TodoHasFilter) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TodoHasFilter(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TodoHasFilter", str)
	}
	return nil
}

func (e TodoHasFilter) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TodoOrderable string

const (
	TodoOrderableText TodoOrderable = "text"
)

var AllTodoOrderable = []TodoOrderable{
	TodoOrderableText,
}

func (e TodoOrderable) IsValid() bool {
	switch e {
	case TodoOrderableText:
		return true
	}
	return false
}

func (e TodoOrderable) String() string {
	return string(e)
}

func (e *TodoOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TodoOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TodoOrderable", str)
	}
	return nil
}

func (e TodoOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserHasFilter string

const (
	UserHasFilterID    UserHasFilter = "id"
	UserHasFilterName  UserHasFilter = "name"
	UserHasFilterTest  UserHasFilter = "test"
	UserHasFilterTodos UserHasFilter = "todos"
)

var AllUserHasFilter = []UserHasFilter{
	UserHasFilterID,
	UserHasFilterName,
	UserHasFilterTest,
	UserHasFilterTodos,
}

func (e UserHasFilter) IsValid() bool {
	switch e {
	case UserHasFilterID, UserHasFilterName, UserHasFilterTest, UserHasFilterTodos:
		return true
	}
	return false
}

func (e UserHasFilter) String() string {
	return string(e)
}

func (e *UserHasFilter) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserHasFilter(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserHasFilter", str)
	}
	return nil
}

func (e UserHasFilter) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserOrderable string

const (
	UserOrderableName UserOrderable = "name"
	UserOrderableTest UserOrderable = "test"
)

var AllUserOrderable = []UserOrderable{
	UserOrderableName,
	UserOrderableTest,
}

func (e UserOrderable) IsValid() bool {
	switch e {
	case UserOrderableName, UserOrderableTest:
		return true
	}
	return false
}

func (e UserOrderable) String() string {
	return string(e)
}

func (e *UserOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserOrderable", str)
	}
	return nil
}

func (e UserOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
