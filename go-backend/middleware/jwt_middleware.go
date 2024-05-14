package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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

func NewJWTuidToken(user *domain.User, secret string) (string, error) {
	myClaims := jwt.RegisteredClaims{
		Issuer: "fantasyforum",
		Subject: fmt.Sprintf("%d", user.ID),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	signedToken, err := token.SignedString([]byte(secret))
	return signedToken, err
}

// Return parsed token if verification success
func VerifyToken(token string, secret []byte) (*jwt.Token, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token, 
		&jwt.RegisteredClaims{}, 
		func (token *jwt.Token) (interface{}, error) {
			return secret, nil
		}, 
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	return parsedToken, err
}

// For JWT token authentication
func (jm *JWTmiddleware) GinHandler(c *gin.Context) {
	// Before handle request
	// Extract token from request header
	// Example Authorization fields:
	// Authorization: bear <token>
	credential := c.Request.Header.Get("Authorization")
	authFields := strings.Split(credential, " ") // []string{"bear", "<token>"}
	if len(authFields) < 2 {
		log.Fatal("Invalid Authorization field in header")
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid Authorization field in header"})
		c.Abort()
		return
	}
	token := authFields[1]  // []string{"bear", "<token>"}

	// Verify token
	// A simplest verifyFunc just need to return the secret used in signature
	parsedToken, err := VerifyToken(token, jm.secret)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid token: " + err.Error()})
		c.Abort()
		return
	}

	if claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims); ok {
		// Set token's user id in gin context
		userID, err := claims.GetSubject()
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", userID)
	} else {
		log.Print("Unknow claims type")
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid token: Unknow claims type"})
		c.Abort()
		return
	}

	c.Next()

	// After send response
}

