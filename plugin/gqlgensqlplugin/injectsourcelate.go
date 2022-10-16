package gqlgensqlplugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"text/template"

	"github.com/vektah/gqlparser/v2/ast"
)

func (ggs GqlGenSqlPlugin) InjectSourceLate(schema *ast.Schema) *ast.Source {
	log.Println("InjectSourceLate")
	builderHandler := NewSqlBuilderHandler()
	for _, c := range schema.Types {
		if sqlDirective := c.Directives.ForName(DirectiveSQL); sqlDirective != nil {
			// Has Trigger directive
			builder := NewSqlBuilder()
			builder.TypeName = c.Name
			a := make(SqlBuilderRefs)
			f := getSqlBuilderFields(c.Fields, schema, a)
			builder.Fields = f
			for k, v := range a {
				builderHandler.Refs[k] = v
				if _, ok := builderHandler.List[k]; !ok {
					builderHandler.List[k] = *v // @TODO needed to generate all queries and mutations for ref types also ?
				}

			}
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
			builderHandler.List[builder.TypeName] = builder
		}
	}
	result := getExtendsSource(builderHandler)
	log.Println(result)
	return &ast.Source{
		Name:    fmt.Sprintf("%s/gqlgenSql.graphql", ggs.Name()),
		Input:   result,
		BuiltIn: true,
	}
}

func getExtendsSource(builder SqlBuilderHandler) string {

	b, _ := json.Marshal(builder.Refs)
	log.Println(string(b))
	// to generate refs []sqlBuilder needs a methode to find out all refs where it is nessasary to gen a new input file
	temp := `

{{- range $key, $value := .Refs.ReferObjects}}

input {{$key}}Ref {
	{{- range $fieldKey, $field := $value.InputRefFields}}
  {{$field.Name}}: {{$field.RefGqlType}}
	{{- end}}
}

{{- end}}

{{- range $key, $value := .List}}
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



type Add{{$value.TypeName}}Payload {
	{{$value.TypeName}}(filter: {{$value.TypeName}}Filter, order: {{$value.TypeName}}Order, first: Int, offset: Int): [{{$value.TypeName}}]
}

type Update{{$value.TypeName}}Payload {
	{{$value.TypeName}}(filter: {{$value.TypeName}}Filter, order: {{$value.TypeName}}Order, first: Int, offset: Int): [{{$value.TypeName}}]
	numIds: Int
}

type Delete{{$value.TypeName}}Payload {
	{{$value.TypeName}}(filter: {{$value.TypeName}}Filter, order: {{$value.TypeName}}Order, first: Int, offset: Int): [{{$value.TypeName}}]
	numIds: Int
	msg: String
}

input Add{{$value.TypeName}}Input{
	{{- range $fieldKey, $field := $value.Fields}}
  {{$field.Name}}: {{$field.RefGqlType}}!
	{{- end}}
}

input {{$value.TypeName}}Patch{
	{{- range $fieldKey, $field := $value.InputPatchFields}}
  {{$field.Name}}: {{$field.RefGqlType}}
	{{- end}}
}

input Update{{$value.TypeName}}Input{
	filter: {{$value.TypeName}}Filter
	set: {{$value.TypeName}}Patch
	remove: {{$value.TypeName}}Patch
}

{{- end}}
extend type Query {
	{{- range $key, $value := .List}}
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

extend type Mutation {
	{{- range $key, $value := .List}}
	{{- if $value.Mutation.Add}}
	add{{$value.TypeName}}(input: [Add{{$value.TypeName}}Input!]!): Add{{$value.TypeName}}Payload
	{{- end}}
	{{- if $value.Mutation.Update}}
	update{{$value.TypeName}}(input: Update{{$value.TypeName}}Input!):  Update{{$value.TypeName}}Payload
	{{- end}}
	{{- if $value.Mutation.Delete}}
	delete{{$value.TypeName}}(filter: {{$value.TypeName}}Filter! ): Delete{{$value.TypeName}}Payload
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

func customizeSqlBuilderQuery(s *SqlBuilderQuery, a *ast.Argument) error {
	for _, e := range a.Value.Children {
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

func fillSqlBuilderByName(schema *ast.Schema, name string, knownValues SqlBuilderRefs) {

	val := schema.Types[name]
	if val.BuiltIn {
		return
	}
	if _, isOk := knownValues[val.Name]; isOk {
		return
	} else {
		tmp := NewSqlBuilder()
		tmp.TypeName = val.Name
		knownValues[val.Name] = &tmp
		f := getSqlBuilderFields(val.Fields, schema, knownValues)
		tmp.Fields = f

	}
}

func getSqlBuilderFields(fields ast.FieldList, schema *ast.Schema, knownValues SqlBuilderRefs) []SqlBuilderField {
	res := make([]SqlBuilderField, 0)
	for _, field := range fields {
		res = append(res, SqlBuilderField{
			Name:    field.Name,
			GqlType: field.Type.Name(),
			Primary: field.Directives.ForName(DirectiveSQLPrimary) != nil,
			BuiltIn: schema.Types[field.Type.Name()].BuiltIn,
			Raw:     field,
		})
		fillSqlBuilderByName(schema, field.Type.Name(), knownValues)
	}
	return res
}
