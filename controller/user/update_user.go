package userctrl

import (
	"net/http"
	"prakerja_batch11/config"
	basemodel "prakerja_batch11/model/base"
	usermodel "prakerja_batch11/model/user"

	"github.com/labstack/echo/v4"
)

func UpdateUser(e echo.Context) error {
	var updateRequest usermodel.UpdateUser
	e.Bind(&updateRequest)

	id := e.Param("id")

	var user usermodel.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return e.JSON(http.StatusNotFound, basemodel.Response{
			Status:  false,
			Message: "User not found",
			Data:    err,
		})
	}

	if err := config.DB.Where("id = ?", updateRequest.Id).Omit("email", "password").Updates(usermodel.User{Name: updateRequest.Name, Address: updateRequest.Address,
		Contact: updateRequest.Contact}).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, basemodel.Response{
			Status:  false,
			Message: "Failed to update user data",
			Data:    err,
		})
	}

	return e.JSON(http.StatusOK, basemodel.Response{
		Status:  true,
		Message: "User data updated successfully",
		Data:    updateRequest,
	})
}
