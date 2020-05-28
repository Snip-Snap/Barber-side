package database

import (
	"bufio"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Db is used to access the db via other packages.
var Db *sql.DB

func parseCreds(fn string) string {
	infile, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	filescanner := bufio.NewScanner(infile)

	var creds string
	for filescanner.Scan() {
		creds = filescanner.Text()
	}
	infile.Close()
	return creds
}

// ConnectPSQL opens a connection with my psql database.
func ConnectPSQL() {

	creds := parseCreds("/run/secrets")
	//code when you run manually without docker
	// creds := parseCreds("./../../dbcreds.config")

	var err error

	Db, err = sql.Open("postgres", creds)
	print("connect psql")
	if err != nil {
		log.Fatal(err)
	}
}

// ClosePSQL closes connection with psql database.
func ClosePSQL() {
	Db.Close()
}
