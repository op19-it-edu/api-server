package middleware

import (
	"net/http"

	ctr "api-server/controller"
	m "api-server/model"

	"github.com/labstack/echo"
)

// 各ctrの関数が実行される前に実行される関数をmiddlewareとして定義
// jwtから取り出したuser_idがDBに無ければエラーを返す

func JWTCustomMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := ctr.UserIdFromToken(c)
			loginUser := m.GetUser(userID)

			// userIDがゼロの場合はエラーを返す
			if loginUser.ID == 0 {
				return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "user not signuped",
				}
			}

			err := next(c)
			return err
		}
	}
}
