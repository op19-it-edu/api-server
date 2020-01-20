package model

import (
	// gormパッケージ
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 共通的なDB操作を定義

func InitDB() *gorm.DB {

	// DBサーバーの情報
	DBMS := "mysql"
	user := "root"
	password := "password"
	protocol := "tcp(127.0.0.1:3306)"
	DBname := "todotweet_db"

	DBInfo := user + ":" + password + "@" + protocol + "/" + DBname + "?parseTime=true"

	// DB接続
	db, err := gorm.Open(DBMS, DBInfo)
	if err != nil {
		panic("failed to connect database")
	}

	// 接続したDBを返す
	return db
}

func MigrateDB() {

	// 接続したDBを変数dbに格納
	db := InitDB()

	// DBをentityの定義に従ってmigrationする
	db.AutoMigrate(&User{})

	// DBを閉じる
	defer db.Close()

}
