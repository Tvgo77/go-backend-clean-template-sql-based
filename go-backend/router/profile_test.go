package router

import (
	"context"
	"encoding/json"
	"go-backend/domain"
	"go-backend/middleware"
	"go-backend/setup"
	testHelper "go-backend/test_helper"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)
func TestProfile(t *testing.T) {
	// Setup database
	db, err := testHelper.SetupDB()
	defer testHelper.TeardownDB(db)
	assert.NoError(t, err)

	// Setup gin engine and router
	env := setup.NewEnv()
	ginEngine := gin.Default()
	privateRouter := ginEngine.Group("")
	jm := middleware.NewJWTmiddleware(env)
	privateRouter.Use(jm.GinHandler)
	ProfileRouterSetup(env, db, privateRouter)

	// Test 1: Success fetch
	t.Run("Success fetch", func (t *testing.T) {
		// Setup request
		reqBody := domain.FetchProfileRequest{UID: 1}
		req, err := testHelper.NewJSONreq("GET", "/profile/1", &reqBody)
		assert.NoError(t, err)

		token, err := middleware.NewJWTuidToken(&domain.User{ID: 77}, env.TokenSecret)
		assert.NoError(t, err)

		req.Header.Set("Authorization", "Bearer " + token)
		resp := httptest.NewRecorder()

		// Create user in database first
		db.SavePoint("one")
		defer db.Rollbackto("one")
		user := testHelper.TestUser
		db.InsertOne(context.Background(), &user)

		// Start test
		ginEngine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Result().StatusCode)  // Expect http OK 200

		fetchResponse := domain.FetchProfileResponse{}
		json.Unmarshal(resp.Body.Bytes(), &fetchResponse)
		assert.Equal(t, user.Profile.Name, fetchResponse.Profile.Name)
	})

	// Test 2: User not exsit
	t.Run("User not exsit", func (t *testing.T) {
		// Setup request
		reqBody := domain.FetchProfileRequest{UID: 1}
		req, err := testHelper.NewJSONreq("GET", "/profile/1", &reqBody)
		assert.NoError(t, err)

		token, err := middleware.NewJWTuidToken(&domain.User{ID: 77}, env.TokenSecret)
		assert.NoError(t, err)

		req.Header.Set("Authorization", "Bearer " + token)
		resp := httptest.NewRecorder()

		// Start test
		ginEngine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusNotFound, resp.Result().StatusCode)  // Expect 404 Not Found
	})

	// Test 3: Invalid token
	t.Run("Invalid token", func (t *testing.T) {
		// Setup request
		reqBody := domain.FetchProfileRequest{UID: 1}
		req, err := testHelper.NewJSONreq("GET", "/profile/1", &reqBody)
		assert.NoError(t, err)

		req.Header.Set("Authorization", "Bearer " + "invalid_token")
		resp := httptest.NewRecorder()

		// Start test
		ginEngine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusUnauthorized, resp.Result().StatusCode)  // Expect 401 Unauthorized
	})

	// Test 4: Success update
	t.Run("Success update", func (t *testing.T) {
		// Setup request
		reqBody := domain.UpdateProfileRequest{UID: 1}
		req, err := testHelper.NewJSONreq("POST", "/profile/1", &reqBody)
		assert.NoError(t, err)

		token, err := middleware.NewJWTuidToken(&domain.User{ID: 1}, env.TokenSecret)
		assert.NoError(t, err)

		req.Header.Set("Authorization", "Bearer " + token)
		resp := httptest.NewRecorder()

		// Create user in database first
		db.SavePoint("four")
		defer db.Rollbackto("four")
		user := testHelper.TestUser
		db.InsertOne(context.Background(), &user)

		// Start test
		ginEngine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Result().StatusCode)  // Expect 200 OK
	})

	// Test 5: When update other's
	t.Run("When update other", func (t *testing.T) {
		// Setup request
		reqBody := domain.UpdateProfileRequest{UID: 1}
		req, err := testHelper.NewJSONreq("POST", "/profile/1", &reqBody)
		assert.NoError(t, err)

		token, err := middleware.NewJWTuidToken(&domain.User{ID: 77}, env.TokenSecret)
		assert.NoError(t, err)

		req.Header.Set("Authorization", "Bearer " + token)
		resp := httptest.NewRecorder()

		// Create user in database first
		db.SavePoint("four")
		defer db.Rollbackto("four")
		user := testHelper.TestUser
		db.InsertOne(context.Background(), &user)

		// Start test
		ginEngine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusUnauthorized, resp.Result().StatusCode)  // Expect 401 Unauthorized
	})
}