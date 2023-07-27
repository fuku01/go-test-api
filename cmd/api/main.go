package main // mainパッケージであることを宣言

import (
	"context"
	"log"

	"firebase.google.com/go/auth"
	"go.uber.org/fx"

	"github.com/fuku01/go-test-api/app/config"
	"github.com/fuku01/go-test-api/app/handler"
	"github.com/fuku01/go-test-api/app/infra/firebase"
	"github.com/fuku01/go-test-api/app/infra/mysgl"
	"github.com/fuku01/go-test-api/app/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	return e
}

func NewFirebaseAuth() (*auth.Client, error) {
	ctx := context.Background()
	firebaseApp, err := config.GetFirebaseAuth()
	if err != nil {
		return nil, err
	}
	authClient, err := firebaseApp.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return authClient, nil
}

func NewDB() (*gorm.DB, error) {
	DBURL, err := config.GetDBURL()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	fx.New(
		fx.Provide(
			NewEcho,
			NewFirebaseAuth,
			NewDB,

			// 依存性を注入します
			mysgl.NewTodoRepository,
			mysgl.NewUserRepository,
			firebase.NewFirebaseAuthRepository,

			usecase.NewTodoUsecase,
			usecase.NewUserUsecase,

			handler.NewTodoHandler,
			handler.NewUserHandler,
		),
		fx.Invoke(RegisterRoutes), // RegisterRoutes関数を起動
	).Run()
}

// ! ルーティング
func RegisterRoutes(lc fx.Lifecycle, th handler.TodoHandler, uh handler.UserHandler, e *echo.Echo) {
	e.GET("/todos", th.GetAll)                         // GETメソッドで/todosにアクセスしたときの処理を定義
	e.POST("/create", th.Create)                       // POSTメソッドで/createにアクセスしたときの処理を定義
	e.DELETE("/delete/:ID", th.Delete)                 // DELETEメソッドで/deleteにアクセスしたときの処理を定義
	e.GET("/loginuser", uh.GetLoginUser)               // GETメソッドで/loginuserにアクセスしたときの処理を定義
	e.GET("/todoswithtags", th.GetAllWithTags)         // GETメソッドで/todoswithtagsにアクセスしたときの処理を定義
	e.POST("/createwithtags", th.CreateWithTags)       // POSTメソッドで/createwithtagsにアクセスしたときの処理を定義
	e.DELETE("/deletewithtags/:ID", th.DeleteWithTags) // DELETEメソッドで/deletewithtagsにアクセスしたときの処理を定義
	e.PUT("/editwithtags/:ID", th.EditWithTags)        // PUTメソッドで/editwithtagsにアクセスしたときの処理を定義

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// 非同期にサーバーを起動する
			go func() {
				if err := e.Start(":8000"); err != nil { // サーバーを起動する
					log.Printf("echo start error: %v\n", err) // サーバーが起動できなかったときのエラー処理
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// サーバーを停止する
			return e.Shutdown(ctx)
		},
	})
}
