package domain

import (
	"context"

	"gorm.io/gorm"
)

type Database interface {
	// Necessary *gorm.DB methods
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
	Count(context.Context, interface{}) (int, error)
}