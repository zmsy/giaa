/*
 * Main server loop.
 */

package main

import (
	"net/http"

	_ "github.com/lib/pq"

	muxHandlers "github.com/gorilla/handlers"
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
	http.ListenAndServe(
		":8000",
		muxHandlers.CORS(
			muxHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			muxHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			muxHandlers.AllowedOrigins([]string{"*"}))(router))
}
