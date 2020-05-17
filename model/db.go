package model

import (
	conf "api-server/config"

	// gormパッケージ
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 共通的なDB操作を定義

func InitDB() *gorm.DB {

	// デコードしたjsonをcfに格納
	cf := conf.DecodeJson()

	// DBサーバーの情報
	DBMS := cf.DBMS
	user := cf.User
	password := cf.Password
	protocol := cf.Protocol
	DBname := cf.DBName

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
	// AddForeignKey()で外部キーを指定
	db.AutoMigrate(&Relation{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("follow_id", "users(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&TodoTweet{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	// DBを閉じる
	defer db.Close()

}
