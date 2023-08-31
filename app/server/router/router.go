package router

import (
	"app/controllers"

	"github.com/labstack/echo/v4"
)

func DefineRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")

	todoController := controllers.NewTodoController()

	apiGroup.GET("/todo", todoController.Index)
}