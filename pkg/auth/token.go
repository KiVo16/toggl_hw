package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type JWTManager struct {
	secret []byte
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWTManager(secret string) JWTManager {
	return JWTManager{
		secret: []byte(secret),
	}
}

func (m JWTManager) ValidateTokenAndExtractData(tokenData string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenData, claims, func(token *jwt.Token) (interface{}, error) {
		return m.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token invalid")
	}

	return claims, nil
}

func (JWTManager) ExtractJWTFromHeaderHTTP(headers http.Header) (*string, error) {
	reqToken := headers.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) < 2 {
		return nil, fmt.Errorf("JWT token missing")
	}

	return &splitToken[1], nil
}
