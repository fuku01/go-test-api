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

	// TagテーブルとTodoテーブルを結合して、全てのTodoに加えて、そのTodoが持つ全てのTagを取得するメソッドを定義
	GetAllWithTags(c echo.Context) error // echo.Contextとは、HTTPリクエストとレスポンスを扱うための構造体。

	// *CreateWithTagsメソッド（トランザクションを使用して、TodoとTagを同時に作成）
	CreateWithTags(c echo.Context) error // echo.Contextとは、HTTPリクエストとレスポンスを扱うための構造体。

	// *DeleteWithTagsメソッド（トランザクションを使用して、TodoとTagを同時に削除）
	DeleteWithTags(c echo.Context) error // echo.Contextとは、HTTPリクエストとレスポンスを扱うための構造体。

	// !EditWithTagsメソッド(トランザクションを使用して、Todoとそれに紐付くTagの追加と削除を行う)
	EditWithTags(c echo.Context) error // echo.Contextとは、HTTPリクエストとレスポンスを扱うための構造体。
}

// @ 構造体の型。
type todoHandler struct {
	tu usecase.TodoUsecase
	uu usecase.UserUsecase
}

// IDとContentだけを含む新しい構造体（練習用）
type todoContent struct {
	ID      uint   `json:"ID"`
	Content string `json:"content"`
}

// *TodoとそのTodoが持つ全てのTagを含む新しい構造体
type todoWithTags struct {
	Todo     *model.Todo `json:"todo"`
	TagNames []string    `json:"tagNames"`
}

// ! Edit用のリクエスト構造体
type editTodoRequest struct {
	Content      string   `json:"content"`      // 新たに設定したいTODOの内容
	AddTagNames  []string `json:"addTagNames"`  // 追加したいタグの名前のリスト
	DeleteTagIDs []uint   `json:"deleteTagIDs"` // 削除したいタグのIDのリスト
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

	// IDとContentだけを含む新しい構造体を作成。（練習用）
	content := []todoContent{}   // ? TodoContent構造体の[配列]を作成。
	for _, todo := range todos { // ? todosの中身を順番にtodoに代入。_は、index番号であり、今回は使用しないため_としている。
		content = append(content, todoContent{ // ? appendとは、[配列]に要素を追加するメソッド。（追加先の[配列]と追加する要素。）
			ID:      todo.ID,
			Content: todo.Content,
		})
	}

	return c.JSON(http.StatusOK, content) // 200ステータスコードとcontentを返す
}

// GetAllWithTagsメソッドを定義
func (h todoHandler) GetAllWithTags(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization") // リクエストヘッダーからAuthorizationを取得
	token := strings.TrimPrefix(authHeader, "Bearer ")    // Bearerを削除

	todosWithTags, err := h.tu.GetAllWithTags(token) // 全てのtodoとそのtodoが持つ全てのtagを取得
	if err != nil {                                  // エラーがあれば
		return err // エラーを返す
	}

	return c.JSON(http.StatusOK, todosWithTags) // 200ステータスコードとtodosWithTagsを返す
}

// *CreateWithTagsメソッドを定義
func (h todoHandler) CreateWithTags(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization") // リクエストヘッダーからAuthorizationを取得
	token := strings.TrimPrefix(authHeader, "Bearer ")    // Bearerを削除

	todoWithTags := &todoWithTags{} // TodoとそのTodoが持つ全てのTagを含む新しい構造体を作成
	if err := c.Bind(todoWithTags); err != nil {
		return err
	}

	newTodo, err := h.tu.CreateWithTags(todoWithTags.Todo.Content, token, todoWithTags.TagNames)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newTodo)
}

// Createメソッドを定義
func (h todoHandler) Create(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	todo := &model.Todo{}                // 「Todo」型のポインタを生成
	if err := c.Bind(todo); err != nil { // フロントから受け取ったJSONをtodoにバインド
		return err // エラーがあればエラーを返す
	}

	newTodo, err := h.tu.Create(todo.Content, token) // フロントから受け取ったcontentをtodoに代入
	if err != nil {                                  // エラーがあれば
		return err // エラーを返す
	}

	return c.JSON(http.StatusOK, newTodo) // 200ステータスコードとcreatedTodoを返す
}

// * DeleteWithTagsメソッドを定義
func (h todoHandler) DeleteWithTags(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	idParam := c.Param("ID") // URLからTODOのidパラメータを取得

	ID, err := strconv.Atoi(idParam) // idパラメータをintに変換
	if err != nil {                  // 変換エラーがあれば
		return err // エラーを返す
	}

	if err := h.tu.DeleteWithTags(uint(ID), token); err != nil { // 変換したidを用いて削除。idをuint（符号なし整数）に変換。
		return err // エラーを返す
	}

	return c.NoContent(http.StatusNoContent) // 204ステータスコードを返す
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

// ! EditWithTagsメソッドを定義(トランザクションを使用して、Todoとそれに紐付くTagの追加と削除を行う)
func (h todoHandler) EditWithTags(c echo.Context) error {
	// リクエストヘッダーからAuthorizationを取得
	authHeader := c.Request().Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	idParam := c.Param("ID") // URLからTODOのidパラメータを取得

	// idParamはstring型なので、これをint型に変換する必要があります。
	ID, err := strconv.Atoi(idParam) // idパラメータをintに変換
	if err != nil {                  // 変換エラーがあれば
		return err // エラーを返す
	}

	request := &editTodoRequest{}           // Edit用のリクエスト構造体を作成
	if err := c.Bind(request); err != nil { // フロントから受け取ったJSONをrequestにバインド
		return err
	}

	editedTodo, err := h.tu.EditWithTags(uint(ID), token, request.Content, request.AddTagNames, request.DeleteTagIDs)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, editedTodo) // 200ステータスコードとeditedTodoを返す
}
