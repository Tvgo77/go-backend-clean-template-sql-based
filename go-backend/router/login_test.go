package router

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-backend/domain"
	"go-backend/middleware"
	"go-backend/setup"
	testHelper "go-backend/test_helper"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// Setup Database
	db, err := testHelper.SetupDB()
	defer testHelper.TeardownDB(db)
	assert.NoError(t, err)

	// Setup gin engine and router
	env := setup.NewEnv()
	ginEngine := gin.Default()
	publicRouter := ginEngine.Group("")
	LoginRouterSetup(env, db, publicRouter)

	// Test 1: Successful login
	t.Run("Success login", func(t *testing.T) {
		// Setup test request
		reqBody := testHelper.LoginReqBody
		req, err := testHelper.NewJSONreq("POST", "/login", &reqBody)
		assert.NoError(t, err)

		// Setup gin context
		resp := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(resp, ginEngine)
		c.Request = req

		// Create user in database first
		db.SavePoint("one")
		defer db.Rollbackto("one")
		user := testHelper.TestUser
		db.InsertOne(context.Background(), &user)

		// Start test
		ginEngine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Result().StatusCode)  // Expect http OK 200
		
		loginResp := domain.LoginResponse{}
		json.Unmarshal(resp.Body.Bytes(), &loginResp)

		// Verify token
		token := loginResp.Token
		parsedToken, err := middleware.VerifyToken(token, []byte(env.TokenSecret))
		assert.NoError(t, err)
		subject, err := parsedToken.Claims.GetSubject()
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%d", user.ID), subject)  // Expect user id is same with token's

	})

	// Test 2: Wrong password
	t.Run("Wrong password", func(t *testing.T) {
		// Setup test request
		reqBody := testHelper.LoginReqBody
		reqBody.Password = "wrong password"
		req, err := testHelper.NewJSONreq("POST", "/login", &reqBody)
		assert.NoError(t, err)

		// Setup gin context
		resp := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(resp, ginEngine)
		c.Request = req

		// Create user in database first
		db.SavePoint("two")
		defer db.Rollbackto("two")
		user := testHelper.TestUser
		db.InsertOne(context.Background(), &user)

		// Start test
		ginEngine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusUnauthorized, resp.Result().StatusCode)  // Expect http Unauthorized 401
	})

	// Test 3: User not exist
	t.Run("User not exist", func(t *testing.T) {
		// Setup test request
		reqBody := testHelper.LoginReqBody
		req, err := testHelper.NewJSONreq("POST", "/login", &reqBody)
		assert.NoError(t, err)

		// Setup gin context
		resp := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(resp, ginEngine)
		c.Request = req

		// Start test
		ginEngine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusUnauthorized, resp.Result().StatusCode)  // Expect http Unauthorized 401
	})

	// Test 4: Bad Request
	t.Run("Bad Request", func(t *testing.T) {
		// Setup test request
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte("bad request")))
		assert.NoError(t, err)

		// Setup gin context
		resp := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(resp, ginEngine)
		c.Request = req

		// Start test
		ginEngine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Result().StatusCode)  // Expect http bad request 400
	})
}