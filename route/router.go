package route

import (
	ctr "api-server/controller"
	mid "api-server/middleware"

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
	e.POST("/signup", ctr.Signup)
	e.POST("/login", ctr.Login)

	api := e.Group("/api/v1")
	api.Use(middleware.JWTWithConfig(ctr.Config))
	api.Use(mid.JWTCustomMiddleware())

	// 以下、JWT認証が必要
	// accountに関するRUD
	api.GET("/user", ctr.GetSelfAccount)
	api.PUT("/user", ctr.UpdateAccount)
	api.DELETE("/user", ctr.DeleteAccount)

	// followeeを検索するため全ユーザーを返す
	api.GET("/users", ctr.GetAllAccounts)
	// followeeを新規作成する
	api.POST("/relation/:id", ctr.CreateRelation)
	api.DELETE("/relation/:id", ctr.DeleteRelation)
	api.GET("/followee/:id", ctr.GetFollowee)
	api.GET("/follower/:id", ctr.GetFollower)

	// 該当ユーザー情報を表示するため(後々、該当ユーザーのtodoも返す)
	api.GET("/user/:id", ctr.GetAccount)

	return e
}
