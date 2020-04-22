package barber

import (
	"api/internal/database"
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Barber represents a barber in a barbershop.
type Barber struct {
	BarberID    string
	ShopID      int
	UserName    string
	Password    string
	FirstName   string
	LastName    string
	PhoneNumber string
	Gender      *string
	Dob         string
	HireDate    string
	DismissDate *string
	SeatNum     int
}

// SaveOne inserts a specified new barber into the DB.
func (barber Barber) SaveOne() {
	insertBarber := "insert into barber (shopid, userName, hashedpassword," +
		"firstName, lastName, phonenumber, dob, gender, hiredate," +
		"dismissdate, seatnum) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
	stmt, err := database.Db.Prepare(insertBarber)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	hashpw, err := bcrypt.GenerateFromPassword([]byte(barber.Password),
		bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(barber.ShopID, barber.UserName, string(hashpw),
		barber.FirstName, barber.LastName, barber.PhoneNumber, barber.Dob,
		barber.Gender, barber.HireDate, barber.DismissDate, barber.SeatNum)
	if err != nil {
		log.Fatal(err)
	}

}

// GetAll selects all barbers from DB and returns them to resolver.
func GetAll() []Barber {
	getAllBarbers := "select * from barber"
	stmt, err := database.Db.Prepare(getAllBarbers)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var barbers []Barber
	for rows.Next() {
		var barber Barber
		// Save directly into arguments of Scan
		err := rows.Scan(&barber.BarberID, &barber.ShopID, &barber.UserName,
			&barber.Password, &barber.FirstName, &barber.LastName,
			&barber.PhoneNumber, &barber.Gender, &barber.Dob, &barber.HireDate,
			&barber.DismissDate, &barber.SeatNum)
		if err != nil {
			log.Fatal(err)
		}
		barbers = append(barbers, barber)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return barbers

}

// Get selects a specified barber via its ID and modifies the param barber.
func (barber *Barber) Get() {
	selectBarber := "select * from barber where barberid = $1"

	row := database.Db.QueryRow(selectBarber, barber.BarberID)

	// var newBarber Barber
	err := row.Scan(&barber.BarberID, &barber.ShopID, &barber.UserName,
		&barber.Password, &barber.FirstName, &barber.LastName,
		&barber.PhoneNumber, &barber.Gender, &barber.Dob,
		&barber.HireDate, &barber.DismissDate, &barber.SeatNum)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned.")
	case nil:
		fmt.Println(barber.BarberID, barber.FirstName)
	default:
		panic(err)
	}
}
