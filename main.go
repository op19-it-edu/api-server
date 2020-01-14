package main

import (
	// controllerパッケージをimport
	"./controller"

	// go get "github.com/dgrijalva/jwt-go" を実行して事前に取得しとく
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// エンドポイント（ここから処理が始まる）
func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// ログを吐かせるミドルウェア
	e.Use(middleware.Logger())
	// panicが発生したときにサーバーを維持するミドルウェア
	e.Use(middleware.Recover())

	// pathと関数の紐付け（ルーティング）
	// http://localhost:1323/hello を叩くとHelloが呼ばれる
	e.GET("/hello", controller.Hello)

	// サーバー起動(ここでportを指定する)
	e.Start(":1323")
}
