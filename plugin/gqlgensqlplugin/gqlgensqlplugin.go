package gqlgensqlplugin

import "log"

const (
	DirectiveSQL        = "SQL"
	DirectiveSQLPrimary = "SQL_PRIMARY"
	ArgumentQuery       = "query"
)

type GqlGenSqlPlugin struct {
	handler SqlBuilderHandler
}

func NewGqlGenSqlPlugin() GqlGenSqlPlugin {
	log.Println("asddsa")
	return GqlGenSqlPlugin{}
}

func (ggs *GqlGenSqlPlugin) Name() string {
	return "gqlgensql"
}
