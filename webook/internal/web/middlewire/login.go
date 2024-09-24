package middlewire

import (
	"encoding/gob"
	"gin_test/webook/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		gob.Register(time.Now())
		// Don't need validate
		if ctx.Request.URL.Path == "/users/login" ||
			ctx.Request.URL.Path == "/users/signup" {
			return
		}

		authCode := ctx.GetHeader("Authorization")

		if authCode == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		claims := &web.UserClaims{}
		token, err := jwt.ParseWithClaims(authCode, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil || !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("claims", claims)
	}
}
