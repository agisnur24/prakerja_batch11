package articlectrl

import (
	"net/http"
	"prakerja_batch11/config"
	"prakerja_batch11/model/article"
	basemodel "prakerja_batch11/model/base"

	"github.com/labstack/echo/v4"
)

func GetArticles(e echo.Context) error {
	var articles []article.Article

	result := config.DB.Find(&articles)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, basemodel.Response{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, basemodel.Response{
		Status:  true,
		Message: "Success",
		Data:    articles,
	})
}
