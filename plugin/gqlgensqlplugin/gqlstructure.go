package gqlgensqlplugin

import (
	"fmt"
	"log"

	"github.com/vektah/gqlparser/v2/ast"
)

type SqlBuilderList map[string]SqlBuilder
type SqlBuilderRefs map[string]*SqlBuilder
type SqlBuilderHandler struct {
	List SqlBuilderList
	Refs SqlBuilderRefs
}

type SqlBuilder struct {
	TypeName string             `json:"type_name,omitempty"`
	Fields   []SqlBuilderField  `json:"fields,omitempty"`
	Query    SqlBuilderQuery    `json:"query,omitempty"`
	Mutation SqlBuilderMutation `json:"mutation,omitempty"`
}

func (sbh SqlBuilderRefs) ReferObjects() map[string]SqlBuilder {
	res := make(map[string]SqlBuilder)
	for key, value := range sbh {
		res[key] = *value
	}
	return res
}

func NewSqlBuilderHandler() SqlBuilderHandler {
	return SqlBuilderHandler{
		List: make(SqlBuilderList, 0),
		Refs: make(SqlBuilderRefs),
	}
}

func (s SqlBuilder) PrimaryField() SqlBuilderField {
	for _, a := range s.Fields {
		if a.Primary {
			return a
		}
	}
	for _, a := range s.Fields {
		if a.GqlType == "ID" {
			return a
		}
	}
	log.Panicf("Type %s: Can not find a Field with Directive %s or with type ID!", s.TypeName, DirectiveSQLPrimary)
	return SqlBuilderField{}
}

func (s SqlBuilder) OrderAbleFields() []SqlBuilderField {
	res := make([]SqlBuilderField, 0)
	for _, a := range s.Fields {
		switch a.GqlType {
		case "String", "DateTime", "Int", "Float":
			res = append(res, a)
		}
	}
	return res
}

func (s SqlBuilder) AggregateFields() []SqlBuilderField {
	return s.OrderAbleFields()
}

// / All field which can be added by mutation from client
func (s SqlBuilder) InputFields() []SqlBuilderField {
	res := make([]SqlBuilderField, 0)
	for _, a := range s.Fields {
		if !a.Primary {
			res = append(res, a)
		}
	}
	return res
}

// / All field which can be added by mutation from client
func (s SqlBuilder) InputRefFields() []SqlBuilderField {
	return s.Fields
}

func (s SqlBuilder) InputPatchFields() []SqlBuilderField {
	res := make([]SqlBuilderField, 0)
	for _, v := range s.Fields {
		if !v.Primary {
			res = append(res, v)
		}
	}
	return res
}

type SqlBuilderField struct {
	Name    string `json:"name,omitempty"`
	GqlType string `json:"gql_type,omitempty"`
	Primary bool
	BuiltIn bool
	Raw     *ast.FieldDefinition
}

func (sbf *SqlBuilderField) RefGqlType() string {
	if sbf.BuiltIn {
		return sbf.GqlType
	}
	if isGqlArray(sbf.Raw.Type.String()) {
		return fmt.Sprintf("[%sRef!]", sbf.GqlType)
	}
	return fmt.Sprintf("%sRef", sbf.GqlType)
}

func isGqlArray(v string) bool {
	return v[0] == '['
}

type SqlBuilderQuery struct {
	Get          bool     `json:"get,omitempty"`
	Query        bool     `json:"query,omitempty"`
	Aggregate    bool     `json:"aggregate,omitempty"`
	DirectiveExt []string `json:"directiveEtx,omitempty"`
}

func (sbq *SqlBuilderQuery) HasQueries() bool {
	return sbq.Query || sbq.Get || sbq.Aggregate
}

func NewSqlBuilder() SqlBuilder {
	return SqlBuilder{
		Fields: make([]SqlBuilderField, 0),
		Query: SqlBuilderQuery{
			Get:          true,
			Query:        true,
			Aggregate:    true,
			DirectiveExt: make([]string, 0),
		},
		Mutation: SqlBuilderMutation{
			Add:          true,
			Update:       true,
			Delete:       true,
			DirectiveExt: make([]string, 0),
		},
	}
}

type SqlBuilderMutation struct {
	Add          bool     `json:"add,omitempty"`
	Update       bool     `json:"update,omitempty"`
	Delete       bool     `json:"delete,omitempty"`
	DirectiveExt []string `json:"directive_ext,omitempty"`
}
