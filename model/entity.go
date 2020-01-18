package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 構造体などを定義する

// usersテーブル
type User struct {
	// gorm.ModelはID, CreatedAt, UpdatedAt, DeletedAtを自動挿入する
	gorm.Model
	// タグでDBに作成されるカラム名やデータ型を指定する
	Name        string `gorm:"type:text;column:user_name"`
	Password    string `gorm:"type:text;column:user_password"`
	Discription string `gorm:"type:text;column:user_discription"`

	// has manyの関係を明示
	Relations  []Relation
	TodoTweets []TodoTweet
}

// relationsテーブル
type Relation struct {
	gorm.Model

	// 外部キー（usersテーブル）
	UserID   uint
	FollowID uint
}

// todo_tweetsテーブル
type TodoTweet struct {
	gorm.Model

	// 外部キー（usersテーブル）
	UserID uint

	// typeでデータ型を指定
	Todo     string    `gorm:"type:text"`
	Deadline time.Time `gorm:"type:datetime"`
	Done     bool      `gorm:"type:boolean"`
}
