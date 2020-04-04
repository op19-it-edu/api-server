package controller

import (
	"net/http"

	// modelパッケージをimportしてmというエイリアスをつける
	m "api-server/model"

	"github.com/labstack/echo"
)

func GetTimeline(c echo.Context) error {

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

	castedID := uint(loginUser.ID)
	followeeIDs := m.GetFolloweeID(castedID)

	// 取り出したFolloweeIDをuintに変換
	castedIDs := make([]uint, len(followeeIDs))
	for n := range followeeIDs {
		castedIDs[n] = uint(followeeIDs[n])
	}

	followeeTodos := m.GetFolloweeTodos(castedIDs)

	return c.JSON(http.StatusOK, followeeTodos)
}
