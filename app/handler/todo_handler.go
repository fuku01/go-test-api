package handler

import (
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
	// 全てのTodoを取得するメソッドを定義
	GetAll(c echo.Context) error //　echo.Contextとは、HTTPリクエストとレスポンスを扱うための構造体。

	// 新しいTodoを作成するメソッドを定義
	Create(c echo.Context) error // echo.Contextとは、HTTPリクエストとレスポンスを扱うための構造体。

	// 指定したTodoを削除するメソッドを定義
	Delete(c echo.Context) error // echo.Contextとは、HTTPリクエストとレスポンスを扱うための構造体。
}

// @ 構造体の型。
type todoHandler struct {
	tu usecase.TodoUsecase
	uu usecase.UserUsecase
}

// *IDとContentだけを含む新しい構造体（練習用）
type todoContent struct {
	ID      uint   `json:"ID"`
	Content string `json:"content"`
}

// @ /mainのルーティングで、この構造体を使用する（呼び出す）ための関数を定義。
// ? tu2,uu2に引数で各々のインターフェースを満たすオブジェクトを受け取り、TodoHandlerのインターフェースを満たすような新しいtodoHandler構造体を作成して返す。
func NewTodoHandler(tu2 usecase.TodoUsecase, uu2 usecase.UserUsecase) TodoHandler {
	return &todoHandler{tu: tu2, uu: uu2}
}

// @ フロントからのHTTPリクエストを受け取り、/usecase層で実装した【具体的な処理】を呼び出し、フロントへ返すレスポンスを生成。

// GetAllメソッドを定義
func (h todoHandler) GetAll(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization") // リクエストヘッダーからAuthorizationを取得
	token := strings.TrimPrefix(authHeader, "Bearer ")    // Bearerを削除

	todos, err := h.tu.GetAll(token) // 全てのtodoを取得
	if err != nil {                  // エラーがあれば
		fmt.Println("エラー：", err)
		return err // エラーを返す
	}

	// *IDとContentだけを含む新しい構造体を作成。（練習用）
	content := []todoContent{}   // ? TodoContent構造体の[配列]を作成。
	for _, todo := range todos { // ? todosの中身を順番にtodoに代入。_は、index番号であり、今回は使用しないため_としている。
		content = append(content, todoContent{ // ? appendとは、[配列]に要素を追加するメソッド。（追加先の[配列]と追加する要素。）
			ID:      todo.ID,
			Content: todo.Content,
		})
	}

	return c.JSON(http.StatusOK, content) // 200ステータスコードとcontentを返す
}

// Createメソッドを定義
func (h todoHandler) Create(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	todo := &model.Todo{}                // 「Todo」型のポインタを生成
	if err := c.Bind(todo); err != nil { // フロントから受け取ったJSONをtodoにバインド
		return err // エラーがあればエラーを返す
	}

	createdTodo, err := h.tu.Create(todo.Content, token) // フロントから受け取ったcontentをtodoに代入
	if err != nil {                                      // エラーがあれば
		return err // エラーを返す
	}

	return c.JSON(http.StatusOK, createdTodo) // 200ステータスコードとcreatedTodoを返す
}

// Dleteメソッドを定義
func (h todoHandler) Delete(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	idParam := c.Param("ID") // URLからTODOのidパラメータを取得

	ID, err := strconv.Atoi(idParam) // idパラメータをintに変換
	if err != nil {                  // 変換エラーがあれば
		return err // エラーを返す
	}

	if err := h.tu.Delete(uint(ID), token); err != nil { // 変換したidを用いて削除。idをuint（符号なし整数）に変換。
		return err // エラーを返す
	}

	return c.NoContent(http.StatusNoContent) // 204ステータスコードを返す
}
