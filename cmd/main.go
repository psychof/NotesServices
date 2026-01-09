package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/psychof/NotesServices/internal/config"
	handler "github.com/psychof/NotesServices/internal/handlers"
	"github.com/psychof/NotesServices/internal/services"
	"github.com/psychof/NotesServices/internal/storage"
)

func main() {

	config := config.MustLoad()

	ctx := context.Background()

	fmt.Print(config)

	logger := SetupLoger(config.Env)

	logger.Info("Server start")

	storage, err := storage.New(ctx, config.Database.ConnString)

	if err != nil {
		logger.Error("Error setup database%s", err)
	}
	r := handler.Handlers()

	s := http.Server{
		WriteTimeout: config.Server.TimeOut,
		ReadTimeout:  config.Server.TimeOut,
		IdleTimeout:  config.Server.IdleTime,
		Handler:      r,
	}

	_ = services.New(logger, storage, storage)

	if err := s.ListenAndServe(); err != nil {
		logger.Error("Error starting server")
	}

	logger.Info("Server started")

	fmt.Print(config)

}

func SetupLoger(env string) *slog.Logger {

	var l *slog.Logger

	switch env {
	case "local":
		l = slog.New(slog.NewJSONHandler(os.Stdin, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case "dev":
		l = slog.New(slog.NewJSONHandler(os.Stdin, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case "prod":
		l = slog.New(slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return l
}
