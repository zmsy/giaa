package handlers

import (
	"fmt"
	"net/http"

	"github.com/zmsy/giaa/db"
	"github.com/zmsy/giaa/responses"
)

// TeamsHandler returns a list of teams in the fantasy league.
func TeamsHandler(w http.ResponseWriter, r *http.Request) {

	// get the lastUpdatedTime from the database
	teams, err := queryTeams()
	if err != nil {
		errorOut, _ := responses.ErrorResponse(err.Error())
		http.Error(w, errorOut, 500)
		return
	}

	// coerce the elements into JSON
	out, err := responses.SuccessResponse(teams)
	if err != nil {
		errorOut, _ := responses.ErrorResponse(err.Error())
		http.Error(w, errorOut, 500)
		return
	}

	fmt.Fprintf(w, out)
}

func queryTeams() ([]db.Team, error) {
	teams := []db.Team{}
	rows, err := db.Conn.Query(`
		select t.id, t.espn_id, t.name, t.abbrev, t.divisionid,
			t.logo, t.gamesback, t.wins, t.losses, t.ties,
			t.acquisitions, t.drops, t.trades, t.movetoactive,
			t.movetoir, lower(m.firstname) as firstname,
			lower(m.lastname) as lastname
		from fantasy.teams t
			join fantasy.members m
				on t.primaryowner = m.espn_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		team := db.Team{}
		err = rows.Scan(
			&team.ID,
			&team.EspnID,
			&team.Name,
			&team.Abbrev,
			&team.DivisionID,
			&team.Logo,
			&team.GamesBack,
			&team.Wins,
			&team.Losses,
			&team.Ties,
			&team.Acquisitions,
			&team.Drops,
			&team.Trades,
			&team.MoveToActive,
			&team.MoveToIR,
			&team.FirstName,
			&team.LastName,
		)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	err = rows.Err()
	return teams, err
}
