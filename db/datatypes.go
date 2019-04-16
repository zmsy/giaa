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
	ID             int             `json:"id"`
	Fullname       sql.NullString  `json:"fullName"`
	Position       sql.NullString  `json:"position"`
	Team           sql.NullString  `json:"team"`
	PercentOwned   sql.NullFloat64 `json:"percentOwned"`
	PercentStarted sql.NullFloat64 `json:"percentStarted"`
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
	ID           int             `json:"id"`
	EspnID       sql.NullInt64   `json:"espnId"`
	Name         sql.NullString  `json:"name"`
	Abbrev       sql.NullString  `json:"abbrev"`
	DivisionID   sql.NullString  `json:"divisionId"`
	Logo         sql.NullString  `json:"logo"`
	GamesBack    sql.NullFloat64 `json:"gamesBack"`
	Wins         sql.NullInt64   `json:"wins"`
	Losses       sql.NullInt64   `json:"losses"`
	Ties         sql.NullInt64   `json:"ties"`
	Acquisitions sql.NullInt64   `json:"acquisitions"`
	Drops        sql.NullInt64   `json:"drops"`
	Trades       sql.NullInt64   `json:"trades"`
	MoveToActive sql.NullInt64   `json:"moveToActive"`
	MoveToIR     sql.NullInt64   `json:"moveToIr"`
	FirstName    sql.NullString  `json:"firstName"`
	LastName     sql.NullString  `json:"lastName"`
}
