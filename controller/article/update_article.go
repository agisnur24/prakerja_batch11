package articlectrl

import (
	"net/http"
	"prakerja_batch11/config"
	articlemodel "prakerja_batch11/model/article"
	basemodel "prakerja_batch11/model/base"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateArticle(e echo.Context) error {
	var request articlemodel.UpdateRequest
	e.Bind(&request)

	id := e.Param("id")

	var article articlemodel.Article
	if err := config.DB.First(&article, id).Error; err != nil {
		return e.JSON(http.StatusNotFound, basemodel.Response{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", id).Omit("user_id").Updates(articlemodel.Article{Title: request.Title, Description: request.Description,
		Image: request.Image}).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, basemodel.Response{
			Status:  false,
			Message: "Failed to update user data",
			Data:    err,
		})
	}

	return e.JSON(http.StatusOK, basemodel.Response{
		Status:  true,
		Message: "Article updated successfully",
		Data:    article,
	})
}
