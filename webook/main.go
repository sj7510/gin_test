package main

import (
	"gin_test/webook/internal/repository"
	"gin_test/webook/internal/repository/dao"
	"gin_test/webook/internal/service"
	"gin_test/webook/internal/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	u := initUser(db)
	server := web.InitServer()
	u.RegisterUserRoutes(server)

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
