package controller

import (
	"net/http"

	// modelパッケージをimportしてmというエイリアスをつける
	m "api-server/model"

	"strconv"

	"github.com/labstack/echo"
)

func CreateRelation(c echo.Context) error {

	// jwtからuserIDを取り出す
	userID := UserIdFromToken(c)

	// requestから受け取った値を変数に格納する
	followeeID, _ := strconv.Atoi(c.Param("id"))

	// followeeIDがDBに無ければエラーを返す
	if followee := m.GetUser(followeeID); followee.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not found",
		}
	}

	castedLoginUserID := uint(userID)
	castedFolloweeID := uint(followeeID)
	// RelationsテーブルのuserIDにはLoginUserIDを、followIDにはFolloweeIDを保存
	m.CreateRelation(castedLoginUserID, castedFolloweeID)

	return c.JSON(http.StatusOK, "creating relation succeeded")
}

func GetFollowee(c echo.Context) error {

	// requestから受け取った値を変数に格納する
	targetUserID, _ := strconv.Atoi(c.Param("id"))

	// targetUserIDがDBに無ければエラーを返す
	if targetUser := m.GetUser(targetUserID); targetUser.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not found",
		}
	}

	castedID := uint(targetUserID)
	followeeIDs := m.GetFolloweeID(castedID)
	followee := m.GetRelators(followeeIDs)

	return c.JSON(http.StatusOK, followee)
}

func GetFollower(c echo.Context) error {

	// requestから受け取った値を変数に格納する
	targetUserID, _ := strconv.Atoi(c.Param("id"))

	// targetUserIDがDBに無ければエラーを返す
	if targetUser := m.GetUser(targetUserID); targetUser.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not found",
		}
	}

	castedID := uint(targetUserID)
	followerIDs := m.GetFollowerID(castedID)
	follower := m.GetRelators(followerIDs)

	return c.JSON(http.StatusOK, follower)
}

func DeleteRelation(c echo.Context) error {

	// jwtからuserIDを取り出す
	userID := UserIdFromToken(c)

	// requestから受け取った値を変数に格納する
	followeeID, _ := strconv.Atoi(c.Param("id"))

	// followeeIDがDBに存在しない場合はエラーを返す
	if followee := m.GetUser(followeeID); followee.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not found",
		}
	}

	castedLoginUserID := uint(userID)
	castedFolloweeID := uint(followeeID)
	m.DeleteRelation(castedLoginUserID, castedFolloweeID)

	return c.JSON(http.StatusOK, "deleting relation succeeded")
}
