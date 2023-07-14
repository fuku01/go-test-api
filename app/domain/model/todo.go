package model // modelパッケージであることを宣言

import "gorm.io/gorm" // gormを使用するためのパッケージ

// Todo構造体を定義
type Todo struct {
	gorm.Model        // gorm.Modelを埋め込む
	Content    string `json:"content"` // Contentを定義
}
