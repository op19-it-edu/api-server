package model

// usersテーブルの操作を定義

// nameとpasswordを引数にとって新規レコードを追加する関数
func CreateUser(name, password string) {

	db := InitDB()
	user := User{
		UserName:     name,
		UserPassword: password,
	}

	db.Create(&user)

	defer db.Close()

}

// usersテーブルの全レコードを取得する関数
func GetUsers() []User {

	db := InitDB()
	users := []User{}

	db.Find(&users)

	defer db.Close()

	return users

}

// userNameが一致するレコードを返す
func FindUser(u *User) User {
	db := InitDB()
	user := User{}
	db.Where(u).First(&user)
	return user
}

// userIDがuidに一致するレコードを返す関数
func GetUser(uid int) User {

	db := InitDB()
	user := User{}

	db.First(&user, uid)

	defer db.Close()

	return user

}

// userIDが一致するレコード全てを返す(fllowee, followerを返すため)
func GetRelators(IDs []int) []User {
	db := InitDB()
	users := []User{}

	for _, v := range IDs {
		user := User{}
		db.Where("id = ?", v).Find(&user)
		user.UserPassword = ""
		users = append(users, user)
	}

	return users

}

func UpdateUser(uid int, name, password, description string) {

	db := InitDB()
	user := User{}

	// 空白のフィールドは無視してレコードを更新する
	db.Model(&user).Where("id = ?", uid).Updates(User{UserName: name, UserPassword: password, UserDescription: description})

	defer db.Close()

}

// userIDがuidに一致するレコードを削除する関数
func DeleteUser(uid int) {

	db := InitDB()
	user := User{}

	// Deleteする際のid指定はuintじゃないといけないので、intからuintにキャストしておく
	castedID := uint(uid)

	// キャストしたuserIDを構造体に代入する
	user.ID = castedID

	db.Unscoped().Delete(&user)

	defer db.Close()

}
