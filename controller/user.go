package controller

import (
	"net/http"

	// modelパッケージをimportしてmというエイリアスをつける
	m "api-server/model"

	"github.com/labstack/echo"
	"strconv"
)

// 以下、handler関数の定義
// model/user.goの関数を呼び出してリクエストを処理する

// アカウント作成用のhandler
func SignUp(c echo.Context) error {

	name := c.FormValue("name")
	password := c.FormValue("password")

	m.CreateUser(name, password)

	return c.String(http.StatusOK, "user created")

}

// アカウント一覧表示用のhandler(DB確認用)
func GetAccounts(c echo.Context) error {

	users := m.GetUsers()

	return c.JSON(http.StatusOK, users)

}

// アカウント情報表示用のhandler
func GetAccount(c echo.Context) error {

	uid, _ := strconv.Atoi(c.Param("uid"))

	user := m.GetUser(uid)

	return c.JSON(http.StatusOK, user)

}

// アカウント編集用のhandler
func UpdateAccount(c echo.Context) error {

	uid, _ := strconv.Atoi(c.Param("uid"))

	name := c.FormValue("name")
	if name != "" {
		m.UpdateUserName(uid, name)
	}

	password := c.FormValue("password")
	if password != "" {
		m.UpdateUserPassword(uid, password)
	}

	discription := c.FormValue("discription")
	if discription != "" {
		m.UpdateUserDiscription(uid, discription)
	}

	return c.String(http.StatusOK, "user updated")

}

// アカウント削除用のhandler
func DeleteAccount(c echo.Context) error {

	uid, _ := strconv.Atoi(c.Param("uid"))

	m.DeleteUser(uid)

	return c.String(http.StatusOK, "user deleted")

}
