package model

import "gorm.io/gorm"

// @「DBのテーブル」と紐付けるための「Todo」モデル型を定義する。

type Todo struct {
	gorm.Model //  gorm.Modelを埋め込む。（ID, created_at, updated_at, deleted_atが含まれる。）

	Content string `json:"content"` // jsonの「content」と紐付け

	UserID uint // Userテーブルとの関連付け。(TodoがUserを1つ持つことを表す。)
}
