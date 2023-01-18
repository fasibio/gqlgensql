{{ reserveImport "fmt"  }}
{{ reserveImport "gorm.io/gorm"  }}
{{ reserveImport "gorm.io/gorm/clause"  }}
{{ reserveImport "strings" }}
	
{{$methodeName := "ExtendsDatabaseQuery"}}


{{- $root := .}}

{{- range $objectName, $object := .Handler.List.Objects }}
{{- if $object.HasSqlDirective}}

func (d *{{$object.Name}}FiltersInput) {{$methodeName}}(db *gorm.DB, alias string) []clause.Expression {
	res := make([]clause.Expression, 0)
    if d.And != nil {
		for _, v := range d.And{
			res = append(res, clause.And(v.ExtendsDatabaseQuery(db, alias)...))
		}
	}

	if d.Or != nil {
		for _, v := range d.Or{
			res = append(res, clause.Or(v.ExtendsDatabaseQuery(db, alias)...))
		}
	}

	if d.Not != nil {
		res = append(res, clause.Not(d.Not.ExtendsDatabaseQuery(db, alias)...))
	}
  {{- range $entityKey, $entity := $object.Entities }}
  {{- $entityGoName :=  $root.GetGoFieldName $objectName $entity}}

	if d.{{$entityGoName}} != nil {
    {{-  if $entity.BuiltIn}}
    res = append(res, d.{{$entityGoName}}.{{$methodeName}}(db, fmt.Sprintf("%s.%s",alias,"{{snakecase $entityGoName}}"))...)
    {{- else}}
    tableName := db.Config.NamingStrategy.TableName("{{$entityGoName}}")
    db := db.Joins(fmt.Sprintf("JOIN %s ON %s.{{$root.PrimaryKeyOfObject $entity.GqlTypeName | snakecase}} = %s.{{$object.ForeignNameKeyName $entity.GqlTypeName}}",tableName, tableName, alias))
    res = append(res, d.{{$entityGoName}}.{{$methodeName}}(db, tableName)...)
    {{- end}}
	}
  {{- end}}

	return res
}
{{- end}}
{{- end}}

func getExpressions(db *gorm.DB, query interface{}, args ...interface{}) []clause.Expression {
	return db.Statement.BuildCondition(query, args...)
}

func (d *StringFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string) []clause.Expression {
	res := make([]clause.Expression, 0)
	if d.And != nil {
		for _, v := range d.And {
			r := clause.And(getExpressions(db, fmt.Sprintf("%s = ?", fieldName), *v)...)
			res = append(res, r)
		}
	}
	if d.Contains != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s Like ?", fieldName), fmt.Sprintf("%%%s%%",*d.Contains))...)
		res = append(res, r)
	}

	if d.Containsi != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("lower(%s) Like ?", fieldName), fmt.Sprintf("%%%s%%",strings.ToLower(*d.Containsi)))...)
		res = append(res, r)
	}

	if d.EndsWith != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s Like ?", fieldName), fmt.Sprintf("%%%s",*d.EndsWith))...)
		res = append(res, r)
	}

	if d.Eq != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s = ?", fieldName), *d.Eq)...)
		res = append(res, r)
	}

	if d.Eqi != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("lower(%s) = ?", fieldName), strings.ToLower(*d.Eqi))...)
		res = append(res, r)
	}

	if d.In != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s in ?", fieldName), d.In)...)
		res = append(res, r)
	}

	if d.Ne != nil {
		r := clause.Not(getExpressions(db, fmt.Sprintf("%s = ?", fieldName), d.Ne)...)
		res = append(res, r)
	}

	if d.Not != nil {
		tmp := d.Not.ExtendsDatabaseQuery(db, fieldName)
		res = append(res, clause.Not(tmp...))
	}

	if d.NotContains != nil {
		r := clause.Not(getExpressions(db, fmt.Sprintf("%s Like %%?%%", fieldName), *d.NotContains)...)
		res = append(res, r)
	}

	if d.NotContainsi != nil {
		r := clause.Not(getExpressions(db, fmt.Sprintf("lower(%s) Like %%?%%", fieldName), strings.ToLower(*d.NotContainsi))...)
		res = append(res, r)
	}

	if d.NotIn != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s in ?", fieldName), d.NotIn)...)
		res = append(res, r)
	}

	if d.NotNull != nil {
		r := clause.Not(getExpressions(db, fmt.Sprintf("%s IS NULL", fieldName))...)
		res = append(res, r)
	}

	if d.Null != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s IS NULL", fieldName))...)
		res = append(res, r)
	}

	if d.Or != nil {
		for _, v := range d.Or {
			r := clause.Or(getExpressions(db, fmt.Sprintf("%s = ?", fieldName), *v)...)
			res = append(res, r)
		}
	}

	if d.StartsWith != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s Like ?", fieldName), fmt.Sprintf("%s%%",*d.StartsWith))...)
		res = append(res, r)
	}

	return res
}

func (d *IntFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string) []clause.Expression {

	res := make([]clause.Expression, 0)

	if d.And != nil {
		for _, v := range d.And {
			r := clause.And(getExpressions(db, fmt.Sprintf("%s = ?", fieldName), *v)...)
			res = append(res, r)
		}
	}

	if d.Between != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s BETWEEN ? AND ?", fieldName), d.Between.Start, d.Between.End)...)
		res = append(res, r)
	}

	if d.Eq != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s = ?", fieldName), *d.Eq)...)
		res = append(res, r)
	}
	if d.Gt != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s > ?", fieldName), *d.Gt)...)
		res = append(res, r)
	}

	if d.Gte != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s >= ?", fieldName), *d.Gte)...)
		res = append(res, r)
	}

	if d.In != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s in ?", fieldName), d.In)...)
		res = append(res, r)
	}

	if d.Lt != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s < ?", fieldName), d.Lt)...)
		res = append(res, r)
	}

	if d.Lte != nil {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s <= ?", fieldName), d.Lte)...)
		res = append(res, r)
	}

	if d.Ne != nil {
		r := clause.Not(getExpressions(db, fmt.Sprintf("%s = ?", fieldName), d.Ne)...)
		res = append(res, r)
	}
	if d.Not != nil {
		tmp := d.Not.ExtendsDatabaseQuery(db, fieldName)
		res = append(res, clause.Not(tmp...))
	}

	if d.NotIn != nil {
		r := clause.Not(getExpressions(db, fmt.Sprintf("%s in ?", fieldName), d.NotIn)...)
		res = append(res, r)
	}

	if d.NotNull != nil && *d.NotNull {
		r := clause.Not(getExpressions(db, fmt.Sprintf("%s IS NULL", fieldName))...)
		res = append(res, r)
	}

	if d.Null != nil && *d.Null {
		r := clause.And(getExpressions(db, fmt.Sprintf("%s IS NULL", fieldName))...)
		res = append(res, r)
	}

	if d.Or != nil {
		for _, v := range d.Or {
			r := clause.Or(getExpressions(db, fmt.Sprintf("%s = ?", fieldName), *v)...)
			res = append(res, r)
		}
	}

	return res
}