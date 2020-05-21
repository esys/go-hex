package main

import (
	"go-hex/internal/user/domain"
	"go-hex/internal/user/rest"
	"go-hex/internal/user/storage/sqlite"

	"github.com/gin-gonic/gin"
)

func main() {
	db := sqlite.OpenDatabase(":memory:")
	userRepository := sqlite.NewSQLiteUserRepository(db)
	userService := domain.NewUserService(userRepository)
	userHandler := rest.NewUserHandler(userService)

	r := gin.Default()
	r.GET("/user", userHandler.Get)
	r.GET("/user/:id", userHandler.GetById)
	r.POST("/user", userHandler.Create)
	r.Run()
}
