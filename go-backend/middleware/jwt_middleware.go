package middleware

import (
	"log"
	"net/http"
	"strings"

	"go-backend/domain"
	"go-backend/setup"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTmiddleware struct {
	secret []byte
}

func NewJWTmiddleware(env *setup.Env) *JWTmiddleware {
	return &JWTmiddleware{secret: []byte(env.TokenSecret)}
}

// For JWT token authentication
func (jm *JWTmiddleware) GinHandler(c *gin.Context) {
	// Before handle request
	// Extract token from request header
	// Example Authorization fields:
	// Authorization: bear <token>
	credential := c.Request.Header.Get("Authorization")
	token := strings.Split(credential, " ")[1] // []string{"bear", "<token>"}

	// Verify token
	// A simplest verifyFunc just need to return the secret used in signature
	parsedToken, err := jwt.ParseWithClaims(
		token, 
		&jwt.RegisteredClaims{}, 
		func (token *jwt.Token) (interface{}, error) {
			return jm.secret, nil
		}, 
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid token: " + err.Error()})
		c.Abort()
		return
	}

	if claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims); ok {
		// Set token's user id in gin context
		userID, err := claims.GetSubject()
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", userID)
	} else {
		log.Fatal("Unknow claims type")
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid token: Unknow claims type"})
		c.Abort()
		return
	}

	c.Next()

	// After send response
}

