package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	dynmaicMiddleware := []func(http.Handler) http.Handler{app.session.Enable, noSurf, app.authenticate}
	r.Use(app.recoverPanic)
	r.Use(app.logRequest)
	r.Use(secureHeaders)

	r.With(dynmaicMiddleware...).Get("/", app.home)
	r.With(dynmaicMiddleware...).Get("/about", app.about)
	r.With(dynmaicMiddleware...).Get("/contact", app.contact)
	r.With(dynmaicMiddleware...).Get("/{category}", app.category)
	r.Route("/post", func(r chi.Router) {
		r.Use(dynmaicMiddleware...)
		r.With(app.authenticationRequired).Get("/create", app.createArticleForm)
		r.With(app.authenticationRequired).Post("/create", app.createArticle)
		r.With(app.authenticationRequired).Get("/edit/{id}", app.editArticleForm)
		r.With(app.authenticationRequired).Post("/edit/{id}", app.editArticle)
		r.Get("/{id}", app.showArticle)
	})

	r.Route("/author", func(r chi.Router) {
		r.Use(dynmaicMiddleware...)
		r.Get("/signup", app.signupAuthorForm)
		r.Post("/signup", app.signupAuthor)
		r.Get("/login", app.loginAuthorForm)
		r.Post("/login", app.loginAuthor)
		r.With(app.authenticationRequired).Post("/logout", app.logoutAuthor)
	})

	fileServer := http.FileServer(http.Dir("./ui/static"))

	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix("/static", fileServer)
		fs.ServeHTTP(w, r)
	})

	return r
}
