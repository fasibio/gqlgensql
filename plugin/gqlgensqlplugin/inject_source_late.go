package gqlgensqlplugin

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"text/template"

	"github.com/vektah/gqlparser/v2/ast"
)

func (ggs *GqlGenSqlPlugin) InjectSourceLate(schema *ast.Schema) *ast.Source {
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
	ggs.handler = builderHandler
	// log.Println(result)
	return &ast.Source{
		Name:    fmt.Sprintf("%s/gqlgenSql.graphql", ggs.Name()),
		Input:   result,
		BuiltIn: false,
	}
}

//go:embed inject_source_late.gql.go.tpl
var gqltemplate string

func getExtendsSource(builder SqlBuilderHandler) string {
	tmpl, _ := template.New("sourcebuilder").Parse(gqltemplate)
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
			Primary: field.Directives.ForName(DirectiveSQL) != nil,
			BuiltIn: schema.Types[field.Type.Name()].BuiltIn,
			Raw:     field,
		})
		fillSqlBuilderByName(schema, field.Type.Name(), knownValues)
	}
	return res
}
