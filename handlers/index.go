/*
Handler functions for specific routes on the API.
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zmsy/giaa/db"
)

// IndexHandler simply returns the last updated timestamp
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// get the lastUpdatedTime from the database
	updt, err := queryLastUpdated()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// coerce the elements into JSON
	out, err := json.Marshal(updt)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// write the output back to the http response
	fmt.Fprintf(w, string(out))
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
			&lastUpdated.Id,
			&lastUpdated.UpdateTime,
		)
		if err != nil {
			return lastUpdated, err
		}
	}
	err = rows.Err()
	return lastUpdated, err
}
