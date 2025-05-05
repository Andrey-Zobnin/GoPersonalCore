
package main

import (
    "persona-core/db"
    "persona-core/handler"
    "persona-core/repository"
    "persona-core/service"

    "github.com/gin-gonic/gin"
)

func main() {
    database := db.Init()
    userRepo := repository.NewUserRepository(database)
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    r := gin.Default()
    r.POST("/users", userHandler.CreateUser)
    r.GET("/users/:id", userHandler.GetUser)
    r.Run(":8080")
}
