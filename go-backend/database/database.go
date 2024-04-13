package database

import (
	"context"
	"go-backend/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	db *gorm.DB
}

func NewDatabase(dsn string) (domain.Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &postgresDB{db: db}, nil
}

func (p *postgresDB) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return p.db.First(dest, conds)
}

func (p *postgresDB) Select(query interface{}, args ...interface{}) (tx *gorm.DB) {
	return p.db.Select(query, args)
}

func (p *postgresDB) Where(query interface{}, args ...interface{}) (tx *gorm.DB) {
	return p.db.Where(query, args)
}

func (p *postgresDB) WithContext(ctx context.Context) (tx *gorm.DB) {
	return p.db.WithContext(ctx)
}

/* Custom CRUD convinent interface */
// Note: Only non-zero of struct field will be used for condition and updates
// If you want to deal with zero value, use map[string]interface{}{} as input args

func (p *postgresDB) Ping() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (p *postgresDB) InsertOne(ctx context.Context, src interface{}) error {
	result := p.db.WithContext(ctx).Create(src)
	return result.Error
}

// Query row goes into dest
func (p *postgresDB) FindOne(ctx context.Context, dest interface{}, conds interface{}) error {
	result := p.db.WithContext(ctx).Where(conds).First(dest)
	return result.Error
}


// Old one should contain primary key field
func (p *postgresDB) UpdateOne(ctx context.Context, old interface{}, new interface{}) error {
	result := p.db.WithContext(ctx).Model(old).Updates(new)
	return result.Error
}

// Arg should contain primary key field
func (p *postgresDB) DeleteOne(ctx context.Context, arg interface{}) error {
	result := p.db.WithContext(ctx).Delete(arg)
	return result.Error
}

