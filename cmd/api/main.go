package main // mainパッケージであることを宣言

import (
	"github.com/labstack/echo/v4" // echoを使用するためのパッケージ
	"github.com/labstack/echo/v4/middleware"
)

// サーバーを起動する関数
func main() {
	e := echo.New()                              // Rest APIを使用するためのインスタンスを作成
	e.Use(middleware.CORS())                     // CORSを許可する
	e.GET("/hello", func(c echo.Context) error { // GETメソッドで/helloにアクセスしたときの処理を定義
		return c.String(200, "Hello World") // 200ステータスコードと"Hello World"を返す
	})
	e.Logger.Fatal(e.Start(":8000")) // サーバーを起動
}
