package test_helper

import (
	"bytes"
	"encoding/json"
	"go-backend/database"
	"go-backend/domain"
	"go-backend/setup"
	"net/http"
)

type TestResponse struct {
	Message string `json:"message"`
}

func SetupDB() (domain.Database, error) {
	passwd := setup.NewEnv().DBpassword
	dsn := "host=localhost user=postgres dbname=postgres port=5432 sslmode=disable password=" + passwd
	db, err := database.NewDatabase(dsn)
	if err != nil {
		return nil, err
	}

	tx := db.Begin()
	err = tx.AutoMigrate(&domain.User{})
	return database.NewDatabaseFromExist(tx), err
}

func TeardownDB(db domain.Database) {
	db.Rollback()
}

func NewJSONreq(method string, url string, body any) (*http.Request, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	
	return req, err
}