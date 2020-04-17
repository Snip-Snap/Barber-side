package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphqltest/api/generated"
	"graphqltest/api/model"

	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) SignupClient(ctx context.Context, input model.NewClient) (*model.Response, error) {
	print("here\n")
	statement, err := db.Prepare("insert into client (fullname, gender, phonenumber, username, hashedpassword) values($1, $2, $3, $4, $5)")
	CheckError(err)
	print("here\n")
	//TODO: Check inputs for uniqueness and appropriate characters.
	// pw := string(input.Password)
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password),
		bcrypt.DefaultCost)
	CheckError(err)

	hashedInputpw := string(hash)
	_, err = statement.Exec(input.FullName,
		input.Gender, input.PhoneNumber, input.UserName, hashedInputpw)
	CheckError(err)

	res := &model.Response{Error: "Okay"}

	return res, nil
}

func (r *mutationResolver) SignUpBarber(ctx context.Context, input model.NewBarber) (*model.Response, error) {
	stmt, err := db.Prepare("insert into barber (fullname, gender, phonenumber, username, hashedpassword) values($1, $2, $3, $4, $5)")

	if err != nil {
		return nil, err
	}

	hashpw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(input.FullName, input.Gender, input.PhoneNumber, input.UserName, string(hashpw))
	if err != nil {
		return nil, err
	}
	res := &model.Response{Error: "Okay"}

	return res, nil
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
