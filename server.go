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
	"github.com/fasibio/gqlgensql/graph/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultPort = "9999"

type DeleteCompanyHook struct{}

func (h DeleteCompanyHook) Received(ctx context.Context, dbHelper *db.GqlGenSqlDB, input *model.CompanyFiltersInput) (*gorm.DB, model.CompanyFiltersInput, error) {
	log.Println("Received")
	return dbHelper.Db, *input, nil
}
func (h DeleteCompanyHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	log.Println("BeforeCallDb")
	return db, nil
}
func (h DeleteCompanyHook) BeforeReturn(ctx context.Context, db *gorm.DB, res *model.DeleteCompanyPayload) (*model.DeleteCompanyPayload, error) {
	log.Println("BeforeReturn")
	return res, nil
}

type UpdateCompanyHook struct{}

func (h UpdateCompanyHook) Received(ctx context.Context, dbHelper *db.GqlGenSqlDB, input *model.UpdateCompanyInput) (*gorm.DB, model.UpdateCompanyInput, error) {
	log.Println("Received")
	return dbHelper.Db, *input, nil
}
func (h UpdateCompanyHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data *model.Company) (*gorm.DB, *model.Company, error) {
	log.Println("BeforeCallDb")
	return db, data, nil
}
func (h UpdateCompanyHook) BeforeReturn(ctx context.Context, db *gorm.DB, res *model.UpdateCompanyPayload) (*model.UpdateCompanyPayload, error) {
	log.Println("BeforeReturn")
	return res, nil
}

type AddCompanyHook struct{}

func (h AddCompanyHook) Received(ctx context.Context, dbHelper *db.GqlGenSqlDB, input []*model.CompanyInput) (*gorm.DB, []*model.CompanyInput, error) {
	log.Println("Received")
	return dbHelper.Db, input, nil
}
func (h AddCompanyHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data []model.Company) (*gorm.DB, []model.Company, error) {
	log.Println("BeforeCallDb")
	return db, data, nil
}
func (h AddCompanyHook) BeforeReturn(ctx context.Context, db *gorm.DB, res *model.AddCompanyPayload) (*model.AddCompanyPayload, error) {
	log.Println("BeforeReturn")
	return res, nil
}

type GetUserHook struct{}

func (h GetUserHook) Received(ctx context.Context, dbHelper *db.GqlGenSqlDB, id int) (*gorm.DB, error) {
	log.Println("Received")
	return dbHelper.Db, nil
}
func (h GetUserHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	log.Println("BeforeCallDb")
	return db, nil
}
func (h GetUserHook) AfterCallDb(ctx context.Context, data *model.User) (*model.User, error) {
	log.Println("AfterCallDb")
	return data, nil
}
func (h GetUserHook) BeforeReturn(ctx context.Context, data *model.User, db *gorm.DB) (*model.User, error) {
	log.Println("BeforeReturn")
	return data, nil
}

type QueryCompanyHook struct{}

func (h QueryCompanyHook) Received(ctx context.Context, dbHelper *db.GqlGenSqlDB, filter *model.CompanyFiltersInput, order *model.CompanyOrder, first, offset *int) (*gorm.DB, *model.CompanyFiltersInput, *model.CompanyOrder, *int, *int, error) {
	log.Println("Received")
	return dbHelper.Db, filter, order, first, offset, nil
}
func (h QueryCompanyHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	log.Println("BeforeCallDb")
	return db, nil
}
func (h QueryCompanyHook) AfterCallDb(ctx context.Context, data []*model.Company) ([]*model.Company, error) {
	log.Println("AfterCallDb")
	return data, nil
}
func (h QueryCompanyHook) BeforeReturn(ctx context.Context, data []*model.Company, db *gorm.DB) ([]*model.Company, error) {
	log.Println("BeforeReturn")
	return data, nil
}

func main() {
	dbCon, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbCon = dbCon.Debug()
	dborm := db.NewGqlGenSqlDB(dbCon)
	db.AddGetHook[model.User](&dborm, "GetUser", GetUserHook{})
	db.AddQueryHook[model.Company, model.CompanyFiltersInput, model.CompanyOrder](&dborm, "QueryCompany", QueryCompanyHook{})
	db.AddAddHook[model.Company, model.CompanyInput, model.AddCompanyPayload](&dborm, "AddCompany", AddCompanyHook{})
	db.AddUpdateHook[model.Company, model.UpdateCompanyInput, model.UpdateCompanyPayload](&dborm, "UpdateCompany", UpdateCompanyHook{})
	db.AddDeleteHook[model.Company, model.CompanyFiltersInput, model.DeleteCompanyPayload](&dborm, "DeleteCompany", DeleteCompanyHook{})
	dborm.Init()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Sql: &dborm}, Directives: generated.DirectiveRoot{
		Validated: DirectiveImpl,
		A:         DirectiveImpl,
		B:         DirectiveImpl,
		C:         DirectiveImpl,
		D:         DirectiveImpl,
		E:         DirectiveImpl,
		F:         DirectiveImpl,
		G:         DirectiveImpl,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func DirectiveImpl(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	return next(ctx)
}
