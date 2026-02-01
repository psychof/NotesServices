package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/psychof/NotesServices/internal/pkg"
	"github.com/psychof/NotesServices/internal/services"
)

type Handlers struct {
	Services *services.Services
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{Services: services}
}

func (s *Handlers) SetupRouter() *echo.Echo {

	r := echo.New()
	r.Validator = pkg.New()

	return s.SetupRoutes(r)

}

func(s *Handlers) SetupRoutes(r *echo.Echo) *echo.Echo {
	
	r.POST("/notes",s.Services.AddNotes)
	r.GET("/notes/:id",s.Services.RemoveNotes)
	
 
	return r 
}
