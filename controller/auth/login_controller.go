package auth

import (
	"net/http"
	"prakerja_batch11/config"
	"prakerja_batch11/middleware"
	"prakerja_batch11/model/base"
	usermodel "prakerja_batch11/model/user"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LogInController(e echo.Context) error {
	var logIn usermodel.LoginRequest
	e.Bind(logIn)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(logIn.Password), 12)
	logIn.Password = string(hashPassword)

	var user usermodel.User
	email := e.QueryParam(logIn.Email)
	if findEmail := config.DB.Where("email = ?", email).First(&user).Error; findEmail != nil {
		if gorm.IsRecordNotFoundError(findEmail) {
			return e.JSON(http.StatusNotFound, base.Response{
				Status:  false,
				Message: "Email not found",
				Data:    nil,
			})
		}
	}

	password := bcrypt.CompareHashAndPassword([]byte(logIn.Password), []byte(user.Password))

	if logIn.Email != user.Email || password != nil {
		return e.JSON(http.StatusUnauthorized, base.Response{
			Status:  false,
			Message: "Email or password is incorrect",
			Data:    nil,
		})
	}

	var loginResponse = usermodel.LoginResponse{
		Token: middleware.GenerateTokenJWT(user.Id, user.Name),
	}

	return e.JSON(http.StatusOK, base.Response{
		Status:  true,
		Message: "Loged In",
		Data:    loginResponse,
	})
}
