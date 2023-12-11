package gocatgo

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"path"

	"github.com/aidarkhanov/nanoid"
	"github.com/gorilla/mux"
	"github.com/vaaleyard/gocatgo/models"
)

func (app *App) Upload(w http.ResponseWriter, r *http.Request) {
	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileheader, err := r.FormFile("file")
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
		shortid = shortid + path.Ext(fileheader.Filename)

		result := models.Pastebin{}
		result.GetShortID(app.DB, shortid)

		// shortid does not exist
		if result.ShortID == "" {
			break
		}
	}

	model := models.Pastebin{File: buf.Bytes(), ShortID: shortid}
	model.New(app.DB)
	// I don't think there is a better way to do this
	var Scheme string
	if r.Header.Get("X-Forwarded-For") == "" {
		Scheme = "http"
	} else {
		Scheme = r.Header.Get("X-Forwarded-Proto")
	}
	fmt.Fprintf(w, "%s://%s/%s\n", Scheme, r.Host, model.ShortID)
}

func (app *App) Fetch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	paste := models.Pastebin{ShortID: vars["shortid"]}
	paste.Get(app.DB)

	fmt.Fprintf(w, "%v", string(paste.File))
}

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	// I don't think there is a better way to do this
	var Scheme string
	if r.Header.Get("X-Forwarded-For") == "" {
		Scheme = "http"
	} else {
		Scheme = r.Header.Get("X-Forwarded-Proto")
	}
	home := fmt.Sprintf(
		`
   gocatgo: another cool pastebin.

   * Usage:
     # Manually
       $ cat file.txt | curl -F "file=@-" %[1]s
         %[1]s/Rit
     # Or
       $ curl -F "file=@file.txt" %[1]s
     # Passing any string
       $ echo "some cool code" | curl -F "file=@-" %[1]s

   * Alias:
     # Run
       $ echo "$(curl %[1]s/alias)" >> ~/.bashrc
     # Use
       $ cat file.txt | gcg

   * GoCatGo is open source, you check it here:
        https://github.com/vaaleyard/gocatgo/
`,
        Scheme+"://"+r.Host)

	fmt.Fprintf(w, "%s", home)
}

func (app *App) Sha256(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%x", app.GetSha256())
}
