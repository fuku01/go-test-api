package model

import "gorm.io/gorm"

// 「User」の型を定義
type User struct {
	gorm.Model          // gorm.Modelを埋め込む
	Name         string `json:"name"`         // jsonの「name」と紐付け
	Firebase_uid string `json:"firebase_uid"` // jsonのfirebase_uidと紐付け
	Todos        []Todo // ?.Todosとの関連付け。UserがTodoを複数持つことを表す。
}
