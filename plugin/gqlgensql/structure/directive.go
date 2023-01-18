package structure

import (
	"github.com/fasibio/gqlgensql/plugin/gqlgensql/helper"
	"github.com/vektah/gqlparser/v2/ast"
)

type SQLDirective struct {
	Query    SQLDirectiveQuery
	Mutation SQLDirectiveMutation
}

func (sd *SQLDirective) HasQueries() bool {
	return sd.Query.Query || sd.Query.Get
}

func (sd *SQLDirective) HasMutation() bool {
	return sd.Mutation.Add || sd.Mutation.Delete || sd.Mutation.Update
}

type SQLDirectiveHandler struct {
	DirectiveExt []string
}

type SQLDirectiveMutation struct {
	SQLDirectiveHandler
	Add    bool
	Update bool
	Delete bool
}

type SQLDirectiveQuery struct {
	SQLDirectiveHandler
	Get   bool
	Query bool
}

func getDefaultFilledSqlBuilderMutation(defaultValue bool) SQLDirectiveMutation {
	return SQLDirectiveMutation{
		Add:                 defaultValue,
		SQLDirectiveHandler: SQLDirectiveHandler{DirectiveExt: []string{}},
		Update:              defaultValue,
		Delete:              defaultValue,
	}
}

func getDefaultFilledSqlBuilderQuery(defaultValue bool) SQLDirectiveQuery {
	return SQLDirectiveQuery{
		Get:                 defaultValue,
		SQLDirectiveHandler: SQLDirectiveHandler{DirectiveExt: []string{}},
		Query:               defaultValue,
	}
}

func customizeSqlBuilderQuery(a *ast.Argument) SQLDirectiveQuery {
	res := getDefaultFilledSqlBuilderQuery(false)
	for _, e := range a.Value.Children {
		v, _ := e.Value.Value(nil)
		switch e.Name {
		case "query":
			res.Query = v.(bool)
		case "get":
			res.Get = v.(bool)
		case "directiveEtx":
			res.DirectiveExt = helper.GetArrayOfInterface[string](v)
		}

	}
	return res
}

func customizeSqlBuilderMutation(a *ast.Argument) SQLDirectiveMutation {
	res := getDefaultFilledSqlBuilderMutation(false)
	for _, e := range a.Value.Children {
		v, _ := e.Value.Value(nil)
		switch e.Name {
		case "add":
			res.Add = v.(bool)
		case "update":
			res.Update = v.(bool)
		case "delete":
			res.Delete = v.(bool)
		case "directiveEtx":
			res.DirectiveExt = helper.GetArrayOfInterface[string](v)
		}
	}
	return res
}
