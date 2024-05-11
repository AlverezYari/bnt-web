package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"web.bluenimbustech.com/internal/models"
)

type application struct {
	logger *slog.Logger
	pages  *models.PageModel
}

func main() {
	// Define a new command-line flag for the network address
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Create a new logger instance
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// Parse the command-line flags
	flag.Parse()
	// Setup Application instance
	app := &application{
		logger: logger,
		pages:  &models.PageModel{},
		/* 	snippets: &models.SnippetModel{DB: db}, */
	}

	// Log the server is starting
	logger.Info("starting server", "addr", *addr)
	// Start the server & log any errors
	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}
