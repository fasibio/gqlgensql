
{{ reserveImport "gorm.io/gorm" }}


{{ reserveImport .Imports }}

type GqlGenSqlDB struct {
	db *gorm.DB
}

func NewGqlGenSqlDB(db *gorm.DB) GqlGenSqlDB {
	return GqlGenSqlDB{db}
}

func (db *GqlGenSqlDB) Init() {
	db.db.AutoMigrate({{.ModelsMigrations}})
}
