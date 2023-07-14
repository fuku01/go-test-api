package main // mainパッケージであることを宣言

import (
	"github.com/fuku01/go-test-api/app/config"
	"github.com/fuku01/go-test-api/app/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// マイグレーションを実行する関数(migrationフォルダ内で「go run main.go」コマンドで実行する)
func main() {
	// DBのURLを取得
	DBURL, err := config.GetDBURL()
	if err != nil {
		panic(err) // エラーがあればプログラムを強制終了
	}
	// DBに接続
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		panic(err) // エラーがあればプログラムを強制終了
	}
	// Todo構造体に基づいてテーブルを作成
	err = db.AutoMigrate(&model.Todo{})
	if err != nil {
		panic(err) // エラーがあればプログラムを強制終了
	}

}
