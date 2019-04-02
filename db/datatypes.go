package db

// Player represents a single batter in the players database.
type Player struct {
	id             int
	fullname       string
	position       string
	team           string
	percentowned   float32
	percentstarted float32
}
