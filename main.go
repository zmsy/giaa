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
	router.HandleFunc("/api/last_updated", handlers.LastUpdatedHandler)
	router.HandleFunc("/api/rosters", handlers.RostersHandler)
	router.HandleFunc("/api/rosters/{teamId}", handlers.RosterHandler)
	router.HandleFunc("/api/teams", handlers.TeamsHandler)
	router.HandleFunc("/api/players/{playerId}", handlers.PlayersHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}
