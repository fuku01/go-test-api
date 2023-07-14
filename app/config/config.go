package config // configパッケージであることを宣言

import (
	"errors"
	"fmt"
	"os"
)

// Config構造体を定義
type Config struct {
	DBURL string
}

// DBのURLを取得する関数
func GetDBURL() (string, error) {
	DBName := os.Getenv("MYSQL_DATABASE")
	if DBName == "" {
		return "", errors.New("MYSQL_DATABASEの環境変数が設定されていません")
	}
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		return "", errors.New("MYSQL_HOSTの環境変数が設定されていません")
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		return "", errors.New("MYSQL_USERの環境変数が設定されていません")
	}
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		return "", errors.New("MYSQL_PASSWORDの環境変数が設定されていません")
	}
	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort == "" {
		return "", errors.New("MYSQL_PORTの環境変数が設定されていません")
	}
	DBURL := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(DBURL)
	return DBURL, nil
}
