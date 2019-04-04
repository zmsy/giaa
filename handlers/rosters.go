/*
Roster information either all or by specific id.
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zmsy/giaa/db"
)

// RostersHandler fetches all of the current active ESPN rosters.
func RostersHandler(w http.ResponseWriter, r *http.Request) {

	// get the lastUpdatedTime from the database
	rosters, err := queryRosters()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// coerce the elements into JSON
	out, err := json.Marshal(rosters)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// write the output back to the http response
	fmt.Fprintf(w, string(out))
}

func queryRosters() ([]db.RosterEntry, error) {
	rosterEntries := []db.RosterEntry{}
	rows, err := db.Conn.Query(`
		select
			id,
			team_espn_id,
			fullname,
			acquisitiondate,
			acquisitiontype,
			active,
			droppable,
			injurystatus
		from fantasy.rosters
	`)
	if err != nil {
		return rosterEntries, err
	}
	defer rows.Close()
	for rows.Next() {
		rosterEntry := db.RosterEntry{}
		err = rows.Scan(
			&rosterEntry.ID,
			&rosterEntry.EspnID,
			&rosterEntry.FullName,
			&rosterEntry.AcquisitionDate,
			&rosterEntry.AcquisitionType,
			&rosterEntry.Active,
			&rosterEntry.Droppable,
			&rosterEntry.InjuryStatus,
		)
		if err != nil {
			return rosterEntries, err
		}
		rosterEntries = append(rosterEntries, rosterEntry)
	}
	err = rows.Err()
	return rosterEntries, err
}
