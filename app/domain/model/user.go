package model

import "gorm.io/gorm"

// @「DBのテーブル」と紐付けるための「User」型を定義する。
type User struct {
	gorm.Model          // gorm.Modelを埋め込む。（ID, created_at, updated_at, deleted_atが含まれる。）
	Name         string `json:"name"`         // jsonの「name」と紐付け
	Firebase_uid string `json:"firebase_uid"` // jsonのfirebase_uidと紐付け
	Todos        []Todo // ?Todosテーブルとの関連付け。UserがTodoを複数持つことを表す。
}
