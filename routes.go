package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	static_fs := http.FileServer(http.Dir(*STATIC_PATH))
	download_fs := http.FileServer(http.Dir(*ASSETS_PATH))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static_fs))
	r.PathPrefix("/download/").Handler(http.StripPrefix("/download/", download_fs))
}