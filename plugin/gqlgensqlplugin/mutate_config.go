package gqlgensqlplugin

import (
	"log"

	"github.com/99designs/gqlgen/codegen/config"
)

func (ggs GqlGenSqlPlugin) MutateConfig(cfg *config.Config) error {
	log.Println("MutateConfig")
	cfg.Resolver.
	return nil
}
