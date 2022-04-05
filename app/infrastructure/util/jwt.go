package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	JWTSecret = "a-very-secret-string"
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
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

func (j *JWT) ParseToken(tokenString string) (*CustomClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaim); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}
