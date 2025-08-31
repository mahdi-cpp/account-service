package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing
var secretKey = []byte("$Mahdi@_123456")

// CustomClaims Define your custom claims (optional)
type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func main() {
	createAccessToken()
}

func createAccessToken() {

	// Create claims with expiration time
	claims := CustomClaims{
		UserID: "018f3a8b-1b32-729a-f7e5-5467c1b2d3e4",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Second)),
			Issuer:    "com.iris.all",
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Token:", signedToken)
}

func createRefreshToken() {

	// Create claims with expiration time
	claims := CustomClaims{
		UserID: "09355512620",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			Issuer:    "com.iris.all",
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Token:", signedToken)
}

func VerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
