package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphqltest/api/generated"
	"graphqltest/api/internal/barber"
	"graphqltest/api/model"
)

func (r *mutationResolver) SignupClient(ctx context.Context, input model.NewClient) (*model.Response, error) {
	// print("here\n")
	// statement, err := Db.Prepare("insert into client (fullname, gender, phonenumber, username, hashedpassword) values($1, $2, $3, $4, $5)")
	// CheckError(err)
	// print("here\n")
	// //TODO: Check inputs for uniqueness and appropriate characters.
	// // pw := string(input.Password)
	// hash, err := bcrypt.GenerateFromPassword([]byte(input.Password),
	// 	bcrypt.DefaultCost)
	// CheckError(err)

	// hashedInputpw := string(hash)
	// _, err = statement.Exec(input.FullName,
	// 	input.Gender, input.PhoneNumber, input.UserName, hashedInputpw)
	// CheckError(err)

	res := &model.Response{Error: "Okay"}

	return res, nil
}

func (r *mutationResolver) SignUpBarber(ctx context.Context, input model.NewBarber) (*model.Response, error) {
	// Why doesn't barber := barber.Barber work? It can't infer its type?
	var barber barber.Barber
	barber.ShopID = input.ShopID
	barber.UserName = input.UserName
	barber.Password = input.Password
	barber.FirstName = input.FirstName
	barber.LastName = input.LastName
	barber.PhoneNumber = input.PhoneNumber
	barber.Gender = input.Gender
	barber.Dob = input.Dob
	barber.HireDate = input.HireDate
	barber.DismissDate = input.DismissDate
	barber.SeatNum = input.SeatNum

	barber.SaveOne()

	res := &model.Response{Error: "Error message"}

	return res, nil
}

func (r *queryResolver) GetAllBarbers(ctx context.Context) ([]*model.Barber, error) {
	var resultBarbers []*model.Barber
	var dbBarbers []barber.Barber
	dbBarbers = barber.GetAll()

	for _, barber := range dbBarbers {
		resultBarbers = append(resultBarbers, &model.Barber{
			BarberID:    barber.BarberID,
			ShopID:      barber.ShopID,
			UserName:    barber.UserName,
			Password:    barber.Password,
			FirstName:   barber.FirstName,
			LastName:    barber.LastName,
			PhoneNumber: barber.PhoneNumber,
			Gender:      barber.Gender,
			Dob:         barber.Dob,
			HireDate:    barber.HireDate,
			DismissDate: barber.DismissDate,
			SeatNum:     barber.SeatNum})
	}

	return resultBarbers, nil
}

func (r *queryResolver) Response(ctx context.Context) (*model.Response, error) {
	res := &model.Response{Error: "nothing here"}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
