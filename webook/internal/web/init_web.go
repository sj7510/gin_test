package web

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func RegisterRouters() *gin.Engine {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
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

	u := NewUserHandler()
	u.registerUserRoutes(server)

	return server

}
