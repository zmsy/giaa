/*
 * Main server loop.
 */

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zmsy/giaa/handlers"
)

// our main function
func main() {
	router := mux.NewRouter()
	http.HandleFunc("/api/index", handlers.IndexHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}
