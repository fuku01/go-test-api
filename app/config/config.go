package config // configパッケージであることを宣言

import (
	"errors" // エラーを扱うためのパッケージ
	"fmt"    // フォーマットを扱うためのパッケージ
	"os"     // OSの機能を扱うためのパッケージ
)

// Configの型を定義
type Config struct {
	DBURL string
}

// * maigration実行処理で使用する「DBのURL」を取得する関数
func GetDBURL() (string, error) {
	DBName := os.Getenv("MYSQL_DATABASE")
	if DBName == "" {
		return "", errors.New("MYSQL_DATABASEの環境変数が設定されていません") // !環境変数が設定されていなければエラーを返す
	}
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		return "", errors.New("MYSQL_HOSTの環境変数が設定されていません") // !環境変数が設定されていなければエラーを返す
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		return "", errors.New("MYSQL_USERの環境変数が設定されていません") // !環境変数が設定されていなければエラーを返す
	}
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		return "", errors.New("MYSQL_PASSWORDの環境変数が設定されていません") // !環境変数が設定されていなければエラーを返す
	}
	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort == "" {
		return "", errors.New("MYSQL_PORTの環境変数が設定されていません") // !環境変数が設定されていなければエラーを返す
	}
	DBURL := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local" // DBのURLを定義
	fmt.Println(DBURL)
	return DBURL, nil // DBのURLを返す
}
