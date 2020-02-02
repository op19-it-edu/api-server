package controller

import (
	"net/http"

	// modelパッケージをimportしてmというエイリアスをつける
	m "api-server/model"

	"strconv"

	"github.com/labstack/echo"
)

// 以下、handler関数の定義
// model/user.goの関数を呼び出してリクエストを処理する

// ログインユーザーのアカウント情報表示用のhandler
func GetSelfAccount(c echo.Context) error {

	// jwtから取り出したuserIDがテーブルに存在しなければエラーを返す
	userID := userIdFromToken(c)
	loginUser := m.GetUser(userID)

	// userIDがゼロの場合はエラーを返す
	if loginUser.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not signuped",
		}
	}

	return c.JSON(http.StatusOK, loginUser)
}

// アカウント編集用のhandler
func UpdateAccount(c echo.Context) error {

	// jwtから取り出したuuserIDがテーブルに存在しなければエラーを返す
	userID := userIdFromToken(c)
	loginUser := m.GetUser(userID)

	// userIDがゼロの場合はエラーを返す
	if loginUser.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not signuped",
		}
	}

	// requestから受け取った各値を変数に格納する
	name := c.FormValue("name")
	password := c.FormValue("password")
	description := c.FormValue("description")

	m.UpdateUser(userID, name, password, description)

	return c.String(http.StatusOK, "user updated")

}

// アカウント削除用のhandler
func DeleteAccount(c echo.Context) error {

	// jwtから取り出したuidがテーブルに存在しなければエラーを返す
	userID := userIdFromToken(c)
	loginUser := m.GetUser(userID)

	// userIDがゼロの場合はエラーを返す
	if loginUser.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not signuped",
		}
	}

	m.DeleteUser(userID)

	return c.String(http.StatusOK, "user deleted")

}

// アカウント一覧表示用のhandler(followeeを検索するため全ユーザーを返す)
func GetAllAccounts(c echo.Context) error {

	// jwtから取り出したuserIDがテーブルに存在しなければエラーを返す
	userID := userIdFromToken(c)
	loginUser := m.GetUser(userID)

	// userIDがゼロの場合はエラーを返す
	if loginUser.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not signuped",
		}
	}

	// 全ユーザー情報をDBから取得する
	users := m.GetUsers()

	return c.JSON(http.StatusOK, users)

}

// 該当ユーザーのアカウント情報表示用のhandler(後々、該当ユーザーのtodoも返す)
func GetAccount(c echo.Context) error {

	// jwtから取り出したuserIDがテーブルに存在しなければエラーを返す
	userID := userIdFromToken(c)
	loginUser := m.GetUser(userID)

	// userIDがゼロの場合はエラーを返す
	if loginUser.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not signuped",
		}
	}

	// requestから受け取った各値を変数に格納する
	uid, _ := strconv.Atoi(c.Param("id"))

	// DBから該当ユーザーを取り出す
	targetUser := m.GetUser(uid)

	// targetUserIDがDBに無ければエラーを返す
	if targetUser.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not found",
		}
	}

	// passworは返したくないので、空にしておく
	targetUser.UserPassword = ""

	return c.JSON(http.StatusOK, targetUser)

}
