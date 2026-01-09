package services

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/psychof/NotesServices/internal/domain"
)

type Services struct {
	Logger      *slog.Logger
	NotesSaver  NotesSaver
	NoteRemover NoteRemover
}

type NotesSaver interface {
	AddNotes(ctx context.Context, title string, text string, timeStemp *time.Time) error
}

type NoteRemover interface {

}

func New(logger *slog.Logger, NotesSaver NotesSaver, NoteRemover NoteRemover) *Services {
	return &Services{Logger: logger, NotesSaver: NotesSaver, NoteRemover: NoteRemover}
}

func(s *Services) AddNotes(c echo.Context) error {

	notes := &domain.Notes{}

	r := c.Request()

	err := json.NewDecoder(r.Body).Decode(notes)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error parse data from json")
	}

	if err := c.Validate(notes); err != nil {
		return c.JSON(http.StatusBadGateway, "Bed credentinals")
	}

	err = s.NotesSaver.AddNotes(c.Request().Context(), notes.Title, notes.Text, notes.Time_stamp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error paste data to database")
	}

	return nil
}
