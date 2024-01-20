package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/nathancoleman/todos/internal"
)

func main() {
	if err := http.ListenAndServe(":8080", internal.Router()); err != nil {
		slog.Warn("Stopped serving", "error", err.Error())
		os.Exit(1)
	}

	slog.Warn("Stopped serving")
}
