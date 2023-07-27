package model

import "gorm.io/gorm"

// @「DBのテーブル」と紐付けるための「Tag」モデル型を定義する。

type Tag struct {
	gorm.Model // gorm.Modelを埋め込む。（ID, created_at, updated_at, deleted_atが含まれる。）

	Name string `json:"name"` // jsonの「name」と紐付け

	Todo *Todo // Todo構造体との関連付け。（TagがTodoを1つ持つことを表す。）

	TodoID uint `gorm:"not null"` // Todoテーブルとの関連付け。(TagがTodoを1つ持つことを表す。) ※not nullとは、空の値を許可しないという意味。
}
