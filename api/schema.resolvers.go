package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/generated"
	"api/internal/barber"
	"api/internal/database"
	"api/internal/methods"
	"api/jwt"
	"api/model"
	"context"
	"strings"
)

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

	if err := barber.SaveOne(); err != nil {
		res := &model.Response{Response: "", Error: "Error in SaveOne()."}
		// Figure out how to return an error.
		return res, err
	}

	res := &model.Response{Response: "Barber inserted", Error: ""}

	return res, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.UserLogin) (*model.Response, error) {
	var barber barber.Barber

	barber.UserName = input.Username
	barber.Password = input.Password
	if kosher := barber.Authenticate(); !kosher {
		res := &model.Response{Response: "", Error: "Authentication error."}
		// I removed &methods.WrongUsernameOrPasswordError{}.
		// Figure out how to parse it in apollo android as an error response.
		return res, nil
	}
	token, err := jwt.GenerateToken(barber.UserName)
	if err != nil {
		res := &model.Response{Response: "", Error: "Error generating token"}
		return res, err
	}
	res := &model.Response{Response: token, Error: ""}
	return res, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (*model.Response, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		res := &model.Response{Response: "", Error: "Access Denied"}
		return res, err
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		res := &model.Response{Response: "", Error: "Error Generating token"}
		return res, err
	}
	res := &model.Response{Response: token, Error: ""}
	return res, nil
}

func (r *queryResolver) GetAllBarbers(ctx context.Context) ([]*model.Barber, error) {
	var resultBarbers []*model.Barber
	var dbBarbers []barber.Barber

	dbBarbers, err := barber.GetAll()
	if err != nil {
		return nil, err
	}

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

func (r *queryResolver) GetBarberByUsername(ctx context.Context, username string) (*model.Barber, error) {
	var resultBarber *model.Barber
	var dbBarber barber.Barber

	dbBarber.UserName = username
	if err := dbBarber.Get(); err != nil {
		return nil, err
	}

	resultBarber = &model.Barber{
		BarberID:    dbBarber.BarberID,
		ShopID:      dbBarber.ShopID,
		UserName:    dbBarber.UserName,
		Password:    dbBarber.Password,
		FirstName:   dbBarber.FirstName,
		LastName:    dbBarber.LastName,
		PhoneNumber: dbBarber.PhoneNumber,
		Gender:      dbBarber.Gender,
		Dob:         dbBarber.Dob,
		HireDate:    dbBarber.HireDate,
		DismissDate: dbBarber.DismissDate,
		SeatNum:     dbBarber.SeatNum}
	return resultBarber, nil
}

func (r *queryResolver) GetAppointmentsByUsername(ctx context.Context, username string) ([]*model.BarberAppointment, error) {
	getAppointments := `select  bfirstname, blastname,
		shopname, streetaddr,
		apptdate, starttime, endtime, paymenttype, clientcancelled, barbercancelled,
		cfirstname, clastname,
		servicename, servicedescription, price, customduration
		from appt_details where username = $1 order by starttime`

	stmt, err := database.Db.Prepare(getAppointments)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// var resultBarberAppointment []*model.BarberAppointment
	resultBarberAppointment := []*model.BarberAppointment{}
	for rows.Next() {
		barberAppt := &model.BarberAppointment{
			Barber:      new(model.Barber),
			Shop:        new(model.Shop),
			Appointment: &model.Appointment{},
			Client:      &model.MinClient{},
			Service:     &model.Service{}}

		err := rows.Scan(
			&barberAppt.Barber.FirstName, &barberAppt.Barber.LastName,
			&barberAppt.Shop.ShopName, &barberAppt.Shop.StreetAddr,
			&barberAppt.Appointment.ApptDate, &barberAppt.Appointment.StartTime,
			&barberAppt.Appointment.EndTime, &barberAppt.Appointment.PaymentType,
			&barberAppt.Appointment.ClientCancelled,
			&barberAppt.Appointment.BarberCancelled,
			&barberAppt.Client.FirstName, &barberAppt.Client.LastName,
			&barberAppt.Service.ServiceName, &barberAppt.Service.ServiceDescription,
			&barberAppt.Service.Price, &barberAppt.Service.Duration)

		barberAppt.Appointment.ApptDate = methods.RemoveSuffix(
			barberAppt.Appointment.ApptDate)

		tmpS := methods.RemovePrefix(barberAppt.Appointment.StartTime)
		tmpS = strings.TrimSuffix(tmpS, ":00Z")
		barberAppt.Appointment.StartTime = tmpS

		if err != nil {
			return nil, err
		}

		resultBarberAppointment = append(resultBarberAppointment, barberAppt)
	}

	return resultBarberAppointment, nil
}

func (r *queryResolver) GetAppointmentByDateRange(ctx context.Context, input model.BarberDateRange) ([]*model.BarberAppointment, error) {
	getApptByDateRange := `select  bfirstname, blastname,
		shopname, streetaddr,
		apptdate, extract(hour from starttime), extract(hour from endtime), paymenttype, clientcancelled, barbercancelled,
		cfirstname, clastname,
		servicename, servicedescription, price, customduration
		from appt_details where apptdate >= $1 and apptdate <= $2 and username=$3
		order by starttime`

	stmt, err := database.Db.Prepare(getApptByDateRange)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(input.StartDate, input.EndDate, input.UserName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// var resultBarberAppointment []*model.BarberAppointment
	resultBarberAppointment := []*model.BarberAppointment{}
	for rows.Next() {
		barberAppt := &model.BarberAppointment{
			Barber:      new(model.Barber),
			Shop:        new(model.Shop),
			Appointment: &model.Appointment{},
			Client:      &model.MinClient{},
			Service:     &model.Service{}}
		err := rows.Scan(
			&barberAppt.Barber.FirstName, &barberAppt.Barber.LastName,
			&barberAppt.Shop.ShopName, &barberAppt.Shop.StreetAddr,
			&barberAppt.Appointment.ApptDate, &barberAppt.Appointment.StartTime,
			&barberAppt.Appointment.EndTime, &barberAppt.Appointment.PaymentType,
			&barberAppt.Appointment.ClientCancelled,
			&barberAppt.Appointment.BarberCancelled,
			&barberAppt.Client.FirstName, &barberAppt.Client.LastName,
			&barberAppt.Service.ServiceName, &barberAppt.Service.ServiceDescription,
			&barberAppt.Service.Price, &barberAppt.Service.Duration)
		if err != nil {
			return nil, err
		}
		resultBarberAppointment = append(resultBarberAppointment, barberAppt)
	}
	return resultBarberAppointment, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
