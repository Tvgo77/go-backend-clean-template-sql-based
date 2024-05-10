package router

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"go-backend/database"
	"go-backend/domain"
	"go-backend/setup"
)

func setupDB() (domain.Database, error) {
	passwd := setup.NewEnv().DBpassword
	dsn := "host=localhost user=postgres dbname=forumdb_test port=5432 sslmode=disable password=" + passwd
	db, err := database.NewDatabase(dsn)
	if err != nil {
		return nil, err
	}

	tx := db.Begin()
	err = tx.AutoMigrate(&domain.User{})
	return database.NewDatabaseFromExist(tx), err
}

func teardownDB(db domain.Database) {
	db.Rollback()
}

func TestSignup(t *testing.T) {
	// Setup test database
	env := setup.NewEnv()
	db, err := setupDB()
	assert.NoError(t, err)
	defer teardownDB(db)

	// Setup test gin engine
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	publicRouter := ginEngine.Group("")

	// Setup gin router
	SignupRouterSetup(env, db, publicRouter)

	// Setup test http request
	reqBody := domain.SignupRequest{
		Email:    "test@gmail.com",
		Password: "password",
	}
	jsonData, err := json.Marshal(reqBody)
	assert.NoError(t, err)
	

	t.Run("Success", func(t *testing.T) {
		resp := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(resp, ginEngine)
		req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		// Start test
		result := db.SavePoint("one")
		if result.Error != nil {
			t.Fatalf("Fail in SavePoint: %v", err)
		}
		defer db.Rollbackto("one")

		ginEngine.ServeHTTP(resp, c.Request)
		assert.Equal(t, http.StatusOK, resp.Result().StatusCode) //
		user := domain.User{}
		err = db.FindOne(context.Background(), &user, &domain.User{Email: reqBody.Email})
		assert.NoError(t, err)
		assert.Equal(t, reqBody.Email, user.Email) //
	})

	t.Run("Email already exists", func(t *testing.T) {
		resp := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(resp, ginEngine)
		req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		result := db.SavePoint("two")
		if result.Error != nil {
			t.Fatalf("Fail in SavePoint: %v", err)
		}
		defer db.Rollbackto("two")

		// Add user to database first
		user := domain.User{
			Email: reqBody.Email,
		}
		err = db.InsertOne(context.Background(), &user)
		assert.NoError(t, err)

		// Start test
		ginEngine.ServeHTTP(resp, c.Request)
		assert.Equal(t, http.StatusConflict, resp.Result().StatusCode) //

		db.Rollback()
	})

	t.Run("Bad Request", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer([]byte("bad request")))
		assert.NoError(t, err)
		resp := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(resp, ginEngine)
		c.Request = req

		// Start test
		result := db.SavePoint("three")
		if result.Error != nil {
			t.Fatalf("Fail in SavePoint: %v", err)
		}
		defer db.Rollbackto("three")

		ginEngine.ServeHTTP(resp, c.Request)
		assert.Equal(t, http.StatusBadRequest, resp.Result().StatusCode)
		db.Rollback()
	})
}
