package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}
type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
type JwtService struct {
}

//NewJWTService method is creates a new instance of JWTService


func getSecretKey() string {
	return os.Getenv("JWT_SECRET_KEY")
}

func (j JwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    "go-jwt",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		panic(err)
	}
	return t
}

func (j JwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v ", t_.Header["alg"])
		}
		return []byte(getSecretKey()), nil
	})
}
