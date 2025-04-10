package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(app *App) error {
	router := http.NewServeMux()
	router.HandleFunc("GET /", app.Home)
	router.HandleFunc("POST /", app.Upload)
	router.HandleFunc("GET /sha256", app.Sha256)
	router.HandleFunc("GET /{fileid}", app.Get)

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// in goroutine so it doesn't block
	go func() {
		slog.Info("Server listening on " + server.Addr)
		if err := server.ListenAndServe(); err != nil {
			slog.Error("failed to start server: ", err)
		}
	}()

	// Channel to listen for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Block until signal is received
	<-quit

	// Context with timeout for graceful shutdown
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %v", err)
	}
	slog.Info("Server stopped")

	return nil
}
