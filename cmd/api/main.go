package main // mainパッケージであることを宣言

import (
	"github.com/fuku01/go-test-api/app/config"
	"github.com/fuku01/go-test-api/app/handler"
	"github.com/fuku01/go-test-api/app/infra/mysgl"
	"github.com/fuku01/go-test-api/app/usecase"
	"github.com/labstack/echo/v4"            // echoを使用するためのパッケージ
	"github.com/labstack/echo/v4/middleware" // echoのミドルウェアを使用するためのパッケージ
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New() // Rest APIを使用するためのインスタンスを作成

	// DBのURLを取得
	DBURL, err := config.GetDBURL()
	if err != nil {
		panic(err) // !エラーがあればプログラムを強制終了
	}
	// DBに接続
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		panic(err) // !エラーがあればプログラムを強制終了
	}

	e.Use(middleware.CORS())                     // CORSを許可する
	e.GET("/hello", func(c echo.Context) error { // ?GETメソッドで/helloにアクセスしたときの処理を定義
		return c.String(200, "Hello World") // ?200ステータスコードと"Hello World"を返す
	})
	tr := mysgl.NewTodoRepository(db)
	tu := usecase.NewTodoUsecase(tr)
	th := handler.NewTodoHandler(tu)
	e.GET("/todos", th.GetAll) // GETメソッドで/todosにアクセスしたときの処理を定義

	e.Logger.Fatal(e.Start(":8000")) // サーバーを起動
}
