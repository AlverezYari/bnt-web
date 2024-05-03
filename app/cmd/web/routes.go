package main

import (
	"net/http"
	"strings"
)

type customFileServer struct {
	fileServer http.Handler
}

func (cfs customFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, ".css") {
		w.Header().Set("Content-Type", "text/css")
	}
	if strings.Contains(r.URL.Path, "tailwind.css") {
		w.Header().Set("Content-Type", "text/css")
	}
	cfs.fileServer.ServeHTTP(w, r)
}

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/html/static/"))
	customHandler := customFileServer{fileServer: http.StripPrefix("/static/", fileServer)}

	mux.Handle("/static/", customHandler)

	mux.HandleFunc("/", app.home)
	// More routes should be added here

	return mux
}
