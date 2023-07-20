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

// @ Todoに関する、handlerメソッドの集まり（インターフェース）を定義。
type TodoHandler interface {
	GetAll(c echo.Context) error // 全てのTodoを取得するメソッドを定義
	Create(c echo.Context) error // 新しいTodoを作成するメソッドを定義
	Delete(c echo.Context) error // 指定したTodoを削除するメソッドを定義
}

// @ 構造体の型。
type todoHandler struct {
	todoUsecase usecase.TodoUsecase
	userUsecase usecase.UserUsecase
}

// @ /mainのルーティングで、この構造体を使用する（呼び出す）ための関数を定義。
func NewTodoHandler(todoUsecase usecase.TodoUsecase, userUsecase usecase.UserUsecase) TodoHandler {
	return &todoHandler{todoUsecase: todoUsecase, userUsecase: userUsecase}
}

// @ フロントからのHTTPリクエストを受け取り、/usecase層で実装した【具体的な処理】を呼び出し、フロントへ返すレスポンスを生成。

// GetAllメソッドを定義
func (h todoHandler) GetAll(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	user, err := h.userUsecase.GetUserByToken(context.Background(), token)
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

	user, err := h.userUsecase.GetUserByToken(context.Background(), token)
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

	user, err := h.userUsecase.GetUserByToken(context.Background(), token)
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
