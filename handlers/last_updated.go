/*
Handler functions for specific routes on the API.
*/

package handlers

import (
	"fmt"
	"net/http"

	"github.com/zmsy/giaa/db"
	"github.com/zmsy/giaa/responses"
)

// LastUpdatedHandler simply returns the last updated timestamp
func LastUpdatedHandler(w http.ResponseWriter, r *http.Request) {

	// get the lastUpdatedTime from the database
	updt, err := queryLastUpdated()
	if err != nil {
		errorOut, _ := responses.ErrorResponse(err.Error())
		http.Error(w, errorOut, 500)
		return
	}

	// coerce the elements into JSON
	out, err := responses.SuccessResponse(updt)
	if err != nil {
		errorOut, _ := responses.ErrorResponse(err.Error())
		http.Error(w, errorOut, 500)
		return
	}

	fmt.Fprintf(w, out)
}

func queryLastUpdated() (db.LastUpdated, error) {
	lastUpdated := db.LastUpdated{}
	rows, err := db.Conn.Query(`
		SELECT
			id,
			updated_time
		FROM fantasy.etl_updates
		`)
	if err != nil {
		return lastUpdated, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&lastUpdated.ID,
			&lastUpdated.UpdateTime,
		)
		if err != nil {
			return lastUpdated, err
		}
	}
	err = rows.Err()
	return lastUpdated, err
}
