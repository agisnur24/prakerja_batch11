package route

import (
	"os"
	authctrl "prakerja_batch11/controller/auth"
	userctrl "prakerja_batch11/controller/user"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.POST("/login", authctrl.LogInController)
	e.POST("/register", authctrl.RegisterController)

	eAuthUser := e.Group("")
	eAuthUser.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	eAuthUser.GET("/users", userctrl.GetUsers)
	eAuthUser.GET("/find_user", userctrl.FindUser)
	eAuthUser.PUT("/update_user/:id", userctrl.UpdateUser)
	eAuthUser.PUT("/delete_user/:id", userctrl.DeleteUser)
}
