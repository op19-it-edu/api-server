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

// user_idがuidに一致するレコードを返す関数
func GetUser(uid int) User {

	db := InitDB()
	user := User{}

	db.First(&user, uid)

	defer db.Close()

	return user

}

func UpdateUser(uid int, name, password, discription string) {

	db := InitDB()
	user := User{}

	if name != "" {
		db.Model(&user).Where("id = ?", uid).Update("user_name", name)
	}

	if password != "" {
		db.Model(&user).Where("id = ?", uid).Update("user_password", password)
	}

	if discription != "" {
		db.Model(&user).Where("id = ?", uid).Update("user_discription", discription)
	}

	defer db.Close()

}

// user_idがuidに一致するレコードを削除する関数
func DeleteUser(uid int) {

	db := InitDB()
	user := User{}

	db.Where("id = ?", uid).Delete(&user)

	defer db.Close()

}
