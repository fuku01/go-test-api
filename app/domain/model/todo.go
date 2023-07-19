package model

import "gorm.io/gorm"

// !このファイルでは、「DBのテーブル」と紐付けるための型を定義する。

// （※ gorm.Modelを埋め込むと、ID, CreatedAt, UpdatedAt, DeletedAtが自動で追加される。）

// 「Todo」の型を定義
type Todo struct {
	gorm.Model        // gorm.Modelを埋め込む
	Content    string `json:"content"` // jsonの「content」と紐付け
	UserID     uint   // ?UserIDとの関連付け。TodoがUserを1つ持つことを表す。
}
