package gqlgensqlplugin

import (
	"fmt"
	"log"

	"github.com/vektah/gqlparser/v2/ast"
)

func (ggs *GqlGenSqlPlugin) InjectSourceEarly() *ast.Source {
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
