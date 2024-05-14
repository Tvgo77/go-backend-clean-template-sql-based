package middleware

import (
	"bytes"
	"encoding/json"
	"go-backend/domain"
	"go-backend/setup"
	"go-backend/test_helper"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestJWTmiddleware(t *testing.T) {
	// Setup gin engine
	env := setup.NewEnv()
	ginEngine := gin.Default()

	// Setup jwt middleware
	jm := NewJWTmiddleware(env)
	ginEngine.Use(jm.GinHandler)

	// Setup test gin handler
	ginEngine.Handle("GET", "/testJWT", func (c *gin.Context) {
		// Already Passed middleware, response directly
		uid := c.GetString("userID")
		c.JSON(http.StatusOK, test_helper.TestResponse{Message: uid})
	})

	// Test 1: valid token
	t.Run("Valid token", func (t *testing.T) {
		// Setup Request
		token, err := NewJWTuidToken(&domain.User{ID: 123}, env.TokenSecret)
		assert.NoError(t, err)

		req, err := http.NewRequest("GET", "/testJWT", bytes.NewBuffer([]byte{}))
		assert.NoError(t, err)
		req.Header.Set("Authorization", "Bearer " + token)

		// Start test
		resp := httptest.NewRecorder()
		ginEngine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Result().StatusCode)

		var respBody test_helper.TestResponse
		json.Unmarshal(resp.Body.Bytes(), &respBody)
		assert.Equal(t, "123", respBody.Message)
	})
	
	// Test 2: invalid token
	t.Run("Invalid token", func (t *testing.T) {
		// Setup Request
		token := "invalid"

		req, err := http.NewRequest("GET", "/testJWT", bytes.NewBuffer([]byte{}))
		assert.NoError(t, err)
		req.Header.Set("Authorization", "Bearer " + token)

		// Start test
		resp := httptest.NewRecorder()
		ginEngine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusUnauthorized, resp.Result().StatusCode)
	})
}