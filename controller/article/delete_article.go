package articlectrl

import (
	"net/http"
	"prakerja_batch11/config"
	"prakerja_batch11/model/article"
	basemodel "prakerja_batch11/model/base"

	"github.com/labstack/echo/v4"
)

func DeleteArticle(e echo.Context) error {
	id := e.Param("id")

	var article article.Article
	if err := config.DB.First(&article, id).Error; err != nil {
		return e.JSON(http.StatusNotFound, basemodel.Response{
			Status:  false,
			Message: "Article not found",
			Data:    err,
		})
	}

	config.DB.Delete(&article)
	return e.JSON(http.StatusOK, basemodel.Response{
		Status:  true,
		Message: "Article has been successfully deleted",
		Data:    nil,
	})
}
