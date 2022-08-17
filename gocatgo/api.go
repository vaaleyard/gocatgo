package gocatgo

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
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

	encryptedFile, err := AESEncrypt(app.AESCipherkey, string(buf.Bytes()))
	if err != nil {
		panic(err)
	}

	model := models.Pastebin{File: string(encryptedFile), ShortID: shortid}
	model.New(app.DB)

	fmt.Fprintf(w, "http://%s/%s\n", app.Host, model.ShortID)
}

func (app *App) Fetch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	paste := models.Pastebin{ShortID: vars["shortid"]}
	paste.Get(app.DB)

	plainTextFile, err := AESDecrypt(app.AESCipherkey, paste.File)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%v", plainTextFile)
}

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	home := fmt.Sprintf(
		`
    gocatgo: another cool pastebin.

    * Usage:
      # cat file.txt | curl -F "file=@-" %s
`, app.Host)

	fmt.Fprintf(w, "%s", home)
}

func (app *App) Sha256(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(app.BinaryFilename)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%x", sha256.Sum256(data))
}
