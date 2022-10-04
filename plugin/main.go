package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/vektah/gqlparser/v2/ast"
)

const (
	DirectiveSQL        = "SQL"
	DirectiveSQLPrimary = "SQL_PRIMARY"
	ArgumentQuery       = "query"
)

type GqlGenSqlPlugin struct {
}

type SqlBuilderList []SqlBuilder

func (sbl SqlBuilderList) AllRef() map[string]SqlBuilder {
	res := make(map[string]SqlBuilder)
	for _, a := range sbl {
		b := AllRefFields(a.Fields)
		for k, v := range b {
			res[k] = v
		}
	}
	b, _ := json.Marshal(res)
	log.Println("AllRef", string(b))
	return res
}

func AllRefFields(fields []SqlBuilderField) map[string]SqlBuilder {
	res := make(map[string]SqlBuilder)
	for _, f := range fields {
		if f.Ref != nil {
			res[f.Ref.TypeName] = *f.Ref
			b := AllRefFields(f.Ref.Fields)
			for k, v := range b {
				res[k] = v
			}
		}
	}
	return res
}

type SqlBuilder struct {
	TypeName string             `json:"type_name,omitempty"`
	Fields   []SqlBuilderField  `json:"fields,omitempty"`
	Query    SqlBuilderQuery    `json:"query,omitempty"`
	Mutation SqlBuilderMutation `json:"mutation,omitempty"`
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

type SqlBuilderField struct {
	Name    string `json:"name,omitempty"`
	GqlType string `json:"gql_type,omitempty"`
	Primary bool
	Ref     *SqlBuilder
}

type SqlBuilderQuery struct {
	Get          bool     `json:"get,omitempty"`
	Query        bool     `json:"query,omitempty"`
	Aggregate    bool     `json:"aggregate,omitempty"`
	DirectiveExt []string `json:"directiveEtx,omitempty"`
}

type SqlBuilderMutation struct {
	Add          bool     `json:"add,omitempty"`
	Update       bool     `json:"update,omitempty"`
	Delete       bool     `json:"delete,omitempty"`
	DirectiveExt []string `json:"directive_ext,omitempty"`
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

func (ggs GqlGenSqlPlugin) Name() string {
	return "gqlgensql"
}
func (ggs GqlGenSqlPlugin) MutateConfig(cfg *config.Config) error {
	log.Println("MutateConfig")
	return nil
}

func getExtendsSource(builder SqlBuilderList) string {
	// to generate refs []sqlBuilder needs a methode to find out all refs where it is nessasary to gen a new input file
	builder.AllRef()
	temp := `
{{- range $key, $value := .}}
enum {{$value.TypeName}}HasFilter {
	{{- range $fieldKey, $field := $value.Fields}}
  {{$field.Name}}
	{{- end}}
}

enum {{$value.TypeName}}Orderable {
	{{- range $fieldKey, $field := $value.OrderAbleFields}}
  {{$field.Name}}
	{{- end}}
}

input {{$value.TypeName}}Filter {
	{{$value.PrimaryField.Name}}: [{{$value.PrimaryField.GqlType}}!]
	has: [{{$value.TypeName}}HasFilter]
	and: [{{$value.TypeName}}Filter]
	or: [{{$value.TypeName}}Filter]
	not: [{{$value.TypeName}}Filter]
}

input {{$value.TypeName}}Order {
	asc: {{$value.TypeName}}Orderable
	desc: {{$value.TypeName}}Orderable
	then: {{$value.TypeName}}Order
}

type {{$value.TypeName}}AggregateResult {
	count: Int
	{{- range $fieldKey, $field := $value.AggregateFields}}
  {{$field.Name}}Min: {{$field.GqlType}}
	{{$field.Name}}Max: {{$field.GqlType}}
	{{- end}}
}

input Add{{$value.TypeName}}Input {
	{{- range $fieldKey, $field := $value.InputFields}}
  {{$field.Name}}: {{$field.GqlType}}!
	{{- end}}
}

type Add{{$value.TypeName}}Payload {
	{{$value.TypeName}}(filter: {{$value.TypeName}}Filter, order: {{$value.TypeName}}Order, first: Int, offset: Int): [{{$value.TypeName}}]
}

{{- end}}
extend type Query {
	empty: String! # Hack for empty query
	{{- range $key, $value := .}}
		{{- if $value.Query.Get}}
  get{{$value.TypeName}}({{$value.PrimaryField.Name}}: {{$value.PrimaryField.GqlType}}!): {{$value.TypeName}}{{ range $directiveKey, $directive := $value.Query.DirectiveExt}} {{$directive}}{{end}}
		{{- end}}
		{{- if $value.Query.Query}}
  query{{$value.TypeName}}(filter: {{$value.TypeName}}Filter, order: {{$value.TypeName}}Order, first: Int, offset: Int): [{{$value.TypeName}}]{{ range $directiveKey, $directive := $value.Query.DirectiveExt}} {{$directive}}{{end}}
		{{- end}}
		{{- if $value.Query.Aggregate}}
  aggregate{{$value.TypeName}}(filter: {{$value.TypeName}}Filter): {{$value.TypeName}}AggregateResult{{ range $directiveKey, $directive := $value.Query.DirectiveExt}} {{$directive}}{{end}}
			{{- end}}
	{{- end}}
}


	`
	tmpl, _ := template.New("sourcebuilder").Parse(temp)
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, builder)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func getArrayOfInterface[K comparable](v interface{}) []K {
	aInterface := v.([]interface{})
	aGen := make([]K, len(aInterface))
	for i, v := range aInterface {
		aGen[i] = v.(K)
	}
	return aGen
}

func customizeSqlBuilderQuery(s *SqlBuilderQuery, a *ast.Argument) error {
	for _, e := range a.Value.Children {
		log.Println("hier", e.Name)
		v, _ := e.Value.Value(nil)
		switch e.Name {
		case "query":
			s.Query = v.(bool)
		case "get":
			s.Get = v.(bool)
		case "aggregate":
			s.Aggregate = v.(bool)
		case "directiveEtx":
			s.DirectiveExt = getArrayOfInterface[string](v)
		}

	}
	return nil
}

func customizeSqlBuilderMutation(s *SqlBuilderMutation, a *ast.Argument) error {
	for _, e := range a.Value.Children {
		v, _ := e.Value.Value(nil)
		switch e.Name {
		case "add":
			s.Add = v.(bool)
		case "update":
			s.Update = v.(bool)
		case "delete":
			s.Delete = v.(bool)
		case "directiveEtx":
			s.DirectiveExt = getArrayOfInterface[string](v)
		}

	}
	return nil
}

func (ggs GqlGenSqlPlugin) GenerateCode(cfg *codegen.Data) error {
	log.Println("GenerateCode")

	return nil
}

func (ggs GqlGenSqlPlugin) InjectSourceEarly() *ast.Source {
	log.Println("InjectSourceEarly")

	input := `

	input SqlMutationParams {
		add: Boolean
		update: Boolean
		delete: Boolean
		directiveEtx: [String!]
	}

	input SqlQueryParams {
		get: Boolean
		query: Boolean
		aggregate: Boolean
		directiveEtx: [String!]
	}
	directive @SQL(query:SqlQueryParams, mutation: SqlMutationParams ) on OBJECT
	directive @SQL_PRIMARY on FIELD_DEFINITION
	`

	return &ast.Source{
		Name:    fmt.Sprintf("%s/directive.graphql", ggs.Name()),
		Input:   input,
		BuiltIn: true,
	}
}

func fillSqlBuilderByName(schema *ast.Schema, name string) *SqlBuilder {
	val, ok := schema.Types[name]
	if !ok {
		log.Println("mhm warum passiert das ???", name)
	}
	if val.BuiltIn {
		return nil
	}
	res := NewSqlBuilder()
	res.TypeName = val.Name
	res.Fields = getSqlBuilderFields(val.Fields, schema)
	return &res
}

func getSqlBuilderFields(fields ast.FieldList, schema *ast.Schema) []SqlBuilderField {
	res := make([]SqlBuilderField, 0)
	for _, field := range fields {
		res = append(res, SqlBuilderField{
			Name:    field.Name,
			GqlType: field.Type.Name(),
			Primary: field.Directives.ForName(DirectiveSQLPrimary) != nil,
			Ref:     fillSqlBuilderByName(schema, field.Type.Name()),
		})
	}
	b, _ := json.Marshal(res)
	log.Println("lalala", string(b))
	return res
}

func (ggs GqlGenSqlPlugin) InjectSourceLate(schema *ast.Schema) *ast.Source {
	log.Println("InjectSourceLate")
	builderList := make(SqlBuilderList, 0)
	for _, c := range schema.Types {
		log.Println("lalal", c.Name, c.BuiltIn)
		if sqlDirective := c.Directives.ForName(DirectiveSQL); sqlDirective != nil {
			// Has Trigger directive
			builder := NewSqlBuilder()
			builder.TypeName = c.Name
			builder.Fields = getSqlBuilderFields(c.Fields, schema)
			if a := sqlDirective.Arguments.ForName(ArgumentQuery); a != nil {
				err := customizeSqlBuilderQuery(&builder.Query, a)
				if err != nil {
					panic(err)
				}
				err = customizeSqlBuilderMutation(&builder.Mutation, a)
				if err != nil {
					panic(err)
				}
			}
			builderList = append(builderList, builder)
		}
	}
	result := getExtendsSource(builderList)
	log.Println(result)
	return &ast.Source{
		Name:    fmt.Sprintf("%s/gqlgenSql.graphql", ggs.Name()),
		Input:   result,
		BuiltIn: true,
	}
}

func main() {
	log.Println("da")

	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	p := GqlGenSqlPlugin{}

	err = api.Generate(cfg, api.AddPlugin(p))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
