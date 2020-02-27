package route

import (
	ctr "api-server/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter() *echo.Echo {

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

	return e
}
