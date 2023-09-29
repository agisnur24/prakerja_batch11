package config

import usermodel "prakerja_batch11/model/user"

func initMigrate() {
	DB.AutoMigrate(&usermodel.User{})
}
