package main // mainパッケージであることを宣言

import (
	"github.com/fuku01/go-test-api/app/config"       // configパッケージを使用するためのパッケージ
	"github.com/fuku01/go-test-api/app/domain/model" // modelパッケージを使用するためのパッケージ
	"gorm.io/driver/mysql"                           // mysqlを使用するためのパッケージ
	"gorm.io/gorm"                                   // gormを使用するためのパッケージ
)

// ! マイグレーションを実行する関数(migrationフォルダ内で「go run main.go」コマンドで実行する)

func main() {
	// DBのURLを取得
	DBURL, err := config.GetDBURL()
	if err != nil {
		panic(err)
	}
	// DBに接続
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// @「User]と「Todo」の型（モデル）に基づいてテーブルを作成
	db.AutoMigrate(&model.User{}, &model.Todo{}, &model.Tag{})
	if err != nil {
		panic(err)
	}

}
