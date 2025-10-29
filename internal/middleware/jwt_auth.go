package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/binhbeng/goex/internal/global"
	e "github.com/binhbeng/goex/internal/pkg/errors"
	"github.com/binhbeng/goex/internal/pkg/response"
	"github.com/binhbeng/goex/internal/pkg/utils/token"
)

func JwtAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		accessToken, err := token.GetAccessToken(authorization)
		if err != nil {
			response.Fail(c, e.Unauthorized, err.Error())
			return
		}
		customClaims := new(token.CustomClaims)

		err = token.Parse(accessToken, customClaims, jwt.WithSubject(global.Subject))
		if err != nil {
			response.FailCode(c, e.Unauthorized)
			return
		}

		exp, err := customClaims.GetExpirationTime()
		if err != nil || exp == nil {
			response.FailCode(c, e.Unauthorized)
			return
		}

		c.Set("user_id", customClaims.UserID)
		c.Set("username", customClaims.Username)
		c.Next()
	}
}
