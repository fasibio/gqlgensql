package gqlgensqlplugin

import (
	"log"

	"github.com/99designs/gqlgen/codegen"
)

func (ggs GqlGenSqlPlugin) GenerateCode(cfg *codegen.Data) error {
	log.Println("GenerateCode")

	return nil
}
