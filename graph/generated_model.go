// Code generated by github.com/vektah/gqlgen, DO NOT EDIT.

package graph

import (
	fmt "fmt"
	io "io"
	strconv "strconv"
	time "time"
)

type AddCommentOperation struct {
	Author  Person    `json:"author"`
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
}
type Authored interface{}
type Bug struct {
	ID         string              `json:"id"`
	HumanId    string              `json:"humanId"`
	Title      string              `json:"title"`
	Status     Status              `json:"status"`
	Labels     []string            `json:"labels"`
	Comments   CommentConnection   `json:"comments"`
	Operations OperationConnection `json:"operations"`
}
type BugConnection struct {
	Edges      []*BugEdge `json:"edges"`
	PageInfo   PageInfo   `json:"pageInfo"`
	TotalCount int        `json:"totalCount"`
}
type BugEdge struct {
	Cursor string `json:"cursor"`
	Node   Bug    `json:"node"`
}
type Comment struct {
	Author  Person `json:"author"`
	Message string `json:"message"`
}
type CommentConnection struct {
	Edges      []CommentEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}
type CommentEdge struct {
	Cursor string  `json:"cursor"`
	Node   Comment `json:"node"`
}
type ConnectionInput struct {
	After  *string `json:"after"`
	Before *string `json:"before"`
	First  *int    `json:"first"`
	Last   *int    `json:"last"`
}
type CreateOperation struct {
	Author  Person    `json:"author"`
	Date    time.Time `json:"date"`
	Title   string    `json:"title"`
	Message string    `json:"message"`
}
type LabelChangeOperation struct {
	Author  Person    `json:"author"`
	Date    time.Time `json:"date"`
	Added   []string  `json:"added"`
	Removed []string  `json:"removed"`
}
type Operation interface{}
type OperationConnection struct {
	Edges      []OperationEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}
type OperationEdge struct {
	Cursor string         `json:"cursor"`
	Node   OperationUnion `json:"node"`
}
type OperationUnion interface{}
type PageInfo struct {
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
}
type Person struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
}
type Repository struct {
	AllBugs BugConnection `json:"allBugs"`
	Bug     *Bug          `json:"bug"`
}
type SetStatusOperation struct {
	Author Person    `json:"author"`
	Date   time.Time `json:"date"`
	Status Status    `json:"status"`
}
type SetTitleOperation struct {
	Author Person    `json:"author"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}

type Status string

const (
	StatusOpen   Status = "OPEN"
	StatusClosed Status = "CLOSED"
)

func (e Status) IsValid() bool {
	switch e {
	case StatusOpen, StatusClosed:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}