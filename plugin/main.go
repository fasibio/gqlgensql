package main

import (
	"bytes"
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

type SqlBuilder struct {
	TypeName string            `json:"type_name,omitempty"`
	Fields   []SqlBuilderField `json:"fields,omitempty"`
	Query    SqlBuilderQuery   `json:"query,omitempty"`
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

type SqlBuilderField struct {
	Name    string `json:"name,omitempty"`
	GqlType string `json:"gql_type,omitempty"`
	Primary bool
}

type SqlBuilderQuery struct {
	Get          bool     `json:"get,omitempty"`
	Query        bool     `json:"query,omitempty"`
	Aggregate    bool     `json:"aggregate,omitempty"`
	DirectiveExt []string `json:"directiveEtx,omitempty"`
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
	}
}

func (ggs GqlGenSqlPlugin) Name() string {
	return "gqlgensql"
}
func (ggs GqlGenSqlPlugin) MutateConfig(cfg *config.Config) error {
	log.Println("MutateConfig")
	// for _, c := range cfg.Schema.Types {
	// 	if sqlDirective := c.Directives.ForName(DirectiveSQL); sqlDirective != nil {
	// 		builder := NewSqlBuilder()
	// 		builder.TypeName = c.Name
	// 		if a := sqlDirective.Arguments.ForName(ArgumentQuery); a != nil {
	// 			err := customizeSqlBuilderQuery(&builder.Query, a)
	// 			if err != nil {
	// 				return err
	// 			}
	// 		}
	// 		// cfg.Sources = append(cfg.Sources, &ast.Source{
	// 		// 	Name:    fmt.Sprintf("%s/%s", ggs.Name(), builder.TypeName),
	// 		// 	Input:   getExtendsSource(builder),
	// 		// 	BuiltIn: true,
	// 		// })

	// 		// cfg.Schema.Types["Query"].Fields = append(cfg.Schema.Types["Query"].Fields, &ast.FieldDefinition{
	// 		// 	Name: builder.TypeName,
	// 		// 	// Directives: getDirectiveList(builder.Query.DirectiveExt),
	// 		// 	Type: ast.ListType(ast.NamedType(c.Name, c.Position), c.Position),
	// 		// })
	// 	}
	// }

	return nil
}

func getExtendsSource(builder []SqlBuilder) string {
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

func getDirectiveList(d []string) ast.DirectiveList {
	res := make(ast.DirectiveList, len(d))
	for _, v := range d {
		tmp := ast.Directive{
			Name: v,
		}
		res = append(res, &tmp)
	}
	return res
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

func (ggs GqlGenSqlPlugin) GenerateCode(cfg *codegen.Data) error {
	log.Println("GenerateCode")

	return nil
}

func (ggs GqlGenSqlPlugin) InjectSourceEarly() *ast.Source {
	log.Println("InjectSourceEarly")

	input := `
	input SqlQueryParams {
		get: Boolean
		query: Boolean
		aggregate: Boolean
		directiveEtx: [String!]
	}
	directive @SQL(query:SqlQueryParams ) on OBJECT
	directive @SQL_PRIMARY on FIELD_DEFINITION
	`

	return &ast.Source{
		Name:    fmt.Sprintf("%s/dirictive.graphql", ggs.Name()),
		Input:   input,
		BuiltIn: true,
	}
}

func (ggs GqlGenSqlPlugin) InjectSourceLate(schema *ast.Schema) *ast.Source {
	log.Println("InjectSourceLate")
	builderList := make([]SqlBuilder, 0)
	for _, c := range schema.Types {
		if sqlDirective := c.Directives.ForName(DirectiveSQL); sqlDirective != nil {
			// Has Trigger directive
			builder := NewSqlBuilder()
			builder.TypeName = c.Name
			for _, field := range c.Fields {
				builder.Fields = append(builder.Fields, SqlBuilderField{
					Name:    field.Name,
					GqlType: field.Type.Name(),
					Primary: field.Directives.ForName(DirectiveSQLPrimary) != nil,
				})
			}
			if a := sqlDirective.Arguments.ForName(ArgumentQuery); a != nil {
				err := customizeSqlBuilderQuery(&builder.Query, a)
				if err != nil {
					panic(err)
				}
			}
			builderList = append(builderList, builder)
			// .Sources = append(cfg.Sources, &ast.Source{
			// 	Name:    fmt.Sprintf("%s/%s", ggs.Name(), builder.TypeName),
			// 	Input:   getExtendsSource(builder),
			// 	BuiltIn: true,
			// })

			// cfg.Schema.Types["Query"].Fields = append(cfg.Schema.Types["Query"].Fields, &ast.FieldDefinition{
			// 	Name: builder.TypeName,
			// 	// Directives: getDirectiveList(builder.Query.DirectiveExt),
			// 	Type: ast.ListType(ast.NamedType(c.Name, c.Position), c.Position),
			// })
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
