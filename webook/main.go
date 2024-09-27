package main

import (
	"gin_test/webook/internal/repository"
	"gin_test/webook/internal/repository/dao"
	"gin_test/webook/internal/service"
	"gin_test/webook/internal/web"
	"gin_test/webook/internal/web/middlewire"
	"gin_test/webook/pkg/ginx/middleware/ratelimit"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	// db := initDB()
	// u := initUser(db)
	// server := initServer()
	// u.RegisterUserRoutes(server)

	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})
	_ = server.Run(":8080")
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initServer() *gin.Engine {
	server := gin.Default()

	// Implementing current limiting through Redis
	cmd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	// 1 minute, 100 requests
	server.Use(ratelimit.NewBuilder(cmd, 1*time.Minute, 100).Build())

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
