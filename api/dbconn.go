package api

import (
	"bufio"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func parseCreds(fn string) string {
	infile, err := os.Open(fn)
	CheckError(err)

	filescanner := bufio.NewScanner(infile)

	var creds string
	for filescanner.Scan() {
		creds = filescanner.Text()
	}
	infile.Close()
	return creds
}

func ConnectPSQL() {

	creds := parseCreds("/run/secrets")
	//code when you run manually without docker
	//	creds := parseCreds("../../dbcreds.config")

	var err error

	db, err = sql.Open("postgres", creds)
	print("connect psql")
	CheckError(err)
}

func ClosePSQL() {
	db.Close()
}
