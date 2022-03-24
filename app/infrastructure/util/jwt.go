package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	JWTSecret = "a-very-secret-string"
)

type TokenInfo struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type CustomClaim struct {
	ID       uint
	Username string
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(JWTSecret),
	}
}

func (j *JWT) GenerateToken(ID uint, username string) (TokenInfo, error) {
	now := time.Now()
	expiredAt := now.Add(time.Hour * 24)
	customClaim := &CustomClaim{
		ID,
		username,
		jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
			Issuer:    "tuantt",
		},
	}
	token, err := j.CreateToken(*customClaim)
	return TokenInfo{
		Token:     token,
		ExpiredAt: expiredAt,
	}, err
}

func (j *JWT) CreateToken(customClaim CustomClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim)
	return token.SignedString(j.SigningKey)
}
