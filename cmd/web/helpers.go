package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"

	"github.com/russross/blackfriday/v2"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) addDefault(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	td.Flash = app.session.PopString(r, "flash")
	return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist!", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefault(td, r))
	if err != nil {
		app.serverError(w, err)
	}

	buf.WriteTo(w)
}

func (app *application) markdownToHtml(content []byte) template.HTML {
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{Flags: blackfriday.CompletePage})

	parsed := blackfriday.Run(content, blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(blackfriday.FencedCode))

	body := template.HTML(parsed)

	return body
}

func (app *application) isAuthenticated(r *http.Request) bool {
	return app.session.Exists(r, "authenticatedUserID")
}
