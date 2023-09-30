package userctrl

import (
	"net/http"
	"prakerja_batch11/config"
	basemodel "prakerja_batch11/model/base"
	usermodel "prakerja_batch11/model/user"

	"github.com/labstack/echo/v4"
)

func FindUser(e echo.Context) error {
	var findUser usermodel.FindUserRequest
	e.Bind(&findUser)

	var user usermodel.User
	if err := config.DB.Where("name = ?", findUser.Name).First(&user).Error; err != nil {
		return e.JSON(http.StatusNotFound, basemodel.Response{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	userResponse := usermodel.ToUserResponse(&user)
	return e.JSON(http.StatusOK, basemodel.Response{
		Status:  true,
		Message: "User matches",
		Data:    userResponse,
	})
}
