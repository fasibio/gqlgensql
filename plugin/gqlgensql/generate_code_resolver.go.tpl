
{{ reserveImport "context"  }}
{{ reserveImport "fmt"  }}
{{ reserveImport "gorm.io/gorm/clause"}}
{{ range $import := .Imports }}
	{{ reserveImport $import }}
{{end}}

{{$root := .}}
{{- range $objectName, $object := .Handler.List.Objects}}
{{- if $object.HasSqlDirective}}
  {{- if $object.SQLDirective.HasQueries}}
    {{- if $object.SQLDirective.Query.Get}}
// Get{{$object.Name}} is the resolver for the get{{$object.Name}} field.
{{$primaryField := $object.PrimaryKeyField }}
func (r *queryResolver) Get{{$object.Name}}(ctx context.Context, {{$primaryField.Name}} {{$root.GetGoFieldType $objectName $primaryField false}}) (*model.{{$object.Name}}, error) {
	var res model.{{$object.Name}}
  d := runtimehelper.GetPreloadSelection(ctx,r.Sql.Db,runtimehelper.GetPreloadsMap(ctx)).First(&res, id)
	return &res, d.Error
}
    {{- end}}
    {{- if $object.SQLDirective.Query.Query}}
// Query{{$object.Name}} is the resolver for the query{{$object.Name}} field.
func (r *queryResolver) Query{{$object.Name}}(ctx context.Context, filter *model.{{$object.Name}}FiltersInput, order *model.{{$object.Name}}Order, first *int, offset *int) (*model.{{$object.Name}}QueryResult, error) {
	var res []*model.{{$object.Name}}
  tableName := r.Sql.Db.Config.NamingStrategy.TableName("{{$object.Name}}")
	db := runtimehelper.GetPreloadSelection(ctx, r.Sql.Db,runtimehelper.GetPreloadsMap(ctx).SubTables[0])
	sql, arguments := runtimehelper.CombineSimpleQuery(filter.ExtendsDatabaseQuery(db, tableName), "OR")
	db.Where(sql, arguments...)
	
	if (order != nil){
		if order.Asc != nil {
			db = db.Order(fmt.Sprintf("%s.%s asc",tableName,order.Asc))
		}
		if order.Desc != nil {
			db = db.Order(fmt.Sprintf("%s.%s desc",tableName,order.Desc))
		}
	}
	var total int64
	db.Model(res).Count(&total)
	if first != nil {
		db = db.Limit(*first)
	}
	if offset != nil {
		db = db.Offset(*offset)
	}
	d := db.Find(&res)
	return &model.{{$object.Name}}QueryResult{
		Data: res,
    Count: len(res),
		TotalCount: int(total),
	},d.Error
}
    {{- end}}
  {{- end}}
  {{- if $object.SQLDirective.HasMutation}}
    {{- if $object.SQLDirective.Mutation.Add}}
// Add{{$object.Name}} is the resolver for the add{{$object.Name}} field.
func (r *mutationResolver) Add{{$object.Name}}(ctx context.Context, input []*model.{{$object.Name}}Input) (*model.Add{{$object.Name}}Payload, error) {
	obj:= make([]model.{{$object.Name}}, len(input))
  for i, v := range input {
    obj[i] = v.MergeToType()
  }
  res := r.Sql.Db.Omit(clause.Associations).Create(&obj)
  return &model.Add{{$object.Name}}Payload{}, res.Error
}
    {{- end}}
    {{- if $object.SQLDirective.Mutation.Update}}
// Update{{$object.Name}} is the resolver for the update{{$object.Name}} field.
func (r *mutationResolver) Update{{$object.Name}}(ctx context.Context, input model.Update{{$object.Name}}Input) (*model.Update{{$object.Name}}Payload, error) {
	panic(fmt.Errorf("not implemented: Update{{$object.Name}} - update{{$object.Name}}"))
}
    {{- end}}
    {{- if $object.SQLDirective.Mutation.Delete}}
// Delete{{$object.Name}} is the resolver for the delete{{$object.Name}} field.
func (r *mutationResolver) Delete{{$object.Name}}(ctx context.Context, filter model.{{$object.Name}}FiltersInput) (*model.Delete{{$object.Name}}Payload, error) {
	panic(fmt.Errorf("not implemented: Delete{{$object.Name}} - delete{{$object.Name}}"))
}
    {{- end}}
  {{- end}}
{{- end}}
{{- end}}