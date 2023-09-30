package userctrl

import (
	"net/http"
	"prakerja_batch11/config"
	basemodel "prakerja_batch11/model/base"
	usermodel "prakerja_batch11/model/user"

	"github.com/labstack/echo/v4"
)

func DeleteUser(e echo.Context) error {
	id := e.Param("id")

	var user usermodel.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return e.JSON(http.StatusNotFound, basemodel.Response{
			Status:  false,
			Message: "User not found",
			Data:    err,
		})
	}

	config.DB.Delete(&user)
	return e.JSON(http.StatusOK, basemodel.Response{
		Status:  true,
		Message: "User data has been successfully deleted",
		Data:    nil,
	})
}
