package handler // handlerパッケージであることを宣言

import "github.com/labstack/echo/v4" // echoを使用するためのパッケージ

type TodoHandler interface { // TodoHandlerインターフェースを定義
	GettTodo(c echo.Context) error // GetTodoメソッドを定義
}
