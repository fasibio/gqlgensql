package gqlgensqlplugin

const (
	DirectiveSQL        = "SQL"
	DirectiveSQLPrimary = "SQL_PRIMARY"
	ArgumentQuery       = "query"
)

type GqlGenSqlPlugin struct {
	handler SqlBuilderHandler
}

func (ggs GqlGenSqlPlugin) Name() string {
	return "gqlgensql"
}
