package main

import (
	// controllerパッケージをimportしてctrというエイリアスをつける（使用するとき長いので）
	ctr "./controller"
	m "./model"

	// go get "github.com/dgrijalva/jwt-go" を実行して事前に取得しとく
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ルーティング設定とサーバーの起動を定義

// エンドポイント（ここから処理が始まる）
func main() {

	// migrationを実行
	m.MigrateDB()

	// Echoのインスタンス作る
	e := echo.New()

	// ログを吐かせるミドルウェア
	e.Use(middleware.Logger())
	// panicが発生したときにサーバーを維持するミドルウェア
	e.Use(middleware.Recover())

	// 以下、pathとhandler関数の紐付け（ルーティング）

	e.GET("/hello", ctr.Hello)
	e.GET("/users", ctr.GetUsersHandler)
	e.GET("/user/create", ctr.CreateUserHandler)
	e.GET("/users/delete", ctr.DeleteUsersHandler)

	// サーバー起動(ここでportを指定する)
	e.Start(":1323")
}
