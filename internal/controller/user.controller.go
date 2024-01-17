package controller

import (
	"got/internal/models"
	"got/internal/utils"
	userView "got/internal/views/user"

	"github.com/labstack/echo/v4"
)

func UserController(c echo.Context) error {
	user := models.User{
		Name:  "Benjamin",
		Email: "b@aol.com",
	}
	return utils.Render(c, userView.Show(user))
}
