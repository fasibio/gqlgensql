package gqlgensqlplugin

import (
	"encoding/json"
	"log"
	"os"

	"github.com/99designs/gqlgen/codegen/config"
)

func (ggs GqlGenSqlPlugin) MutateConfig(cfg *config.Config) error {
	log.Println("MutateConfig")
	b, _ := json.Marshal(cfg)
	os.WriteFile("./lalala.json", b, 0644)
	return nil
}
