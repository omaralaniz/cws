package main

import (
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/omaralaniz/computerandwebstuff/pkg/models/postgres"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	session       *sessions.Session
	posts         *postgres.PostModel
	authors       *postgres.AuthorModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	uri := flag.String("uri", "test", "Postgres URI")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*uri)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := newTemplateData("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		session:       session,
		posts:         &postgres.PostModel{DB: db},
		authors:       &postgres.AuthorModel{DB: db},
		templateCache: templateCache,
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting on server %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(uri string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}
