package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoController struct{}

func NewTodoController() *TodoController {
	return new(TodoController)
}

type TodoIndexData struct{
	Message string `json:"message"`
}

func (c *TodoController) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &TodoIndexData{Message: "This is TodoIndexController"})
}