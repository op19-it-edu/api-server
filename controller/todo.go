package controller

import (
	"net/http"

	// modelパッケージをimportしてmというエイリアスをつける
	m "api-server/model"

	"strconv"

	"github.com/labstack/echo"
	"time"
)

func CreateTodo(c echo.Context) error {

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

	// requestから受け取った各値を変数に格納する
	todo := c.FormValue("todo")
	deadline := c.FormValue("deadline")
	// stringからtimeにparse
	format := "2006-01-02 15:04:05"
	parsedDeadline, _ := time.Parse(format, deadline)

	castedLoginUserID := uint(loginUser.ID)

	m.CreateTodo(castedLoginUserID, todo, parsedDeadline)

	return c.String(http.StatusOK, "todo created")
}

func GetTodo(c echo.Context) error {

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
	// requestから受け取った値を変数に格納する
	targetTodoID, _ := strconv.Atoi(c.Param("id"))
	todo := m.GetTodo(targetTodoID)

	return c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c echo.Context) error {

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
	// requestから受け取った値を変数に格納する
	targetTodoID, _ := strconv.Atoi(c.Param("id"))
	todo := c.FormValue("todo")
	deadline := c.FormValue("deadline")
	done := c.FormValue("done")

	// stringからtimeにparse
	format := "2006-01-02 15:04:05"
	parsedDeadline, _ := time.Parse(format, deadline)

	// string からboolにキャスト
	castedDone, _ := strconv.ParseBool(done)

	m.UpdateTodo(targetTodoID, todo, parsedDeadline, castedDone)

	return c.String(http.StatusOK, "todo updated")

}

func DeleteTodo(c echo.Context) error {

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
	// requestから受け取った値を変数に格納する
	targetTodoID, _ := strconv.Atoi(c.Param("id"))

	// todoIDがDBに存在しない場合はエラーを返す
	if todoID := m.GetTodo(targetTodoID); todoID.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "todo not found",
		}
	}

	m.DeleteTodo(targetTodoID)

	return c.String(http.StatusOK, "todo deleted")
}

// func GetTimeline(c echo.Context) error {

// }
