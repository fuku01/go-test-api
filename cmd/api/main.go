package main // mainパッケージであることを宣言

import (
	"context"

	"github.com/fuku01/go-test-api/app/config"
	"github.com/fuku01/go-test-api/app/handler"
	"github.com/fuku01/go-test-api/app/infra/mysgl"
	"github.com/fuku01/go-test-api/app/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()          // ?Rest APIを使用するためのインスタンスを作成。(echoを使用するために必要)
	e.Use(middleware.CORS()) // ?CORSを許可する。(フロントとの通信を許可するために必要)

	// !Firebaseの認証情報を取得
	ctx := context.Background()
	firebaseApp, err := config.GetFirebaseAuth()
	if err != nil {
		panic(err) // エラーがあれば強制終了
	}
	client, err := firebaseApp.Auth(ctx)
	if err != nil {
		panic(err) // エラーがあれば強制終了
	}

	//!DBの接続情報を取得する
	DBURL, err := config.GetDBURL() // *DBのURLを取得
	if err != nil {
		panic(err) // エラーがあれば強制終了
	}
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{}) // *DBに接続
	if err != nil {
		panic(err) // エラーがあれば強制終了
	}

	// !依存関係の注入
	tr := mysgl.NewTodoRepository(db)
	ur := mysgl.NewUserRepository(db)
	tu := usecase.NewTodoUsecase(tr)
	uu := usecase.NewUserUsecase(ur, client)
	th := handler.NewTodoHandler(tu, uu)

	// !ルーティング
	e.GET("/todos", th.GetAll)         // GETメソッドで/todosにアクセスしたときの処理を定義
	e.POST("/create", th.Create)       // POSTメソッドで/createにアクセスしたときの処理を定義
	e.DELETE("/delete/:ID", th.Delete) // DELETEメソッドで/deleteにアクセスしたときの処理を定義

	// !サーバーを起動
	e.Logger.Fatal(e.Start(":8000")) // サーバーをポート8000で立ち上げる

	//// e.GET("/hello", func(c echo.Context) error { // GETメソッドで/helloにアクセスしたときの処理を定義
	//// 	return c.String(200, "Hello World") // 200ステータスコードと"Hello World"を返す
	//// })
}
