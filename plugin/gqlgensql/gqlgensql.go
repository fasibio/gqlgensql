package gqlgensql

import (
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/fasibio/gqlgensql/plugin/gqlgensql/structure"
)

type GqlGenSqlPlugin struct {
	Handler structure.SqlBuilderHelper
}

func NewGqlGenSqlPlugin() (*GqlGenSqlPlugin, *modelgen.Plugin) {
	sp := &GqlGenSqlPlugin{}
	modelGenPlugin := &modelgen.Plugin{MutateHook: MutateHook(sp), FieldHook: ConstraintFieldHook(sp)}
	return sp, modelGenPlugin
}

func (ggs *GqlGenSqlPlugin) Name() string {
	return "gqlgensql"
}
