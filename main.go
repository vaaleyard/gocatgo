package main

import (
	"log"
	"net/http"
	"os"

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

	log.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
