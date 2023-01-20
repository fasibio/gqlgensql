package structure

import (
	"strings"

	"github.com/huandu/xstrings"
	"github.com/vektah/gqlparser/v2/ast"
)

type Object struct {
	Entities []Entity
	Raw      *ast.Definition
}

func NewObject(raw *ast.Definition) Object {
	return Object{
		Entities: make([]Entity, 0),
		Raw:      raw,
	}
}

func (o Object) Name() string {
	return o.Raw.Name
}

func (o Object) Many2ManyRefEntities() map[string]Entity {
	res := make(map[string]Entity)
	for _, v := range o.Entities {
		if v.HasGormDirective() && strings.Contains(v.GormDirectiveValue(), "many2many:") {
			gs := strings.Split(v.GormDirectiveValue(), ";")
			key := ""
			for _, gd := range gs {
				if strings.Contains(gd, "many2many:") {
					key = strings.Replace(gd, "many2many:", "", 1)
				}
			}
			res[key] = v
		}
	}
	return res
}

func (o Object) HasSqlDirective() bool {
	return o.SQLDirective() != nil
}

func (o Object) PrimaryKeyField() *Entity {
	for _, v := range o.Entities {
		if v.IsPrimary() {
			return &v
		}
	}
	return nil
}

func (o Object) ForeignNameKeyName(fieldName string) string {
	foreignName := xstrings.ToSnakeCase(fieldName + "ID")

out:
	for _, e := range o.Entities {
		if e.Name() == fieldName && e.HasGormDirective() {
			v := e.GormDirectiveValue()
			if strings.Contains(v, "foreignKey:") {
				commands := strings.Split(v, ";")
				for _, c := range commands {
					foreignName = strings.Split(c, ":")[1]
					break out
				}
			}
			if strings.Contains(v, "many2many:") {
				foreignName = ""
			}

			break out
		}
	}
	return foreignName
}

func (o Object) SQLDirective() *SQLDirective {
	if directive := o.Raw.Directives.ForName(string(DirectiveSQL)); directive != nil {
		res := SQLDirective{}
		qa := directive.Arguments.ForName(DirectiveSQLArgumentQuery)
		if qa == nil {
			res.Query = getDefaultFilledSqlBuilderQuery(true)
		} else {
			res.Query = customizeSqlBuilderQuery(qa)
		}
		ma := directive.Arguments.ForName(DirectiveSQLArgumentMutation)
		if ma == nil {
			res.Mutation = getDefaultFilledSqlBuilderMutation(true)
		} else {
			res.Mutation = customizeSqlBuilderMutation(ma)
		}
		return &res
	}
	return nil
}

func (o Object) PrimaryKeys() []Entity {
	res := make([]Entity, 0)
	for _, e := range o.Entities {
		if e.IsPrimary() {
			res = append(res, e)
		}
	}
	return res
}

func (o Object) WhereAbleEntities() []Entity {
	res := make([]Entity, 0)
	for _, e := range o.Entities {
		if e.WhereAble() {
			res = append(res, e)
		}
	}
	return res
}

func (o Object) OrderAbleEntities() []Entity {
	res := make([]Entity, 0)
	for _, e := range o.Entities {
		if e.OrderAble() {
			res = append(res, e)
		}
	}
	return res
}
