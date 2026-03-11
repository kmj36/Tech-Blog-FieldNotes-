package cryption

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret []byte
	expiry time.Duration
}

func NewJWTManager(secret []byte, expiry time.Duration) *JWTManager {
    return &JWTManager{secret: secret, expiry: expiry}
}

func (j *JWTManager) GenerateJWT(userID int, userAccount string, userNickName string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"user_account": userAccount,
		"user_nickname": userNickName,
		"exp": j.expiry,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JWTManager) ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return j.secret, nil
	})

	return token, err
}