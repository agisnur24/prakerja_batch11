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

	if err := config.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		return e.JSON(http.StatusBadRequest, basemodel.Response{
			Status:  false,
			Message: "Email already exist",
			Data:    nil,
		})
	}

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

	return e.JSON(http.StatusCreated, basemodel.Response{
		Status:  true,
		Message: "Registration successful",
		Data:    nil,
	})
}
