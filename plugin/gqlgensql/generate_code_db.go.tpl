
{{ reserveImport "gorm.io/gorm" }}

{{ range $import := .Imports }}
	{{ reserveImport $import }}
{{end}}
type GqlGenSqlDB struct {
	Db *gorm.DB
}

func NewGqlGenSqlDB(db *gorm.DB) GqlGenSqlDB {
	return GqlGenSqlDB{db}
}

func (db *GqlGenSqlDB) Init() {
	db.Db.AutoMigrate({{.ModelsMigrations}})
}


