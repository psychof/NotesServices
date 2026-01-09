package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/psychof/NotesServices/internal/services"
)

func Handlers() *echo.Echo {
	r := echo.New()

	r.Validator = echo.New().Validator

	r.POST("/notes", ) 

	r.GET("/notes/:id",)
	

	return r
}
