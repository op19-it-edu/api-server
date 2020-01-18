package model

// usersテーブルの操作を定義

// usersテーブルの全レコードを取得する関数
func GetUsers() []User {

	db := InitDB()

	// 構造体Userの配列をusersに格納
	users := []User{}

	// DBのusersテーブルの全レコードを取得して、usersに格納
	db.Find(&users)

	defer db.Close()

	// usersを返す
	return users
}

// usersテーブルに新規のユーザーを追加する関数（テスト用のユーザーを構造体に格納して追加）
func CreateUser() {

	db := InitDB()

	// 構造体Userを初期化（テスト用のユーザー情報を格納）してuserに格納
	// これが追加するレコード
	user := User{
		Name:        "hogehoge",
		Password:    "password",
		Discription: "fugafuga",
	}

	// DBにuserレコードを追加
	db.Create(&user)

	defer db.Close()

}

// usersテーブルの全レコードを削除する関数
func DeleteUsers() {

	db := InitDB()

	users := User{}

	// usersテーブルの全レコードを削除
	db.Delete(&users)

	defer db.Close()

}
