package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	f, err := os.OpenFile(fileHandler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := bytes.NewBuffer(nil)
	_, _ = io.Copy(buf, file)

	// model := app.Model{File: string(buf.Bytes()), ShortID: "sxiv"}
	// db.Create(&model)

	// use buf

	fmt.Fprintf(w, "Upload successful")
}
