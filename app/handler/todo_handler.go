package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/usecase"
	"github.com/labstack/echo/v4"
)

// 1. TodoHandlerインターフェースを定義
// 2. TodoHandlerインターフェースを実装する構造体は、この3つのメソッドを実装しなければならない

type TodoHandler interface {
	GetAll(c echo.Context) error // GetAllメソッドを定義
	Create(c echo.Context) error // Createメソッドを定義
	Delete(c echo.Context) error // Deleteメソッドを定義
}

type todoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) TodoHandler {
	return todoHandler{todoUsecase: todoUsecase}
}

// GetAllメソッドを定義
func (h todoHandler) GetAll(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	fmt.Println("トークン：", token)

	user, err := h.todoUsecase.GetUserByToken(context.Background(), token)
	if err != nil {
		fmt.Println("エラー：", err)
		return err
	}

	todos, err := h.todoUsecase.GetAll(user.ID) // 全てのtodoを取得
	if err != nil {                             // エラーがあれば
		fmt.Println("エラー：", err)
		return err // エラーを返す
	}
	return c.JSON(http.StatusOK, todos) // 200ステータスコードとtodosを返す
}

// Createメソッドを定義
func (h todoHandler) Create(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	fmt.Println("トークン：", token)

	user, err := h.todoUsecase.GetUserByToken(context.Background(), token)
	if err != nil {
		return err
	}

	todo := &model.Todo{}                // 「Todo」型のポインタを生成
	if err := c.Bind(todo); err != nil { // フロントから受け取ったJSONをtodoにバインド
		return err // エラーがあればエラーを返す
	}

	createdTodo, err := h.todoUsecase.Create(todo.Content, user.ID) // フロントから受け取ったcontentをtodoに代入
	if err != nil {                                                 // エラーがあれば
		return err // エラーを返す
	}

	return c.JSON(http.StatusOK, createdTodo) // 200ステータスコードとcreatedTodoを返す
}

// Dleteメソッドを定義
func (h todoHandler) Delete(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	fmt.Println("トークン：", token)

	user, err := h.todoUsecase.GetUserByToken(context.Background(), token)
	if err != nil {
		return err
	}

	idParam := c.Param("ID") // URLからTODOのidパラメータを取得

	ID, err := strconv.Atoi(idParam) // idパラメータをintに変換
	if err != nil {                  // 変換エラーがあれば
		return err // エラーを返す
	}

	if err := h.todoUsecase.Delete(uint(ID), user.ID); err != nil { // 変換したidを用いて削除。idをuint（符号なし整数）に変換。
		return err // エラーを返す
	}
	return c.NoContent(http.StatusNoContent) // 204ステータスコードを返す
}
