package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserCustomClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

type IJwtService interface {
	JwtSign(claim *UserCustomClaims) string
	JwtVerify(tokenStr string) (*UserCustomClaims, error)
}

type JwtService struct {
	secret []byte
}

func NewJwtService(secretKey string) *JwtService {
	return &JwtService{secret: []byte(secretKey)}
}

func NewUserCustomClaims(userId int, expireTime time.Duration) *UserCustomClaims {
	expirationTime := time.Now().Add(expireTime)
	return &UserCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
}

func (service *JwtService) JwtSign(claim *UserCustomClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	if tokenStr, err := token.SignedString(service.secret); err != nil {
		fmt.Printf("JWT sign err: %s", err)
		return ""
	} else {
		return tokenStr
	}
}

func (service *JwtService) JwtVerify(tokenStr string) (*UserCustomClaims, error) {
	claims := &UserCustomClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return service.secret, nil
	})

	if err != nil {
		fmt.Printf("JWT verify err: %s", err)
		return nil, err
	}

	return claims, nil
}
