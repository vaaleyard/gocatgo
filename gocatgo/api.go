package gocatgo

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
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

	// TODO: better new short id handling
	shortid, err := shortid.Generate()
	if err != nil {
		panic(err)
	}
	model := models.Pastebin{File: string(buf.Bytes()), ShortID: shortid}
	model.New(app.DB)

	fmt.Fprintf(w, "Upload successful")
}

func (app *App) Fetch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	paste := models.Pastebin{ShortID: vars["shortid"]}
	paste.Get(app.DB)

	fmt.Fprintf(w, "%v", paste.File)
}
