package gqlgensql

import (
	"fmt"
	"log"

	"github.com/fasibio/gqlgensql/plugin/gqlgensql/structure"
	"github.com/vektah/gqlparser/v2/ast"
)

func (ggs *GqlGenSqlPlugin) InjectSourceEarly() *ast.Source {
	log.Println("InjectSourceEarly")

	input := fmt.Sprintf(`

	input SqlMutationParams {
		add: Boolean
		update: Boolean
		delete: Boolean
		directiveEtx: [String!]
	}

	input SqlQueryParams {
		get: Boolean
		query: Boolean
		directiveEtx: [String!]
	}
	directive @%s(%s:SqlQueryParams, %s: SqlMutationParams ) on OBJECT
	directive @%s on FIELD_DEFINITION
	directive @%s on FIELD_DEFINITION

	directive @%s (value: String)on FIELD_DEFINITION
	`, structure.DirectiveSQL,
		structure.DirectiveSQLArgumentQuery,
		structure.DirectiveSQLArgumentMutation,
		structure.DirectiveSQLPrimary,
		structure.DirectiveSQLIndex,
		structure.DirectiveSQLGorm)

	return &ast.Source{
		Name:    fmt.Sprintf("%s/directive.graphql", ggs.Name()),
		Input:   input,
		BuiltIn: true,
	}
}
