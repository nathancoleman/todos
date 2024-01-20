package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/nathancoleman/todos/internal/routes"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})))
	slog.Info("Serving on http://localhost:8080/todos")

	if err := http.ListenAndServe(":8080", routes.Router()); err != nil {
		slog.Warn("Stopped serving", "error", err.Error())
		os.Exit(1)
	}

	slog.Warn("Stopped serving")
}
