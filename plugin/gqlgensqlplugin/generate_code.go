package gqlgensqlplugin

import (
	_ "embed"

	"github.com/99designs/gqlgen/codegen"
)

//go:embed generate_code.go.tpl
var generateCodeTemplate string

func (ggs GqlGenSqlPlugin) GenerateCode(data *codegen.Data) error {
	// filename := "blatest123.go"
	// log.Println("GenerateCode")

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

type ResolverBuild struct {
	*codegen.Data

	TypeName string
}
