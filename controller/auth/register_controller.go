package auth

import (
	"net/http"
	"prakerja_batch11/config"
	"prakerja_batch11/model/base"
	"prakerja_batch11/model/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(e echo.Context) error {
	var user user.User
	e.Bind(&user)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashPassword)

	result := config.DB.Create(&user)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, base.Response{
			Status:  false,
			Message: "Failed add data to database",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, base.Response{
		Status:  true,
		Message: "Success register",
		Data:    nil,
	})
}
