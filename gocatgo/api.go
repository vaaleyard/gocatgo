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

	var shortid string
	for {
		// TODO: create ids with size+1 when all ids with size is over (how?)
		shortid, err = nanoid.Generate(app.Alphabet, 3)
		if err != nil {
			panic(err)
		}

		result := models.Pastebin{}
		result.GetShortID(app.DB, shortid)

		// shortid does not exist
		if result.ShortID == "" {
			break
		}
	}

	model := models.Pastebin{File: string(buf.Bytes()), ShortID: shortid}
	model.New(app.DB)

	fmt.Fprintf(w, "https://%s/%s\n", app.Host, model.ShortID)
}

func (app *App) Fetch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	paste := models.Pastebin{ShortID: vars["shortid"]}
	paste.Get(app.DB)

	fmt.Fprintf(w, "%v", paste.File)
}

func (app *App) Home(w http.ResponseWriter, r *http.Request) {

	home := fmt.Sprintf(
		`
   gocatgo: another cool pastebin.

   * Usage:
     $ cat file.txt | curl -F "file=@-" %[1]s
       %[1]s/Rit

     # will output current binary sha256
     $ curl %[1]s/sha256
       %[2]x

   * Examples:
     # With a file
       $ cat file.txt | curl -F "file=@-" %[1]s
     # or
       $ curl -F "file=@file.txt" %[1]s
     # Passing any string
       $ echo "some cool code" | curl -F "file=@-" %[1]s

   * GoCatGo is open source, you check it here:
        https://github.com/vaaleyard/gocatgo/
   * Roadmap of future development is also available:
        https://github.com/vaaleyard/gocatgo/blob/main/CONTRIBUTING.md#todo
	`, app.Host, app.GetSha256())

	fmt.Fprintf(w, "%s", home)
}

func (app *App) Sha256(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%x", app.GetSha256())
}
