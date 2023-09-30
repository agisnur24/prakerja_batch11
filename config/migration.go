package config

import (
	"prakerja_batch11/model/article"
	usermodel "prakerja_batch11/model/user"
)

func initMigrate() {
	DB.AutoMigrate(&usermodel.User{}, &article.Article{})
}
