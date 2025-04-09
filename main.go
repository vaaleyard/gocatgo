package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/vaaleyard/gocatgo/gocatgo"
)

func main() {
	app := gocatgo.App{
		Host:           "gcg.sh",
		BinaryFilename: "gocatgo.bin",
		Database: gocatgo.Database{
			Host:     os.Getenv("DBHOST"),
			User:     os.Getenv("DBUSER"),
			Password: os.Getenv("DBPASSWORD"),
			Name:     os.Getenv("DBNAME"),
			Port:     os.Getenv("DBPORT"),
		},
	}
	if err := app.Run(); err != nil {
		panic(err)
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /", app.Home)
	router.HandleFunc("POST /", app.Upload)
	router.HandleFunc("GET /sha256", app.Sha256)
	router.HandleFunc("GET /{shortid}", app.Fetch)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Channel to listen for interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Run server in goroutine so it doesn't block
	go func() {
		log.Println("Server listening on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %v", err)
		}
	}()

	// Block until signal is received
	<-stop
	log.Println("Shutting down server...")

	// Context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	os.Exit(0)
}
