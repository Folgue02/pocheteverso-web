package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var PORT = flag.Int("port", 80, "Specifies the port to the server (i.e. '80').")
var STATIC_PATH = flag.String("static", "/var/www/pocheteverso-static/", "Specifies the directory containing the static files.")
var ASSETS_PATH = flag.String("assets", "/var/www/pocheteverso-assets/", "Specifies the directory containing the asset files.")
var VERSION_FLAG = flag.Bool("version", false, "Displays the version of the server.")
var DYNRES_PATH = flag.String("dynres", "/opt/pocheteverso/dynres", "Specifies the directory containing the dynamic resources.")
var SSL_CERT = flag.String("cert", "", "Path to the SSL certificate [must be used in conjunction with -key]")
var SSL_KEY = flag.String("key", "", "Path to the SSL key [must be used in conjunction with -cert]")

const VERSION = "0.2"

func main() {
	flag.Parse()

	if *VERSION_FLAG {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	address := fmt.Sprintf(":%d", *PORT)

	server_mux := mux.NewRouter().StrictSlash(true)

	serverConfig := PvwConfig {
		StaticPath: *STATIC_PATH,
		AssetsPath: *ASSETS_PATH,
		DynResPath: *DYNRES_PATH,
		Port: *PORT,
		SslCertPath: *SSL_CERT,
		SslKeyPath: *SSL_KEY,
	}

	// Routes
	SetupRoutes(server_mux, serverConfig)

	server := http.Server {
		Addr: address,
		Handler: server_mux,
	}

	log.Printf("Listening on %s\n", address)
	log.Println("Configuration:")
	log.Printf("\tStatic files location: %s\n", *STATIC_PATH)
	log.Printf("\tAsset files loation  : %s\n", *ASSETS_PATH)
	
	if !(strings.TrimSpace(*SSL_CERT) == "" && strings.TrimSpace(*SSL_KEY) == "") &&
		!(strings.TrimSpace(*SSL_CERT) != "" && strings.TrimSpace(*SSL_KEY) != "") {
		log.Fatal("-key and -cert must be used together, or not used at all (if http is desired instead of https)")
		os.Exit(1)
	}

	if (strings.TrimSpace(*SSL_CERT) != "") {
		log.Fatal(server.ListenAndServeTLS(*SSL_CERT, *SSL_KEY))
	} else {
		log.Fatal(server.ListenAndServe())
	}
}
