package db

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"github.com/fredenga/traintickets/util"
	"github.com/stretchr/testify/require"
)

func createRandomTicket(t *testing.T) Ticket {
	train := createRandomTrain(t)
	user := createRandomUser(t)
	timetable, _ := testQueries.Gettimetable(context.Background(), train.Route)
	arg := CreateticketParams{
		Route: train.Route,
		TrainNumber: train.TrainNumber,
		CoachNumber: util.RandomNum(),
		SeatNumber: rand.Int31n(train.MaxPassengerNo),
		BookingDate: time.Now(),
		TripDate: util.RandomTime(),
		Fare: timetable.Price,
		Email: user.Email,
	}

	ticket, err := testQueries.Createticket(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ticket)
	require.NotZero(t, ticket.TicketID)
	require.Equal(t, arg.Email, ticket.Email)
	require.Equal(t, arg.TrainNumber, ticket.TrainNumber)
	require.Equal(t, arg.CoachNumber, ticket.CoachNumber)
	require.Equal(t, arg.SeatNumber, ticket.SeatNumber)
	require.Equal(t, arg.Fare, ticket.Fare)
	require.Equal(t, arg.Route, ticket.Route)
	// require.WithinDuration(t, arg.BookingDate, ticket.BookingDate, time.Second)
	// require.WithinDuration(t, arg.TripDate, ticket.TripDate, time.Second)
	return ticket
}

func TestCreateTicket(t *testing.T) {
	createRandomTicket(t,)
}

func TestGetTicket(t *testing.T) {
	tickets, errs := testQueries.Listtickets(context.Background(), 10)
	require.NoError(t, errs)

	for _, ticket := range tickets {
		ticket1, err := testQueries.Getticket(context.Background(), ticket.TicketID)
		require.NoError(t, err)
		require.NotZero(t, ticket1.TicketID)
		require.NotEmpty(t, ticket1)
	}
}

func TestListTickets(t *testing.T) {
	tickets, errs := testQueries.Listtickets(context.Background(), 10)
	require.NoError(t, errs)

	for _, ticket := range tickets {
		require.NotZero(t, ticket.TicketID)
		require.NotEmpty(t, ticket)
	}
}

func TestDeleteTicket(t *testing.T) {
	ticket1 := createRandomTicket(t)
	err := testQueries.Deleteticket(context.Background(), ticket1.TicketID)
	require.NoError(t, err)
	ticket2, err := testQueries.Getticket(context.Background(), ticket1.TicketID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ticket2)

}