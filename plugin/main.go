package main

import (
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/fasibio/gqlgensql/plugin/gqlgensql"
)

func main() {
	log.Println("da")

	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	sqlPlugin, muateHookPlugin := gqlgensql.NewGqlGenSqlPlugin()

	err = api.Generate(cfg, api.AddPlugin(sqlPlugin), api.ReplacePlugin(muateHookPlugin))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
