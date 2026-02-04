package middleware

import (
	"github.com/binhbeng/goex/internal/global"
	"github.com/binhbeng/goex/internal/utils"
	"github.com/binhbeng/goex/internal/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		accessToken, err := token.GetAccessToken(authorization)
		if err != nil {
			utils.HttpUnauthorized(c, "", err)
			return
		}
		customClaims := new(token.CustomClaims)

		err = token.Parse(accessToken, customClaims, jwt.WithIssuer(global.Issuer))
		if err != nil {
			utils.HttpUnauthorized(c, "", err)
			return
		}

		exp, err := customClaims.GetExpirationTime()
		if err != nil || exp == nil {
			utils.HttpUnauthorized(c, "", err)
			return
		}

		c.Set("user_id", customClaims.UserID)
		c.Set("username", customClaims.Username)
		c.Next()
	}
}
