package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Use(app.recoverPanic)
	r.Use(app.logRequest)
	r.Use(secureHeaders)

	r.With(app.session.Enable).Get("/", app.home)
	r.Route("/post", func(r chi.Router) {
		r.Use(app.session.Enable)
		r.Get("/create", app.createArticleForm)
		r.Post("/create", app.createArticle)
		r.Get("/{id}", app.showArticle)
	})

	r.Route("/author", func(r chi.Router) {
		r.Use(app.session.Enable)
		r.Get("/signup", app.signupAuthorForm)
		r.Post("/signup", app.signupAuthor)
		r.Get("/login", app.loginAuthorForm)
		r.Post("/login", app.loginAuthor)
		r.Post("/logout", app.logoutAuthor)
	})

	fileServer := http.FileServer(http.Dir("./ui/static"))

	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix("/static", fileServer)
		fs.ServeHTTP(w, r)
	})

	return r
}
