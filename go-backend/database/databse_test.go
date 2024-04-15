package database

import (
	"context"
	"testing"
	"time"

	"go-backend/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDB() (domain.Database, error){
	// Connect to empty test database
	dsn := "host=localhost user=postgres password=postgres dbname=forumdb_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Set transaction begin
	tx := db.Begin()

	// Do database migration
	err = tx.AutoMigrate(&domain.User{})

	return NewDatabaseFromExist(tx), err
}

func teardownDB(db domain.Database) {
	db.Rollback()
}

func TestBasicCRUD(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Errorf("Database setup error: %v", err)
	}

	defer teardownDB(db)

	t.Run("InsertOne", func (t *testing.T) {
		result := db.SavePoint("InsertOneBegin")
		if result.Error != nil {
			t.Errorf("Fail in SavePoint: %v", err)
		}

		defer db.Rollbackto("InsertOneBegin")
		
		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
		defer cancel()

		user := domain.User{
			Email: "test@test.com",
			PasswordHash: []byte("test value"),
		}
		// InsertOne() should insert user to table and assign back to default primary key ID
		err := db.InsertOne(ctx, &user)
		if err != nil {
			t.Errorf("InsertOne() fail: %v", err)
		}

		var dest = domain.User{ID: user.ID}
		db.First(&dest)

		if !dest.Equals(user) {
			t.Errorf("InsertOne() fail: insert value not match search result.")
		}
	})
}