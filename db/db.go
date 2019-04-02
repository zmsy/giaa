package db

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

var conn *sql.DB

// Connect creates a connection to postgres and stores in db.conn.
func Connect() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	conn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		panic("POSTGRES_HOST environment variable required but not set")
	}
	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		panic("POSTGRES_PORT environment variable required but not set")
	}
	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		panic("POSTGRES_USER environment variable required but not set")
	}
	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		panic("POSTGRES_PASSWORD environment variable required but not set")
	}
	db, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		panic("POSTGRES_DB environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = db
	return conf
}
