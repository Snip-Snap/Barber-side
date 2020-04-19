package database

import (
	"bufio"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

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

func ConnectPSQL() {

	// creds := parseCreds("/run/secrets")
	//code when you run manually without docker
	creds := parseCreds("../../dbcreds.config")

	var err error

	Db, err = sql.Open("postgres", creds)
	print("connect psql")
	if err != nil {
		log.Fatal(err)
	}
}

func ClosePSQL() {
	Db.Close()
}
