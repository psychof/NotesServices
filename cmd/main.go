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
		logger.Error("Error setup database:", slog.Any("", err))
	}

	services := services.New(logger, storage, storage)

	handlers := handler.NewHandlers(services)

	r := handlers.SetupRouter()

	s := http.Server{
		Addr:         config.Server.Addr,
		WriteTimeout: config.Server.TimeOut,
		ReadTimeout:  config.Server.TimeOut,
		IdleTimeout:  config.Server.IdleTime,
		Handler:      r,
	}

	if err := s.ListenAndServe(); err != nil {
		logger.Error("Error starting server")
	}
}

func SetupLoger(env string) *slog.Logger {

	var l *slog.Logger

	switch env {
	case "local":
		l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case "dev":
		l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case "prod":
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return l
}
