package server

import (
	"got/internal/utils"
	"got/internal/views/home"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/assets", "assets")

	e.GET("/", s.HomeHandler)

	return e
}

func (s *Server) HomeHandler(c echo.Context) error {
	return utils.Render(c, home.Home())
}
