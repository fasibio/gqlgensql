
{{ reserveImport "gorm.io/gorm" }}
{{ reserveImport "context" }}
{{- range $import := .Imports }}
	{{ reserveImport $import }}
{{- end}}
{{- $root := .}}

type GqlGenSqlHookM interface {
	{{$root.HookList "model." "" | join "|"}}
}
type GqlGenSqlHookF interface {
	{{$root.HookList "model." "FiltersInput" | join "|"}}
}

type GqlGenSqlHookQueryO interface {
	{{$root.HookList "model." "Order" | join "|"}}
}

type GqlGenSqlHookI interface {
	{{$root.HookList "model." "Input" | join "|"}}
}

type GqlGenSqlHookU interface {
  {{$root.HookList "model.Update" "Input" | join "|"}}
}

type GqlGenSqlHookUP interface {
  {{$root.HookList "model.Update" "Payload" | join "|"}}
}

type GqlGenSqlHookDP interface {
  {{$root.HookList "model.Delete" "Payload" | join "|"}}
}


type GqlGenSqlHookAP interface {
{{$root.HookList "model.Add" "Payload" | join "|"}}
}


type GqlGenSqlDB struct {
	Db *gorm.DB
	Hooks map[string]any
}
func NewGqlGenSqlDB(db *gorm.DB) GqlGenSqlDB {
	return GqlGenSqlDB{
		Db:    db,
		Hooks: make(map[string]any),
	}
}

func (db *GqlGenSqlDB) Init() {
	db.Db.AutoMigrate({{.ModelsMigrations}})
}

func AddGetHook[T GqlGenSqlHookM](db *GqlGenSqlDB, name string, implementation GqlGenSqlHookGet[T]) {
	db.Hooks[name] = implementation
}

func AddQueryHook[M GqlGenSqlHookM, F GqlGenSqlHookF, O GqlGenSqlHookQueryO](db *GqlGenSqlDB, name string, implementation GqlGenSqlHookQuery[M, F, O]) {
	db.Hooks[name] = implementation
}

func AddAddHook[M GqlGenSqlHookM,I GqlGenSqlHookI, AP GqlGenSqlHookAP](db *GqlGenSqlDB, name string, implementation GqlGenSqlHookAdd[M, I, AP]) {
	db.Hooks[name] = implementation
}

func AddUpdateHook[M GqlGenSqlHookM, U GqlGenSqlHookU, UP GqlGenSqlHookUP](db *GqlGenSqlDB, name string, implementation GqlGenSqlHookUpdate[M, U, UP]) {
	db.Hooks[name] = implementation
}

func AddDeleteHook[M GqlGenSqlHookM, F GqlGenSqlHookF, DP GqlGenSqlHookDP](db *GqlGenSqlDB, name string, implementation GqlGenSqlHookDelete[M, F, DP]) {
	db.Hooks[name] = implementation
}

type GqlGenSqlHookGet[obj GqlGenSqlHookM] interface {
	Received(ctx context.Context, dbHelper *GqlGenSqlDB, id int) (*gorm.DB, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	AfterCallDb(ctx context.Context, data *obj) (*obj, error)
	BeforeReturn(ctx context.Context, data *obj, db *gorm.DB) (*obj, error)
}

type GqlGenSqlHookQuery[obj GqlGenSqlHookM, filter GqlGenSqlHookF, order GqlGenSqlHookQueryO] interface {
	Received(ctx context.Context, dbHelper *GqlGenSqlDB, filter *filter, order *order, first, offset *int) (*gorm.DB, *filter, *order, *int, *int, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	AfterCallDb(ctx context.Context, data []*obj) ([]*obj, error)
	BeforeReturn(ctx context.Context, data []*obj, db *gorm.DB) ([]*obj, error)
}

type GqlGenSqlHookAdd[obj GqlGenSqlHookM, input GqlGenSqlHookI, res GqlGenSqlHookAP] interface {
	Received(ctx context.Context, dbHelper *GqlGenSqlDB, input []*input) (*gorm.DB, []*input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB, data []obj) (*gorm.DB,[]obj, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}

type GqlGenSqlHookUpdate[obj GqlGenSqlHookM, input GqlGenSqlHookU,  res GqlGenSqlHookUP]interface{
	Received(ctx context.Context, dbHelper *GqlGenSqlDB, input *input) (*gorm.DB, input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB, data *obj) (*gorm.DB, *obj, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}

type GqlGenSqlHookDelete[obj GqlGenSqlHookM, input GqlGenSqlHookF, res GqlGenSqlHookDP] interface {
	Received(ctx context.Context, dbHelper *GqlGenSqlDB, input *input) (*gorm.DB, input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}