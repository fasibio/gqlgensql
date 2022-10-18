package gqlgensqlplugin

import (
	"fmt"
	"log"

	"github.com/99designs/gqlgen/codegen/config"
)

func (ggs *GqlGenSqlPlugin) MutateConfig(cfg *config.Config) error {
	log.Println("MutateConfig")
	cfg.Directives[DirectiveSQL] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives[DirectiveSQLPrimary] = config.DirectiveConfig{SkipRuntime: true}

	for k := range ggs.handler.List {
		makeResolverFor := []string{fmt.Sprintf("Add%sPayload", k), fmt.Sprintf("Update%sPayload", k), fmt.Sprintf("Delete%sPayload", k)}
		for _, r := range makeResolverFor {
			e := cfg.Models[r]
			e.Fields = make(map[string]config.TypeMapField)
			e.Fields[k] = config.TypeMapField{
				Resolver: true,
			}
			cfg.Models[r] = e
		}

	}

	return nil
}
