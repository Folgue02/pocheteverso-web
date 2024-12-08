package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func handleBackups(pConfig PvwConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(pConfig.DynResFilePath("backups.html"))
		if err != nil {
			panic(err)
		}

		backupsInfo, err := pConfig.GetBackupsInfo()
		if err != nil {
			log.Printf("Error, couldn't read backup info due to the following error: %s\n", err) 
			w.Write([]byte("Internal server error"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("Entries: %d, %v\n", len(backupsInfo), backupsInfo)

		ctx := map[string]any {
			"entries": backupsInfo,
		}

		tmpl.Execute(w, ctx)
	}
}


func SetupRoutes(r *mux.Router, pConfig PvwConfig) {
	static_fs := http.FileServer(http.Dir(*STATIC_PATH))
	download_fs := http.FileServer(http.Dir(*ASSETS_PATH))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static_fs))
	r.PathPrefix("/download/").Handler(http.StripPrefix("/download/", download_fs))
	r.HandleFunc("/backups", handleBackups(pConfig))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/", http.StatusMovedPermanently)
	})
}
