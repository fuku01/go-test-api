package model // modelパッケージであることを宣言

import "gorm.io/gorm" // gormを使用するためのパッケージ

// 「Todo」の型を定義
type Todo struct {
	gorm.Model        // gorm.Modelを埋め込むと、ID, CreatedAt, UpdatedAt, DeletedAtが自動で追加される
	Content    string `json:"content"` // jsonのcontentと紐付け
	UserID     uint   // *UserIDとの関連付け。TodoがUserを1つ持つことを表す。
}

// 「User」の型を定義
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Todos []Todo // *.Todosとの関連付け。UserがTodoを複数持つことを表す。
}
