package controller

import (
	"net/http"
	"time"

	m "api-server/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// jwtのカスタム項目を定義(jwtにuser_idを含めるため)
type jwtCustomClaims struct {
	UID uint `json:"uid"`
	jwt.StandardClaims
}

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: []byte("secret"),
}

func Signup(c echo.Context) error {

	// requestのuserName, userPasswordを変数に格納
	newName := c.FormValue("name")
	newPassword := c.FormValue("password")

	// userName, userPasswordが空の場合はエラーを返す
	if newName == "" || newPassword == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "name or password are empty",
		}
	}

	// 同じUserNameが既に登録済みの場合、エラーを返す
	if newUser := m.FindUser(&m.User{UserName: newName}); newUser.ID != 0 {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "name already exists",
		}
	}

	// 上記処理で問題がなければレコードを新規作成する
	m.CreateUser(newName, newPassword)

	// json形式でuserNameだけbodyに含めてresponseで返却する
	return c.JSON(http.StatusCreated, newName+" signuped!!")
}

func Login(c echo.Context) error {

	loginName := c.FormValue("name")
	loginPassword := c.FormValue("password")

	// userName, userPasswordが一致するレコードが無ければエラーを返す
	loginUser := m.FindUser(&m.User{UserName: loginName})
	if loginUser.ID == 0 || loginUser.UserPassword != loginPassword {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid name or password",
		}
	}

	// jwtカスタム項目としてuserNameを含め、有効期限を24時間にする
	claims := &jwtCustomClaims{
		loginUser.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// カスタム項目を含めてjwtを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成したjwtを暗号化する
	jwtEncoded, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	// responseにjwtを含めて返す
	return c.JSON(http.StatusOK, map[string]string{
		"token": jwtEncoded,
	})
}

// requestのjwtからuserIdを取り出す
func userIdFromToken(c echo.Context) int {
	loginUser := c.Get("user").(*jwt.Token)
	claims := loginUser.Claims.(*jwtCustomClaims)
	userID := claims.UID
	// jwtから取り出したuser_idはuint型なので、handlerでこの関数が呼び出された後の処理で使えるようにintに変換
	castedID := int(userID)
	return castedID
}
