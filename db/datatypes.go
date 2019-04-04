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
	id             int
	fullname       string
	position       string
	team           string
	percentowned   float32
	percentstarted float32
}

// LastUpdated represents the time that the ETL was last run.
type LastUpdated struct {
	ID         int
	UpdateTime time.Time
}

// RosterEntry represents a single player on a team.
type RosterEntry struct {
	ID              int
	EspnID          int
	FullName        string
	AcquisitionDate time.Time
	AcquisitionType string
	Active          bool
	Droppable       bool
	InjuryStatus    string
}
