package auth

import (
	"net/http"
	"prakerja_batch11/config"
	midware "prakerja_batch11/middleware"
	basemodel "prakerja_batch11/model/base"
	"prakerja_batch11/model/login"
	usermodel "prakerja_batch11/model/user"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LogInController(e echo.Context) error {
	var logIn login.LoginRequest
	e.Bind(logIn)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(logIn.Password), 12)
	logIn.Password = string(hashPassword)

	var user usermodel.User
	email := e.QueryParam(logIn.Email)
	if findEmail := config.DB.Where("email = ?", email).First(&user).Error; findEmail != nil {
		if gorm.IsRecordNotFoundError(findEmail) {
			return e.JSON(http.StatusNotFound, basemodel.Response{
				Status:  false,
				Message: "Email not found",
				Data:    nil,
			})
		}
	}

	password := bcrypt.CompareHashAndPassword([]byte(logIn.Password), []byte(user.Password))

	if logIn.Email == user.Email && password != nil {
		return e.JSON(http.StatusUnauthorized, basemodel.Response{
			Status:  false,
			Message: "Password is incorrect",
			Data:    nil,
		})
	}

	if logIn.Email != user.Email && password == nil {
		return e.JSON(http.StatusUnauthorized, basemodel.Response{
			Status:  false,
			Message: "Email is incorrect",
			Data:    nil,
		})
	}

	var loginResponse = &login.LoginResponse{}

	if logIn.Email == user.Email && password != nil {
		genTOken := midware.GenerateTokenJWT(user.Id, user.Name)
		loginResponse.Token = genTOken
		return nil
	}

	return e.JSON(http.StatusOK, basemodel.Response{
		Status:  true,
		Message: "Loged in",
		Data:    loginResponse,
	})
}
