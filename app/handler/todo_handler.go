package handler

import (
	"net/http"
	"strconv"

	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/usecase"
	"github.com/labstack/echo/v4"
)

type TodoHandler interface {
	GetAll(c echo.Context) error
	Create(c echo.Context) error
	Delete(c echo.Context) error
}

type todoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) todoHandler {
	return todoHandler{todoUsecase: todoUsecase}
}

// GetAllメソッドを定義
func (h todoHandler) GetAll(c echo.Context) error {
	todos, err := h.todoUsecase.GetAll() // 全てのtodoを取得
	if err != nil {                      // エラーがあれば
		return err // エラーを返す
	}
	return c.JSON(http.StatusOK, todos) // 200ステータスコードとtodosを返す
}

// Createメソッドを定義
func (h todoHandler) Create(c echo.Context) error {
	todo := &model.Todo{}                // 「Todo」型のポインタを生成
	if err := c.Bind(todo); err != nil { // フロントから受け取ったJSONをtodoにバインド
		return err // エラーがあればエラーを返す
	}

	createdTodo, err := h.todoUsecase.Create(todo.Content) // フロントから受け取ったcontentをtodoに代入
	if err != nil {                                        // エラーがあれば
		return err // エラーを返す
	}

	return c.JSON(http.StatusOK, createdTodo) // 200ステータスコードとcreatedTodoを返す
}

// Dleteメソッドを定義
func (h todoHandler) Delete(c echo.Context) error {
	idParam := c.Param("ID") // URLからidパラメータを取得

	ID, err := strconv.Atoi(idParam) // idパラメータをintに変換
	if err != nil {                  // 変換エラーがあれば
		return err // エラーを返す
	}

	if err := h.todoUsecase.Delete(uint(ID)); err != nil { // 変換したidを用いて削除。idをuint（符号なし整数）に変換。
		return err // エラーを返す
	}
	return c.NoContent(http.StatusNoContent) // 204ステータスコードを返す
}
