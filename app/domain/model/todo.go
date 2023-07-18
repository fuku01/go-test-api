package model // modelパッケージであることを宣言

import "gorm.io/gorm" // gormを使用するためのパッケージ

// 「Todo」の型を定義
type Todo struct {
	gorm.Model        // gorm.Modelを埋め込むと、ID, CreatedAt, UpdatedAt, DeletedAtが自動で追加される
	Content    string `json:"content"` // jsonのcontentと紐付け
}
