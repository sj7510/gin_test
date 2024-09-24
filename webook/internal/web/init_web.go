package web

import (
	"gin_test/webook/internal/web/middlewire"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func InitServer() *gin.Engine {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"x-jwt-token"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				// dev mode
				return true
			}
			// your company domain
			return strings.Contains(origin, "https://xxx.xxx.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	server.Use(middlewire.NewLoginMiddlewareBuilder().Build())

	return server

}
