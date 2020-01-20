package main

import (
	// controllerパッケージをimportしてctrというエイリアスをつける（使用するとき長いので）
	ctr "api-server/controller"
	m "api-server/model"

	// go get "github.com/dgrijalva/jwt-go" を実行して事前に取得しとく
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ルーティング設定とサーバーの起動を定義

// エントリーポイント（ここから処理が始まる）
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

	e.POST("/signup", ctr.SignUp)
	e.GET("/users", ctr.GetAccounts)
	e.GET("/user/:uid", ctr.GetAccount)
	e.PUT("/user/:uid", ctr.UpdateAccount)
	e.DELETE("/user/:uid", ctr.DeleteAccount)

	// サーバー起動(ここでportを指定する)
	e.Start(":1323")
}
