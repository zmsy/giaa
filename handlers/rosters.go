/*
Roster information either all or by specific id.
*/

package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zmsy/giaa/db"
	"github.com/zmsy/giaa/responses"
)

// RostersHandler fetches all of the current active ESPN rosters.
func RostersHandler(w http.ResponseWriter, r *http.Request) {

	// get the lastUpdatedTime from the database
	rosters, err := queryRostersAll()
	if err != nil {
		errorOut, _ := responses.ErrorResponse(err.Error())
		http.Error(w, errorOut, 500)
		return
	}

	// coerce the elements into JSON
	// out, err := json.Marshal(rosters)
	out, err := responses.SuccessResponse(rosters)
	if err != nil {
		errorOut, _ := responses.ErrorResponse(err.Error())
		http.Error(w, errorOut, 500)
		return
	}

	fmt.Fprintf(w, out)
}

// RosterHandler fetches all of the current active ESPN rosters.
func RosterHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	teamID := vars["teamId"]
	teamInt, err := strconv.Atoi(teamID)

	if err != nil {
		errorOut, _ := responses.ErrorResponse(
			fmt.Sprintf("Not a valid team identifier: %s", teamID))
		http.Error(w, errorOut, 500)
		return
	}

	// get the lastUpdatedTime from the database
	rosters, err := queryRostersByTeamID(teamInt)
	if err != nil {
		errorOut, _ := responses.ErrorResponse(err.Error())
		http.Error(w, errorOut, 500)
		return
	}

	// coerce the elements into JSON
	out, err := responses.SuccessResponse(rosters)
	if err != nil {
		errorOut, _ := responses.ErrorResponse(err.Error())
		http.Error(w, errorOut, 500)
		return
	}

	fmt.Fprintf(w, out)
}

// queryRostersAll gets the entirety of the roster list
// from the database.
func queryRostersAll() ([]db.RosterEntry, error) {
	rows, err := db.Conn.Query(`
		select id, team_espn_id, fullname, acquisitiondate,
			acquisitiontype, active, droppable, injurystatus
		from fantasy.rosters
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return parseRowsToRosterEntries(rows)
}

func queryRostersByTeamID(teamID int) ([]db.RosterEntry, error) {
	sqlStatement := `
		select id, team_espn_id, fullname, acquisitiondate, acquisitiontype,
			active, droppable, injurystatus from fantasy.rosters
		where team_espn_id = $1`

	rows, err := db.Conn.Query(sqlStatement, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return parseRowsToRosterEntries(rows)
}

func parseRowsToRosterEntries(rows *sql.Rows) ([]db.RosterEntry, error) {
	rosterEntries := []db.RosterEntry{}
	for rows.Next() {
		rosterEntry := db.RosterEntry{}
		err := rows.Scan(
			&rosterEntry.ID,
			&rosterEntry.TeamID,
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
	err := rows.Err()
	return rosterEntries, err
}
