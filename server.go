package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fasibio/gqlgensql/graph"
	"github.com/fasibio/gqlgensql/graph/db"
	"github.com/fasibio/gqlgensql/graph/generated"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultPort = "9999"

func main() {
	dbCon, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbCon = dbCon.Debug()
	dborm := db.NewGqlGenSqlDB(dbCon)
	dborm.Init()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Sql: &dborm}, Directives: generated.DirectiveRoot{
		Validated: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
			return next(ctx)
		},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
