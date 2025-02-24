package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {

	// Get port or default to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create routes with CORS. Then start server.
	r := createRoutes()
	c := setupCors()

	// Start the server
	logger.Info("starting server", "port", port)
	http.ListenAndServe(":"+port, c.Handler(r))
}
