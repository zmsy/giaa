/*
Query for specific players by ID.
*/

package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zmsy/giaa/db"
	"github.com/zmsy/giaa/responses"
)

// PlayersHandler looks up an individual player's information by
// their ID.
func PlayersHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	playerID := vars["playerId"]
	playerInt, err := strconv.Atoi(playerID)

	player, err := queryPlayerByID(&playerInt)
	if err != nil {
		errorOut, _ := responses.ErrorResponse(err.Error())
		http.Error(w, errorOut, 500)
		return
	}

	// coerce the elements into JSON
	out, err := responses.SuccessResponse(player)
	if err != nil {
		errorOut, _ := responses.ErrorResponse(err.Error())
		http.Error(w, errorOut, 500)
		return
	}

	fmt.Fprintf(w, out)
}

func queryPlayerByID(playerID *int) (db.Player, error) {
	sqlStatement := `
		select id, fullname, position, proteamid,
			percentowned, percentstarted
		from fantasy.players
		where id = $1`
	player := db.Player{}
	err := db.Conn.Get(&player, sqlStatement, *playerID)
	return player, err
}
