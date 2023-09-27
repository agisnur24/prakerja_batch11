package config

import "os/user"

func initMigrate() {
	DB.AutoMigrate(&user.User{})
}
