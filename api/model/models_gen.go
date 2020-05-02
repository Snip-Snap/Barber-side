// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Barber struct {
	BarberID    string  `json:"barberID"`
	ShopID      int     `json:"shopID"`
	UserName    string  `json:"userName"`
	Password    string  `json:"password"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	PhoneNumber string  `json:"PhoneNumber"`
	Gender      *string `json:"gender"`
	Dob         string  `json:"dob"`
	HireDate    string  `json:"hireDate"`
	DismissDate *string `json:"dismissDate"`
	SeatNum     int     `json:"seatNum"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewBarber struct {
	ShopID      int     `json:"shopID"`
	UserName    string  `json:"userName"`
	Password    string  `json:"password"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	PhoneNumber string  `json:"PhoneNumber"`
	Gender      *string `json:"gender"`
	Dob         string  `json:"dob"`
	HireDate    string  `json:"hireDate"`
	DismissDate *string `json:"dismissDate"`
	SeatNum     int     `json:"seatNum"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Response struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}
