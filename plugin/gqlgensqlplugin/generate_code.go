package gqlgensqlplugin

import (
	_ "embed"
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/templates"
)

//go:embed generate_code.go.tpl
var generateCodeTemplate string

//go:embed generate_code_db.go.tpl
var generateDbCodeTemplate string

func (ggs *GqlGenSqlPlugin) GenerateCode(data *codegen.Data) error {
	// filename := "blatest123.go"
	log.Println("GenerateCode")
	ggs.generateDbCode(data)
	// return templates.Render(templates.Options{
	// 	PackageName: "main",
	// 	Filename:    filename,
	// 	Data: &ResolverBuild{
	// 		Data:     cfg,
	// 		TypeName: "BlaCa",
	// 	},
	// 	GeneratedHeader: true,
	// 	Packages:        cfg.Config.Packages,
	// 	Template:        generateCodeTemplate,
	// })

	// for _, o := range data.Objects {
	// 	for _, f := range o.Fields {
	// 		if !f.IsResolver {
	// 			continue
	// 		}

	// 		f.im
	// 	}
	// }
	return nil
}

type DbCodeData struct {
	Data    *codegen.Data
	Handler SqlBuilderHandler
}

func (db *DbCodeData) Imports() string {
	importMap := make(map[string]bool)
	for _, v := range db.Handler.List {

		sp := strings.LastIndex(db.Data.Config.Models[v.TypeName].Model[0], ".")
		importMap[db.Data.Config.Models[v.TypeName].Model[0][:sp]] = true
	}
	for k := range importMap {
		templates.CurrentImports.Reserve(k)
		log.Println("hier", k)
		return k
	}
	return ""
}

func (db *DbCodeData) ModelsMigrations() string {
	res := ""
	for _, v := range db.Handler.List {
		splits := strings.Split(db.Data.Config.Models[v.TypeName].Model[0], "/")
		res += fmt.Sprintf("&%s{},", splits[len(splits)-1])
	}
	return res[:len(res)-1]
}

func (ggs GqlGenSqlPlugin) generateDbCode(data *codegen.Data) error {

	filename := path.Join(data.Config.Resolver.Package, "db/db_gen.go")
	return templates.Render(templates.Options{
		PackageName: "db",
		Filename:    filename,
		Data: &DbCodeData{
			Data:    data,
			Handler: ggs.handler,
		},
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
		Template:        generateDbCodeTemplate,
	})
}

type ResolverBuild struct {
	*codegen.Data

	TypeName string
}
