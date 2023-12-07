package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

func init(){
	var logHandler *slog.JSONHandler
	env := os.Getenv("APP_ENV")
	if env == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	}
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

