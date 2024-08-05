package main

import (
    "todo-api/handlers"
    _ "todo-api/docs"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    echoSwagger "github.com/swaggo/echo-swagger"
)

// @title ToDo API
// @version 1.0
// @description This is a sample server for managing ToDo items.
// @host localhost:8080
// @BasePath /
func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    e.GET("/todos", handlers.GetToDos)
    e.POST("/todos", handlers.CreateToDo)
    e.PUT("/todos/:id", handlers.UpdateToDo)
    e.DELETE("/todos/:id", handlers.DeleteToDo)

    // Swagger
    e.GET("/swagger/*", echoSwagger.WrapHandler)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
