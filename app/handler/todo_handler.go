package handler

import (
	"net/http"

	"github.com/fuku01/go-test-api/app/usecase"
	"github.com/labstack/echo/v4"
)

type TodoHandler interface {
	GetAll(c echo.Context) error
}

type todoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) todoHandler {
	return todoHandler{todoUsecase: todoUsecase}
}
func (h todoHandler) GetAll(c echo.Context) error {
	todos, err := h.todoUsecase.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todos)
}
