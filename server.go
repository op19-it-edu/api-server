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

	// 認証なし
	e.POST("/signup", ctr.SignUp)

	api := e.Group("/api/v1")

	// 以下、JWT認証が必要
	api.GET("/users", ctr.GetAccounts)
	api.GET("/user/:uid", ctr.GetAccount)
	api.PUT("/user/:uid", ctr.UpdateAccount)
	api.DELETE("/user/:uid", ctr.DeleteAccount)

	// サーバー起動(ここでportを指定する)
	e.Start(":1323")
}
