package auth

import (
	"net/http"
	"prakerja_batch11/config"
	midware "prakerja_batch11/middleware"
	basemodel "prakerja_batch11/model/base"
	"prakerja_batch11/model/login"
	usermodel "prakerja_batch11/model/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LogInController(e echo.Context) error {
	var logIn login.LoginRequest
	e.Bind(&logIn)

	var user usermodel.User
	if findEmail := config.DB.Where("email = ?", logIn.Email).First(&user).Error; findEmail != nil {
		return e.JSON(http.StatusNotFound, basemodel.Response{
			Status:  false,
			Message: "Email not found",
			Data:    nil,
		})
	}

	var loginResponse = login.LoginResponse{}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logIn.Password))
	if err != nil {
		return e.JSON(http.StatusUnauthorized, basemodel.Response{
			Status:  false,
			Message: "Incorrect password",
			Data:    nil,
		})
	}

	genToken := midware.GenerateTokenJWT(user.Id, user.Name)
	loginResponse.Token = genToken

	return e.JSON(http.StatusOK, basemodel.Response{
		Status:  true,
		Message: "Loged in",
		Data:    loginResponse,
	})
}
