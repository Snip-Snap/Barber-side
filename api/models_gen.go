// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphqltest

type Barber struct {
	BarberID    string `json:"barberID"`
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	FullName    string `json:"fullName"`
	Gender      *bool  `json:"gender"`
	PhoneNumber string `json:"PhoneNumber"`
}

type Client struct {
	ClientID    string `json:"clientID"`
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	FullName    string `json:"fullName"`
	Gender      *bool  `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
}

type NewBarber struct {
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	FullName    string `json:"fullName"`
	Gender      bool   `json:"gender"`
	PhoneNumber string `json:"PhoneNumber"`
}

type NewClient struct {
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	FullName    string `json:"fullName"`
	Gender      *bool  `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
}

type Response struct {
	Error string `json:"error"`
}
