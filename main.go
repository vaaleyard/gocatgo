package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/throttled/throttled/v2"
	"github.com/throttled/throttled/v2/store/memstore"
	"github.com/vaaleyard/gocatgo/gocatgo"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// rate limit settings
	store, err := memstore.New(65536)
	if err != nil {
		panic(err)
	}
	quota := throttled.RateQuota{
		MaxRate:  throttled.PerMin(20),
		MaxBurst: 5,
	}
	rateLimiterm, err := throttled.NewGCRARateLimiter(store, quota)
	if err != nil {
		panic(err)
	}
	httpRateLimiter := throttled.HTTPRateLimiter{
		RateLimiter: rateLimiterm,
		VaryBy:      &throttled.VaryBy{Path: true},
	}

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

	router := mux.NewRouter()
	router.HandleFunc("/", app.Upload).Methods("POST")
	router.HandleFunc("/", app.Home).Methods("GET")
	router.HandleFunc("/sha256", app.Sha256).Methods("GET")
	router.HandleFunc("/{shortid:[A-Za-z0-9]+(?:\\..*)?}", app.Fetch).Methods("GET")

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      httpRateLimiter.RateLimit(router), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	log.Printf("Listening on %s\n", server.Addr)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
