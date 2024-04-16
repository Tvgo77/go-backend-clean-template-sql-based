package domain

import (
	"context"


	"database/sql"
	"gorm.io/gorm"
)

type Database interface {
	// Necessary *gorm.DB methods
	AutoMigrate(dest ...interface{}) (error)
	Begin(opts ...*sql.TxOptions) *gorm.DB
	SavePoint(name string) *gorm.DB
	Rollbackto(name string) *gorm.DB
	Rollback() *gorm.DB

	Create(value interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Select(query interface{}, args ...interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	WithContext(ctx context.Context) (tx *gorm.DB)
	// ...
	
	Ping() (error)
	InsertOne(context.Context, interface{}) (error)
	FindOne(context.Context, interface{}, interface{}) (error)
	UpdateOne(context.Context, interface{}, interface{}) (error)
	DeleteOne(context.Context, interface{}) (error)
	CountRows(context.Context, interface{}) (int, error)
}