/*
 * Main server loop.
 */

package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/zmsy/giaa/db"
	"github.com/zmsy/giaa/handlers"
)

// our main function
func main() {
	db.Connect()
	router := mux.NewRouter()
	router.HandleFunc("/api/index", handlers.IndexHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}
