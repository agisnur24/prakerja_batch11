package userctrl

import (
	"net/http"
	"prakerja_batch11/config"
	basemodel "prakerja_batch11/model/base"
	usermodel "prakerja_batch11/model/user"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	var users []usermodel.User
	var userResponse []usermodel.UserResponse

	result := config.DB.Find(&users, userResponse)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, basemodel.Response{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, basemodel.Response{
		Status:  true,
		Message: "Success",
		Data:    userResponse,
	})
}