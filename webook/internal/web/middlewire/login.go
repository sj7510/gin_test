package middlewire

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
		// Don't need validate
		if ctx.Request.URL.Path == "/users/login" ||
			ctx.Request.URL.Path == "/users/signup" {
			return
		}
		sess := sessions.Default(ctx)
		if sess.Get("userId") == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		updateTime := sess.Get("update_time")
		now := time.Now().UnixMilli()

		// first login
		if updateTime == nil {
			sess.Set("update_time", now)
			_ = sess.Save()
			return
		}

		updateTimeVal, _ := updateTime.(int64)

		// 1 minutes
		if now-updateTimeVal > 60*1000 {
			// refresh session
			sess.Set("update_time", now)
			_ = sess.Save()
			return
		}
	}
}
