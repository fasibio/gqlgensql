package structure

import (
	"fmt"

	"github.com/vektah/gqlparser/v2/ast"
)

type Entity struct {
	BuiltIn    bool
	Raw        *ast.FieldDefinition
	TypeObject *Object
}

func (e Entity) Name() string {
	return e.Raw.Name
}

func (e Entity) HasGormDirective() bool {
	return e.GormDirectiveValue() != ""
}
func (e Entity) GormDirectiveValue() string {
	d := e.Raw.Directives.ForName(string(DirectiveSQLGorm))
	if d == nil {
		return ""
	}
	return d.Arguments.ForName("value").Value.Raw
}

func (e Entity) GqlTypeName() string {
	return e.Raw.Type.Name()
}

func (e Entity) GqlTypeObj() *Object {
	return e.TypeObject
}

func (e Entity) GqlType(suffix string) string {
	name := e.Raw.Type.Name()

	if e.BuiltIn {
		if e.IsArray() {
			return fmt.Sprintf("[%s!]", name)
		}
		return name
	}
	if e.IsArray() {
		return fmt.Sprintf("[%s%s!]", name, suffix)
	}
	return fmt.Sprintf("%s%s", name, suffix)
}

func (e *Entity) Required() bool {
	return e.Raw.Type.NonNull
}

func (e Entity) RequiredChar() string {
	requiredChar := ""
	if e.Required() {
		requiredChar = "!"
	}
	return requiredChar
}

func (e *Entity) IsArray() bool {
	return e.Raw.Type.String()[0] == '['
}
func (e *Entity) IsArrayElementRequired() bool {
	if !e.IsArray() {
		return false
	}
	return e.Raw.Type.Elem.NonNull
}

func (e *Entity) IsPrimary() bool {
	return e.Raw.Directives.ForName(string(DirectiveSQLPrimary)) != nil
}

func (e *Entity) IsIndex() bool {
	return e.Raw.Directives.ForName(string(DirectiveSQLIndex)) != nil
}

func (e *Entity) WhereAble() bool {
	switch e.Raw.Type.Name() {
	case "String", "DateTime", "Int", "Float", "ID":
		return true
	}
	return false
}

func (e *Entity) OrderAble() bool {
	switch e.Raw.Type.Name() {
	case "String", "DateTime", "Int", "Float", "ID":
		return true
	}
	return false
}
