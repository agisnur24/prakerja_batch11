package auth

import (
	"net/http"
	"prakerja_batch11/config"
	basemodel "prakerja_batch11/model/base"
	usermodel "prakerja_batch11/model/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(e echo.Context) error {
	var user usermodel.User
	e.Bind(&user)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashPassword)

	result := config.DB.Create(&user)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, basemodel.Response{
			Status:  false,
			Message: "Registration failed",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, basemodel.Response{
		Status:  true,
		Message: "Registration successful",
		Data:    nil,
	})
}
