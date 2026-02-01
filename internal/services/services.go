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
	AddNotes(ctx context.Context, title string, text string, timeStemp *time.Time) (int64, error)
}

type NoteRemover interface {
	RemoveNotes(ctx context.Context, note_id int64, user_id int64) error
}

func New(logger *slog.Logger, NotesSaver NotesSaver, NoteRemover NoteRemover) *Services {
	return &Services{Logger: logger, NotesSaver: NotesSaver, NoteRemover: NoteRemover}
}

func (s *Services) AddNotes(c echo.Context) error {

	notes := &domain.Notes{}

	r := c.Request()

	err := json.NewDecoder(r.Body).Decode(notes)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error parse data from json")
	}

	if err := c.Validate(notes); err != nil {
		return c.JSON(http.StatusBadRequest, "Bed credentinals")
	}

	timeLocation, err := time.LoadLocation(notes.Time_zone)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error parse timezone")
	}

	time := time.Now().In(timeLocation)

	_, err = s.NotesSaver.AddNotes(c.Request().Context(), notes.Title, notes.Text, &time)

	if err != nil {
		s.Logger.Error("Error add notes to database:", slog.Any("", err))
		return c.JSON(http.StatusInternalServerError, "Error paste data to database")
	}

	return nil
}

func (s *Services) RemoveNotes(c echo.Context) error {

	id := c.Param("id")

	


	return nil
}
