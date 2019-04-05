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
	ID             int     `json:"id"`
	Fullname       string  `json:"fullName"`
	Position       string  `json:"position"`
	Team           string  `json:"team"`
	PercentOwned   float32 `json:"percentOwned"`
	PercentStarted float32 `json:"percentStarted"`
}

// LastUpdated represents the time that the ETL was last run.
type LastUpdated struct {
	ID         int       `json:"id"`
	UpdateTime time.Time `json:"updateTime"`
}

// RosterEntry represents a single player on a team.
type RosterEntry struct {
	ID              int       `json:"id"`
	TeamID          int       `json:"teamId"`
	FullName        string    `json:"fullName"`
	AcquisitionDate time.Time `json:"acquisitionDate"`
	AcquisitionType string    `json:"acquisitionType"`
	Active          bool      `json:"active"`
	Droppable       bool      `json:"droppable"`
	InjuryStatus    string    `json:"injuryStatus"`
}
