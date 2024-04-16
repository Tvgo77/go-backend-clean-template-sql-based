package database

import (
	"context"
	"testing"
	"time"

	"go-backend/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDB() (domain.Database, error) {
	// Connect to empty test database
	dsn := "host=localhost user=postgres dbname=forumdb_test port=5432 sslmode=disable"
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
		t.Fatalf("Database setup error: %v", err)
	}

	defer teardownDB(db)

	t.Run("Ping", func (t *testing.T){
		err := db.Ping()
		if err != nil {
			t.Fatalf("Ping() fail in connected DB")
		}
	})

	t.Run("InsertOne", func(t *testing.T) {
		result := db.SavePoint("InsertOneBegin")
		if result.Error != nil {
			t.Fatalf("Fail in SavePoint: %v", err)
		}

		defer db.Rollbackto("InsertOneBegin")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		user := domain.User{
			Email:        "test@test.com",
			PasswordHash: []byte("test value"),
		}

		err := db.InsertOne(ctx, &user)
		if err != nil {
			t.Fatalf("InsertOne() fail: %v", err)
		}

		var dest = domain.User{ID: user.ID}
		db.First(&dest)

		if !dest.Equals(user) {
			t.Fatalf("InsertOne() fail: insert value not match search result.")
		}
	})

	t.Run("FindOne", func(t *testing.T) {
		result := db.SavePoint("FindOneBegin")
		if result.Error != nil {
			t.Fatalf("Fail in SavePoint: %v", err)
		}

		defer db.Rollbackto("FindOneBegin")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		// Insert value first for testing FindOne
		user := domain.User{
			Email:        "test@test.com",
			PasswordHash: []byte("test value"),
		}
		db.Create(&user)

		// Call FindOne and Compare result
		var dest domain.User
		err := db.FindOne(ctx, &dest, &domain.User{Email: "test@test.com"})
		if err != nil {
			t.Fatalf("FindOne() fail: %v", err)
		}

		if !dest.Equals(user) {
			t.Fatalf("Findone() fail: Search result does not match")
		}
	})

	t.Run("UpdateOne", func(t *testing.T) {
		result := db.SavePoint("UpdateOneBegin")
		if result.Error != nil {
			t.Fatalf("Fail in SavePoint: %v", err)
		}

		defer db.Rollbackto("UpdateOneBegin")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		user := domain.User{
			Email: "test@test.com",
			PasswordHash: []byte("test value"),
		}

		db.Create(&user)

		newUser := domain.User{
			Email: "newtest@test.com",
		}

		err := db.UpdateOne(ctx, &user, &newUser)
		if err != nil {
			t.Fatalf("UpdateOne() fail: %v", err)
		}

		var dest domain.User
		db.First(&dest, &domain.User{ID: user.ID})

		if !dest.Equals(user) {
			t.Fatalf("UpdateOne() fail: Update value do not match")
		}
	})

	t.Run("DeleteOne", func (t *testing.T){
		result := db.SavePoint("DeleteOneBegin")
		if result.Error != nil {
			t.Fatalf("Fail in SavePoint: %v", err)
		}

		defer db.Rollbackto("DeleteOneBegin")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		user := domain.User{
			Email: "test@test.com",
			PasswordHash: []byte("test value"),
		}

		db.Create(&user)

		err := db.DeleteOne(ctx, &user)
		if err != nil {
			t.Fatalf("DeleteOne() fail: %v", err)
		}

		result = db.First(&user)
		if (result.Error != gorm.ErrRecordNotFound) {
			if (result.Error == nil) {
				t.Fatalf("DeleteOne() fail: didn't delete record")
			} else {
				t.Fatalf("DelteOne() fail: error in db.First() %v", err)
			}
		}
	})

	t.Run("CountRows", func (t *testing.T){
		result := db.SavePoint("CountRowsBegin")
		if result.Error != nil {
			t.Fatalf("Fail in SavePoint: %v", err)
		}

		defer db.Rollbackto("CountRowsBegin")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		user1 := domain.User{
			Email: "test1@test.com",
			PasswordHash: []byte("test value"),
		}

		user2 := domain.User{
			Email: "test2@test.com",
			PasswordHash: []byte("test value"),
		}

		db.Create(&[]domain.User{user1, user2})

		num, err := db.CountRows(ctx, &domain.User{})
		if err != nil {
			t.Fatalf("CountRows() fail: %v", err)
		}

		if num != 2 {
			t.Fatalf("CountRows() fail: count number doesn't match")
		}
	})
}
