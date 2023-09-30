package articlectrl

import (
	"net/http"
	"prakerja_batch11/config"
	"prakerja_batch11/model/article"
	basemodel "prakerja_batch11/model/base"

	"github.com/labstack/echo/v4"
)

func FindArticle(e echo.Context) error {
	id := e.Param("id")
	e.Bind(&id)

	var article article.Article
	if err := config.DB.Where("id = ?", id).First(&article).Error; err != nil {
		return e.JSON(http.StatusNotFound, basemodel.Response{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, basemodel.Response{
		Status:  true,
		Message: "Article found",
		Data:    article,
	})
}
