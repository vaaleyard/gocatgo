package gocatgo

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/aidarkhanov/nanoid"
	"github.com/gorilla/mux"
	"github.com/vaaleyard/gocatgo/models"
)

func (app *App) Upload(w http.ResponseWriter, r *http.Request) {
	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	_, _ = io.Copy(buf, file)

	/*
	   TODO:
	   1. better new id handling
	   2. check if it already exists
	   3. create ids with size+1 when all ids with size is over (how ?)
	*/
	shortid, err := nanoid.Generate(app.Alphabet, 3)
	if err != nil {
		panic(err)
	}

	/*
		result := models.Pastebin{}
		result.GetShortID(app.DB, shortid)
		if result.ShortID == "" {
			log.Println("This ID is new, going with this one...")
		} else {
			panic("this id already exists, generating another one...")
		}
	*/

	model := models.Pastebin{File: string(buf.Bytes()), ShortID: shortid}
	model.New(app.DB)

	fmt.Fprintf(w, "http://%s/%s\n", app.Host, model.ShortID)
}

func (app *App) Fetch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	paste := models.Pastebin{ShortID: vars["shortid"]}
	paste.Get(app.DB)

	fmt.Fprintf(w, "%v", paste.File)
}

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	home :=
		`
    gocatgo: another cool pastebin.

    * Usage:
      # cat file.txt | curl -F "file=@-" gocatgo.sh
`

	fmt.Fprintf(w, "%s", home)
}
