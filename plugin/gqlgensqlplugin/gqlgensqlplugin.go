package gqlgensqlplugin

const (
	DirectiveSQL        = "SQL"
	DirectiveSQLPrimary = "SQL_PRIMARY"
	ArgumentQuery       = "query"
)

type GqlGenSqlPlugin struct {
}

func (ggs GqlGenSqlPlugin) Name() string {
	return "gqlgensql"
}
