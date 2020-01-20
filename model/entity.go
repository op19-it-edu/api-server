package model

import "github.com/jinzhu/gorm"

// 構造体などを定義する

// usersテーブルの定義
type User struct {
	// gorm.ModelはID, CreatedAt, UpdatedAt, DeletedAtを自動挿入する
	gorm.Model
	// タグでDBに作成されるカラム名やデータ型を指定する
	Name        string `gorm:"type:text;column:user_name"`
	Email       string `gorm:"type:text;column:user_password"`
	Password    string `gorm:"type:text;column:user_email"`
	Discription string `gorm:"type:text;column:user_discription"`
}
