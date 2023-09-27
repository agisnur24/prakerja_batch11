package auth

import (
	"net/http"
	"prakerja_batch11/config"
	"prakerja_batch11/middleware"
	"prakerja_batch11/model/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c echo.Context) error {
	var user user.User
	c.Bind(&user)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashPassword)

	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed add data to database",
			Data:    nil,
		})
	}

	var authResponse = usermodel.ResponseAuth{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: middleware.GenerateTokenJWT(user.Id, user.Name),
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success register",
		Data:    authResponse,
	})
}
