/*
Package db datatypes

These are representations of the rows in the database that are
going to be served over the API. Each of these roughly corresponds
to a single database table, although there may be some that
have joins built into them, etc.
*/
package db

import (
	"database/sql"
	"time"
)

// Player represents a single batter in the players database.
type Player struct {
	ID             int             `json:"id" db:"id"`
	Fullname       sql.NullString  `json:"fullName" db:"fullname"`
	Position       sql.NullString  `json:"position" db:"position"`
	Team           sql.NullInt64   `json:"team" db:"proteamid"`
	PercentOwned   sql.NullFloat64 `json:"percentOwned" db:"percentowned"`
	PercentStarted sql.NullFloat64 `json:"percentStarted" db:"percentstarted"`
}

// LastUpdated represents the time that the ETL was last run.
type LastUpdated struct {
	ID         int       `json:"id" db:"id"`
	UpdateTime time.Time `json:"updateTime" db:"updated_time"`
}

// RosterEntry represents a single player on a team.
type RosterEntry struct {
	ID              int            `json:"id" db:"id"`
	TeamID          sql.NullString `json:"teamId" db:"team_espn_id"`
	FullName        sql.NullString `json:"fullName" db:"fullname"`
	AcquisitionDate time.Time      `json:"acquisitionDate" db:"acquisitiondate"`
	AcquisitionType sql.NullString `json:"acquisitionType" db:"acquisitiontype"`
	Active          sql.NullBool   `json:"active" db:"active"`
	Droppable       sql.NullBool   `json:"droppable" db:"droppable"`
	InjuryStatus    sql.NullString `json:"injuryStatus" db:"injurystatus"`
}

// Team represents a fantasy team in the league.
type Team struct {
	ID           int             `json:"id" db:"id"`
	EspnID       sql.NullInt64   `json:"espnId" db:"espn_id"`
	Name         sql.NullString  `json:"name" db:"name"`
	Abbrev       sql.NullString  `json:"abbrev" db:"abbrev"`
	DivisionID   sql.NullString  `json:"divisionId" db:"divisionid"`
	Logo         sql.NullString  `json:"logo" db:"logo"`
	GamesBack    sql.NullFloat64 `json:"gamesBack" db:"gamesback"`
	Wins         sql.NullInt64   `json:"wins" db:"wins"`
	Losses       sql.NullInt64   `json:"losses" db:"losses"`
	Ties         sql.NullInt64   `json:"ties" db:"ties"`
	Acquisitions sql.NullInt64   `json:"acquisitions" db:"acquisitions"`
	Drops        sql.NullInt64   `json:"drops" db:"drops"`
	Trades       sql.NullInt64   `json:"trades" db:"trades"`
	MoveToActive sql.NullInt64   `json:"moveToActive" db:"movetoactive"`
	MoveToIR     sql.NullInt64   `json:"moveToIr" db:"movetoir"`
	FirstName    sql.NullString  `json:"firstName" db:"firstname"`
	LastName     sql.NullString  `json:"lastName" db:"lastname"`
}
