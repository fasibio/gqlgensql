
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
  d := runtimehelper.GetPreloadSelection(ctx,r.Sql.Db,runtimehelper.GetPreloadsMap(ctx,"{{$object.Name}}")).First(&res, id)
	return &res, d.Error
}
    {{- end}}
    {{- if $object.SQLDirective.Query.Query}}
// Query{{$object.Name}} is the resolver for the query{{$object.Name}} field.
func (r *queryResolver) Query{{$object.Name}}(ctx context.Context, filter *model.{{$object.Name}}FiltersInput, order *model.{{$object.Name}}Order, first *int, offset *int) (*model.{{$object.Name}}QueryResult, error) {
	var res []*model.{{$object.Name}}
  tableName := r.Sql.Db.Config.NamingStrategy.TableName("{{$object.Name}}")
	db := runtimehelper.GetPreloadSelection(ctx, r.Sql.Db,runtimehelper.GetPreloadsMap(ctx, "data").SubTables[0])
	if filter != nil{
		sql, arguments := runtimehelper.CombineSimpleQuery(filter.ExtendsDatabaseQuery(db, tableName), "OR")
		db.Where(sql, arguments...)
	}

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
func (r *Resolver) Add{{$object.Name}}Payload() generated.Add{{$object.Name}}PayloadResolver { return &{{lcFirst $object.Name}}PayloadResolver[*model.Add{{$object.Name}}Payload]{r} }
func (r *Resolver) Delete{{$object.Name}}Payload() generated.Delete{{$object.Name}}PayloadResolver { return &{{lcFirst $object.Name}}PayloadResolver[*model.Delete{{$object.Name}}Payload]{r} }
func (r *Resolver) Update{{$object.Name}}Payload() generated.Update{{$object.Name}}PayloadResolver { return &{{lcFirst $object.Name}}PayloadResolver[*model.Update{{$object.Name}}Payload]{r} }


type {{lcFirst $object.Name}}Payload interface {
	*model.Add{{$object.Name}}Payload | *model.Delete{{$object.Name}}Payload | *model.Update{{$object.Name}}Payload
}

type {{lcFirst $object.Name}}PayloadResolver[T  {{lcFirst $object.Name}}Payload] struct {
	*Resolver
}
func (r *{{lcFirst $object.Name}}PayloadResolver[T]) {{$object.Name}}(ctx context.Context, obj T, filter *model.{{$object.Name}}FiltersInput, order *model.{{$object.Name}}Order, first *int, offset *int) (*model.{{$object.Name}}QueryResult, error){
	return r.Query().Query{{$object.Name}}(ctx,filter,order,first,offset)
}
		{{- range $m2mKey, $m2mEntity := $object.Many2ManyRefEntities }}
func (r *mutationResolver) Add{{$m2mEntity.GqlTypeName}}2{{$object.Name}}s(ctx context.Context, input model.{{$m2mEntity.GqlTypeName}}Ref2{{$object.Name}}sInput) (*model.Update{{$object.Name}}Payload, error){
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("{{$object.Name}}")
	sql, arguments := runtimehelper.CombineSimpleQuery(input.Filter.ExtendsDatabaseQuery(r.Sql.Db, tableName), "OR")
	db := r.Sql.Db.Model(&model.{{$object.Name}}{}).Where(sql, arguments...)
	var res []*model.{{$object.Name}}
	db.Find(&res)
	{{- $table1ID := $root.GetGoFieldName $object.Name $object.PrimaryKeyField }}
	{{- $tabe2PrimaryEntity := $root.PrimaryKeyEntityOfObject $m2mEntity.GqlTypeName}}
	{{- $table2ID := $root.GetGoFieldName $m2mEntity.GqlTypeName $tabe2PrimaryEntity }}
	type {{camelcase $m2mKey}} struct {
		{{ucFirst $object.Name}}{{$table1ID}} {{$root.GetGoFieldType $object.Name $object.PrimaryKeyField false}} 
		{{ucFirst $m2mEntity.GqlTypeName}}{{$table2ID}} {{$root.GetGoFieldType $m2mEntity.GqlTypeName $tabe2PrimaryEntity false}} 
	}
	resIds := make([]map[string]interface{}, 0)
	for _, v := range res{
		for _, v1 := range input.Set {
				tmp := make(map[string]interface{})
				tmp["{{ucFirst $object.Name}}{{$table1ID}}"] = v.ID
				tmp["{{ucFirst $m2mEntity.GqlTypeName}}{{$table2ID}}"] = v1
				resIds = append(resIds, tmp)
		}	
	}
	d := r.Sql.Db.Model(&{{camelcase $m2mKey}}{}).Create(resIds)
	return &model.Update{{$object.Name}}Payload{
		Count: int(d.RowsAffected),
	},d.Error
}
		{{- end}}
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
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("{{$object.Name}}")
	sql, arguments := runtimehelper.CombineSimpleQuery(input.Filter.ExtendsDatabaseQuery(r.Sql.Db, tableName), "OR")
	obj := model.{{$object.Name}}{}
	res := r.Sql.Db.Model(&obj).Where(sql, arguments...).Updates(input.Set.MergeToType())
	return &model.Update{{$object.Name}}Payload{
		Count: int(res.RowsAffected),
	}, res.Error
}
    {{- end}}
    {{- if $object.SQLDirective.Mutation.Delete}}
// Delete{{$object.Name}} is the resolver for the delete{{$object.Name}} field.
func (r *mutationResolver) Delete{{$object.Name}}(ctx context.Context, filter model.{{$object.Name}}FiltersInput) (*model.Delete{{$object.Name}}Payload, error) {
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("{{$object.Name}}")
	sql, arguments := runtimehelper.CombineSimpleQuery(filter.ExtendsDatabaseQuery(r.Sql.Db, tableName), "OR")
	obj := model.{{$object.Name}}{}
	res := r.Sql.Db.Where(sql, arguments...).Delete(&obj)
	msg := fmt.Sprintf("%d rows deleted",res.RowsAffected)
	return &model.Delete{{$object.Name}}Payload{
		Count: int(res.RowsAffected),
		Msg: &msg,
	}, res.Error
}
    {{- end}}
  {{- end}}
{{- end}}
{{- end}}