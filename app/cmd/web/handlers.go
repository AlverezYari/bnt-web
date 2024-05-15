package main

import (
	/* 	"fmt" */
	"html/template"
	"net/http"
	/* "strconv" */ // "web.bluenimbustech.com/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	/* 	page, err := app.pages.Get(1) */
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }

	app.pageView(w, r, "home")

}

func (app *application) dynamic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<p>This is dynamically loaded content!</p>`))
}

func (app *application) pageView(w http.ResponseWriter, r *http.Request, id string) {

	if id == "home" {
		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			"./ui/html/pages/home.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		data := &templateData{pageID: id}

		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.serverError(w, r, err)
		}
		return
	}

}
