package model

import (
	"fmt"
	"strings"

	"github.com/fasibio/gqlgensql/plugin/gqlgensql/runtimehelper"
	"gorm.io/gorm"
)

func (d *CompanyFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias)...))
	}

	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf("%s.%s", alias, "id"))...)
	}

	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf("%s.%s", alias, "name"))...)
	}

	if d.MotherCompanyID != nil {
		res = append(res, d.MotherCompanyID.ExtendsDatabaseQuery(db, fmt.Sprintf("%s.%s", alias, "mother_company_id"))...)
	}

	if d.MotherCompany != nil {
		tableName := db.Config.NamingStrategy.TableName("MotherCompany")
		db := db.Joins(fmt.Sprintf("JOIN %s ON %s.id = %s.company_id", tableName, tableName, alias))
		res = append(res, d.MotherCompany.ExtendsDatabaseQuery(db, tableName)...)
	}

	return res
}

func (d *UserFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias)...))
	}

	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf("%s.%s", alias, "id"))...)
	}

	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf("%s.%s", alias, "name"))...)
	}

	if d.CompanyID != nil {
		res = append(res, d.CompanyID.ExtendsDatabaseQuery(db, fmt.Sprintf("%s.%s", alias, "company_id"))...)
	}

	if d.Company != nil {
		tableName := db.Config.NamingStrategy.TableName("Company")
		db := db.Joins(fmt.Sprintf("JOIN %s ON %s.id = %s.company_id", tableName, tableName, alias))
		res = append(res, d.Company.ExtendsDatabaseQuery(db, tableName)...)
	}

	return res
}

func (d *StringFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}
	if d.Contains != nil {
		res = append(res, runtimehelper.Like(fieldName, fmt.Sprintf("%%%s%%", *d.Contains)))
	}

	if d.Containsi != nil {
		res = append(res, runtimehelper.Like(fmt.Sprintf("lower(%s)", fieldName), fmt.Sprintf("%%%s%%", strings.ToLower(*d.Containsi))))
	}

	if d.EndsWith != nil {
		res = append(res, runtimehelper.Like(fieldName, fmt.Sprintf("%%%s", *d.EndsWith)))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}

	if d.Eqi != nil {
		res = append(res, runtimehelper.Equal(fmt.Sprintf("lower(%s)", fieldName), strings.ToLower(*d.Eqi)))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName)...))
	}

	if d.NotContains != nil {
		res = append(res, runtimehelper.NotLike(fieldName, fmt.Sprintf("%%%s%%", *d.NotContains)))
	}

	if d.NotContainsi != nil {
		res = append(res, runtimehelper.NotLike(fmt.Sprintf("lower(%s)", fieldName), fmt.Sprintf("%%%s%%", strings.ToLower(*d.NotContainsi))))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName, d.NotIn))
	}

	if d.NotNull != nil {
		res = append(res, runtimehelper.NotNull(fieldName, d.NotNull))
	}

	if d.Null != nil {
		res = append(res, runtimehelper.Null(fieldName, d.Null))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.StartsWith != nil {
		res = append(res, runtimehelper.Like(fieldName, fmt.Sprintf("%s%%", *d.StartsWith)))
	}

	return res
}

func (d *IntFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string) []runtimehelper.ConditionElement {

	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Between != nil {
		res = append(res, runtimehelper.Between(fieldName, d.Between.Start, d.Between.End))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}
	if d.Gt != nil {
		res = append(res, runtimehelper.More(fieldName, *d.Gt))
	}

	if d.Gte != nil {
		res = append(res, runtimehelper.MoreOrEqual(fieldName, *d.Gte))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Lt != nil {
		res = append(res, runtimehelper.Less(fieldName, *d.Lt))
	}

	if d.Lte != nil {
		res = append(res, runtimehelper.LessOrEqual(fieldName, *d.Lte))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}
	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName)...))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName, d.NotIn))

	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName, *d.NotNull))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName, *d.Null))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}

func (d *BooleanFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Is != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Is))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName)...))
	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName, *d.NotNull))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName, *d.Null))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}
