// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package db

import (
	"github.com/fasibio/gqlgensql/graph/model"
	"gorm.io/gorm"
)

type GqlGenSqlDB struct {
	Db *gorm.DB
}

func NewGqlGenSqlDB(db *gorm.DB) GqlGenSqlDB {
	return GqlGenSqlDB{db}
}

func (db *GqlGenSqlDB) Init() {
	db.Db.AutoMigrate(&model.User{}, &model.Company{})
}
