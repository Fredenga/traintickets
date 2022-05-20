package db

import (
	"context"
	"testing"

	"github.com/fredenga/traintickets/util"
	"github.com/stretchr/testify/require"
)

func createRandomPayment(t *testing.T) Payment {
	ticket := createRandomTicket(t)
	arg := CreatepaymentParams{
		TicketID: ticket.TicketID,
		Amount: ticket.Fare,
		CreditCardNumber: util.RandomNum(),
	}
	payment, err := testQueries.Createpayment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, payment)
	require.Equal(t, arg.Amount, payment.Amount)
	require.Equal(t, arg.TicketID, payment.TicketID)
	require.Equal(t, arg.CreditCardNumber, payment.CreditCardNumber)
	return payment
}

func TestCreatePayment(t *testing.T) {
	createRandomPayment(t)
}

func TestGetPayment(t *testing.T) {
	payments, errs := testQueries.Listpayments(context.Background(), 10)
	require.NoError(t, errs)

	for _, payment := range payments {
		payment1, err := testQueries.Getpayment(context.Background(), payment.TicketID)
		require.NoError(t, err)
		require.NotZero(t, payment1.TicketID)
		require.NotEmpty(t, payment1)
	}
}

func TestListPayments(t *testing.T) {
	payments, errs := testQueries.Listpayments(context.Background(), 10)
	require.NoError(t, errs)

	for _, payment := range payments {
		require.NotZero(t, payment.TicketID)
		require.NotEmpty(t, payment)
	}
}
