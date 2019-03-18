// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package testserver

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type TestUnion interface {
	IsTestUnion()
}

type A struct {
	ID string `json:"id"`
}

func (A) IsTestUnion() {}

type AIt struct {
	ID string `json:"id"`
}

type AbIt struct {
	ID string `json:"id"`
}

type B struct {
	ID string `json:"id"`
}

func (B) IsTestUnion() {}

type EmbeddedDefaultScalar struct {
	Value *string `json:"value"`
}

type InnerDirectives struct {
	Message string `json:"message"`
}

type InnerInput struct {
	ID int `json:"id"`
}

type InnerObject struct {
	ID int `json:"id"`
}

type InputDirectives struct {
	Text          string           `json:"text"`
	Inner         InnerDirectives  `json:"inner"`
	InnerNullable *InnerDirectives `json:"innerNullable"`
	ThirdParty    *ThirdParty      `json:"thirdParty"`
}

type OuterInput struct {
	Inner InnerInput `json:"inner"`
}

type OuterObject struct {
	Inner InnerObject `json:"inner"`
}

type Slices struct {
	Test1 []*string `json:"test1"`
	Test2 []string  `json:"test2"`
	Test3 []*string `json:"test3"`
	Test4 []string  `json:"test4"`
}

type User struct {
	ID      int        `json:"id"`
	Friends []User     `json:"friends"`
	Created time.Time  `json:"created"`
	Updated *time.Time `json:"updated"`
}

type ValidInput struct {
	Break       string `json:"break"`
	Default     string `json:"default"`
	Func        string `json:"func"`
	Interface   string `json:"interface"`
	Select      string `json:"select"`
	Case        string `json:"case"`
	Defer       string `json:"defer"`
	Go          string `json:"go"`
	Map         string `json:"map"`
	Struct      string `json:"struct"`
	Chan        string `json:"chan"`
	Else        string `json:"else"`
	Goto        string `json:"goto"`
	Package     string `json:"package"`
	Switch      string `json:"switch"`
	Const       string `json:"const"`
	Fallthrough string `json:"fallthrough"`
	If          string `json:"if"`
	Range       string `json:"range"`
	Type        string `json:"type"`
	Continue    string `json:"continue"`
	For         string `json:"for"`
	Import      string `json:"import"`
	Return      string `json:"return"`
	Var         string `json:"var"`
	Underscore  string `json:"_"`
}

//  These things are all valid, but without care generate invalid go code
type ValidType struct {
	DifferentCase      string `json:"differentCase"`
	DifferentCaseOld   string `json:"different_case"`
	ValidInputKeywords bool   `json:"validInputKeywords"`
	ValidArgs          bool   `json:"validArgs"`
}

type XXIt struct {
	ID string `json:"id"`
}

type XxIt struct {
	ID string `json:"id"`
}

type AsdfIt struct {
	ID string `json:"id"`
}

type IIt struct {
	ID string `json:"id"`
}

type Status string

const (
	StatusOk    Status = "OK"
	StatusError Status = "ERROR"
)

var AllStatus = []Status{
	StatusOk,
	StatusError,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusOk, StatusError:
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
