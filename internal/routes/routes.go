package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go-rest-api/internal/controller"
)

func Attach(e *echo.Echo) {
	e.GET("/", helloWorld)
	e.GET("/users", controller.NewUsersController().GetUsers)
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
