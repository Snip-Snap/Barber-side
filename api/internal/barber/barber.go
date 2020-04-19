package barber

import (
	"graphqltest/api/internal/database"
	"log"

	"golang.org/x/crypto/bcrypt"
)

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

func (barber Barber) SaveOne() {
	insertBarber := "insert into barber (shopid, userName, hashedpassword,firstName, lastName, phonenumber, dob, gender, hiredate, dismissdate, seatnum) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
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
