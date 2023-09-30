package articlectrl

import (
	"net/http"
	"prakerja_batch11/config"
	"prakerja_batch11/model/article"
	basemodel "prakerja_batch11/model/base"

	"github.com/labstack/echo/v4"
)

func CreateArticle(e echo.Context) error {
	var article article.Article
	e.Bind(&article)

	err := config.DB.Create(&article)
	if err.Error != nil {
		return e.JSON(http.StatusInternalServerError, basemodel.Response{
			Status:  false,
			Message: "Failed to post article",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusCreated, basemodel.Response{
		Status:  true,
		Message: "Successfully posted the article",
		Data:    article,
	})
}
