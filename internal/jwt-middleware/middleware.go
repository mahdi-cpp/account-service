package jwt_middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing
var secretKey = []byte("$Mahdi@_123456")

// CustomClaims Define your custom claims (optional)
type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateAccessToken() {

	// Create claims with expiration time
	claims := CustomClaims{
		UserID: "09125640200",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Second)),
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

func CreateRefreshToken() {

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

func VerifyAuthorizationToken(c *gin.Context) (*CustomClaims, error) {

	// Extract the token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is required")
	}

	//// Check if the header uses the Bearer scheme
	//if len(authHeader) < 7 || strings.ToLower(authHeader[:7]) != "bearer " {
	//	return nil, fmt.Errorf("authorization header format must be 'Bearer {token}'")
	//}
	//
	//// Extract the token part (remove "Bearer " prefix)
	//tokenString := authHeader[7:]

	tokenString := authHeader

	// Parse and validate the token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// Verify token claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// Additional validation (e.g., check expiration, issuer, etc.)
		if time.Now().After(claims.ExpiresAt.Time) {
			return nil, fmt.Errorf("token has expired")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")

}
