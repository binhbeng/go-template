package middleware

import (
	"net/http"
	"github.com/binhbeng/goex/internal/api"
	"github.com/binhbeng/goex/internal/global"
	"github.com/binhbeng/goex/internal/pkg/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		accessToken, err := token.GetAccessToken(authorization)
		if err != nil {
			api.HandleError(c, http.StatusUnauthorized, "" , err)
			return
		}
		customClaims := new(token.CustomClaims)

		err = token.Parse(accessToken, customClaims, jwt.WithSubject(global.Subject))
		if err != nil {
			api.HandleError(c, http.StatusUnauthorized, "" , err)
			return
		}

		exp, err := customClaims.GetExpirationTime()
		if err != nil || exp == nil {
			api.HandleError(c, http.StatusUnauthorized, "" , err)
			return
		}

		c.Set("user_id", customClaims.UserID)
		c.Set("username", customClaims.Username)
		c.Next()
	}
}
