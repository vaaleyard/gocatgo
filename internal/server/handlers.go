package server

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"path"
	"regexp"
	"time"

	"github.com/aidarkhanov/nanoid"
	"github.com/alexliesenfeld/health"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/vaaleyard/gocatgo/internal/repository"
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

	ctx := r.Context()
	queries := repository.New(app.DB)

	var fileID string
	for {
		// TODO: create ids with size+1 when all ids with size is over (how?)
		fileID, err = nanoid.Generate(app.Alphabet, 3)
		if err != nil {
			panic(err)
		}
		fileID = fileID + path.Ext(fileheader.Filename)

		paste, err := queries.GetPaste(ctx, fileID)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				slog.Warn("error connecting to database: ", pgErr.Message, pgErr.Code)
				return
			}
		}

		if paste.FileID == "" {
			break
		}
	}

	err = queries.CreatePaste(ctx, repository.CreatePasteParams{
		FileID:      fileID,
		FileContent: buf.Bytes(),
	})
	if err != nil {
		http.Error(w, "failed to create it in database, please try again: "+err.Error(), http.StatusBadRequest)
		return
	}
	slog.Info("File " + fileID + " has been created")

	// I don't think there is a better way to do this
	var Scheme string
	if r.Header.Get("X-Forwarded-For") == "" {
		Scheme = "http"
	} else {
		Scheme = r.Header.Get("X-Forwarded-Proto")
	}
	fmt.Fprintf(w, "%s://%s/%s\n", Scheme, r.Host, fileID)
}

func (app *App) Get(w http.ResponseWriter, r *http.Request) {
	fileID := r.PathValue("fileid")

	// block unusual paths
	re := regexp.MustCompile(`^[A-Za-z0-9]+(?:\.[A-Za-z0-9]+)?$`)
	if !re.MatchString(fileID) {
		http.Error(w, "Invalid file name", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	queries := repository.New(app.DB)
	paste, err := queries.GetPaste(ctx, fileID)
	if errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "failed to fetch URL, please contact the administrator: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v", string(paste.FileContent))
}

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
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

   * Function:
     # Run
       $ echo "$(curl %[1]s/function)" >> ~/.bashrc
     # Use
       $ gcg file.txt

   * You can check gcg uptime here: https://status.gcg.sh

   * GoCatGo is open source:
        https://github.com/vaaleyard/gocatgo/
`,
		r.Host)

	fmt.Fprintf(w, "%s", home)
}

func (app *App) Sha256(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%x", app.getBinarySha256())
}

func (app *App) Healthcheck() http.HandlerFunc {
	return health.NewHandler(
		health.NewChecker(
			health.WithCacheDuration(1*time.Second),
			health.WithTimeout(10*time.Second),

			health.WithCheck(health.Check{
				Name:    "database",
				Timeout: 2 * time.Second,
				Check:   app.DB.Ping,
			}),

			health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
				slog.Info(fmt.Sprintf("Health status changed to %s", state.Status))
			}),
		),
	)
}
