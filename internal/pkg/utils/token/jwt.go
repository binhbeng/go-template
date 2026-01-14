package token

import (
	"errors"
	"strings"
	"time"
	"github.com/binhbeng/goex/config"
	"github.com/binhbeng/goex/internal/global"
	"github.com/binhbeng/goex/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

type JwtPayload struct {
	UserID   int64   `json:"user_id"`
	Username string `json:"username"`
}

func GetJwtPayload(info any) (jwtPayload JwtPayload) {
	jwtPayload, _ = info.(JwtPayload)
	return
}

func Generate(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(config.Cfg.Jwt.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func Refresh(claims jwt.Claims) (string, error) {
	return Generate(claims)
}

func Parse(accessToken string, claims jwt.Claims, options ...jwt.ParserOption) error {
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (i any, err error) {
		return []byte(config.Cfg.Jwt.SecretKey), err
	}, options...)
	if err != nil {
		return err
	}

	if token.Valid { 
		return nil
	}

	return errors.New("invalid token")
}

func GetAccessToken(authorization string) (accessToken string, err error) {
	if authorization == "" {
		return "", errors.New("authorization header is missing")
	}

	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", errors.New("invalid Authorization header format")
	}

	accessToken = strings.TrimPrefix(authorization, "Bearer ")
	return
}

type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims 
}

func NewCustomClaims(user *model.User, expiresAt time.Time) CustomClaims {
	return CustomClaims{
		JwtPayload: JwtPayload{
			user.ID,
			user.Username,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt), 
			Issuer:    global.Issuer,              
			Subject: global.Subject, 
		},
	}
}
