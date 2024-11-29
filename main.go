package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var PORT = flag.String("port", "80", "Specifies the port to the server (i.e. '80').")
var STATIC_PATH = flag.String("static", "/var/www/pocheteverso-static/", "Specifies the directory containing the static files.")
var ASSETS_PATH = flag.String("assets", "/var/www/pocheteverso-assets/", "Specifies the directory containing the asset files.")
var VERSION_FLAG = flag.Bool("version", false, "Displays the version of the server.")

const VERSION = "0.1"

func main() {
	flag.Parse()

	if *VERSION_FLAG {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	address := fmt.Sprintf(":%s", *PORT)

	server_mux := mux.NewRouter().StrictSlash(true)

	// Routes
	SetupRoutes(server_mux)

	server := http.Server {
		Addr: address,
		Handler: server_mux,
	}

	log.Printf("Listening on %s\n", address)
	log.Println("Configuration:")
	log.Printf("\tStatic files location: %s\n", *STATIC_PATH)
	log.Printf("\tAsset files loation  : %s\n", *ASSETS_PATH)
	log.Fatal(server.ListenAndServe())
}