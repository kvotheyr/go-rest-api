package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go-rest-api/internal/repository"
)

type Controller interface {
	GetUsers() error
}

type UsersController struct {
	usersRepository repository.UsersRepository
}

func NewUsersController() *UsersController {
	return &UsersController{
		usersRepository: repository.NewRepository(),
	}
}

func (u UsersController) GetUsers(c echo.Context) error {
	users, err := u.usersRepository.GetUsers(c.Request().Context())
	if err != nil {
		return errors.Wrap(err, "failed to get users")
	}

	return c.JSON(200, users)
}
