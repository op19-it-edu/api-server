package controller

import (
	"net/http"

	// modelパッケージをimportしてmというエイリアスをつける
	m "../model"

	"github.com/labstack/echo"
)

// 以下、handler関数の定義

// Hello, Worldを返す関数
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// usersテーブルの全レコード取得して返す関数
func GetUsersHandler(c echo.Context) error {
	// modelパッケージのGetUsers()を呼び出す
	users := m.GetUsers()
	return c.JSON(http.StatusOK, users)
}

// usersテーブルに新規レコード追加する関数
func CreateUserHandler(c echo.Context) error {
	// modelパッケージのCreateUser()を呼び出す
	m.CreateUser()
	return c.String(http.StatusOK, "user created")
}

// usersテーブルの全レコードを削除する関数
func DeleteUsersHandler(c echo.Context) error {
	// modelパッケージのDeleteUsersを呼び出す
	m.DeleteUsers()
	return c.String(http.StatusOK, "users deleted")
}
