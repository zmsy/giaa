/*
Package db datatypes

These are representations of the rows in the database that are
going to be served over the API. Each of these roughly corresponds
to a single database table, although there may be some that
have joins built into them, etc.
*/
package db

import (
	"time"
)

// Player represents a single batter in the players database.
type Player struct {
	ID             *int     `json:"id" db:"id"`
	Fullname       *string  `json:"fullName" db:"fullname"`
	Position       *string  `json:"position" db:"position"`
	Team           *int     `json:"team" db:"proteamid"`
	PercentOwned   *float32 `json:"percentOwned" db:"percentowned"`
	PercentStarted *float64 `json:"percentStarted" db:"percentstarted"`
}

// LastUpdated represents the time that the ETL was last run.
type LastUpdated struct {
	ID         *int       `json:"id" db:"id"`
	UpdateTime *time.Time `json:"updateTime" db:"updated_time"`
}

// RosterEntry represents a single player on a team.
type RosterEntry struct {
	ID              *int       `json:"id" db:"id"`
	TeamID          *string    `json:"teamId" db:"team_espn_id"`
	FullName        *string    `json:"fullName" db:"fullname"`
	AcquisitionDate *time.Time `json:"acquisitionDate" db:"acquisitiondate"`
	AcquisitionType *string    `json:"acquisitionType" db:"acquisitiontype"`
	Active          *bool      `json:"active" db:"active"`
	Droppable       *bool      `json:"droppable" db:"droppable"`
	InjuryStatus    *string    `json:"injuryStatus" db:"injurystatus"`
}

// Team represents a fantasy team in the league.
type Team struct {
	ID           int      `json:"id" db:"id"`
	EspnID       *int     `json:"espnId" db:"espn_id"`
	Name         *string  `json:"name" db:"name"`
	Abbrev       *string  `json:"abbrev" db:"abbrev"`
	DivisionID   *string  `json:"divisionId" db:"divisionid"`
	Logo         *string  `json:"logo" db:"logo"`
	GamesBack    *float32 `json:"gamesBack" db:"gamesback"`
	Wins         *int     `json:"wins" db:"wins"`
	Losses       *int     `json:"losses" db:"losses"`
	Ties         *int     `json:"ties" db:"ties"`
	Acquisitions *int     `json:"acquisitions" db:"acquisitions"`
	Drops        *int     `json:"drops" db:"drops"`
	Trades       *int     `json:"trades" db:"trades"`
	MoveToActive *int     `json:"moveToActive" db:"movetoactive"`
	MoveToIR     *int     `json:"moveToIr" db:"movetoir"`
	FirstName    *string  `json:"firstName" db:"firstname"`
	LastName     *string  `json:"lastName" db:"lastname"`
}
