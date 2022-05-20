package db

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"

	"github.com/fredenga/traintickets/util"
	"github.com/stretchr/testify/require"
)
func createRandomTrain(t *testing.T) Train{
	timetable := createRandomTimetable(t)
	train, err := testQueries.Createtrain(context.Background(), CreatetrainParams{
		Type: util.RandomType("types"),
		Class: util.RandomType("class"),
		MaxPassengerNo: util.RandomInt(1000, 2000),
		MaxSpeed: util.RandomInt(100, 200),
		Route: timetable.Route,
	})
	require.NoError(t, err)
	require.NotEmpty(t, train)
	require.NotZero(t, train.TrainNumber)
	return train
}
func TestCreateTrain(t *testing.T){
	createRandomTrain(t)
}

func TestListTrains(t *testing.T) {
	trains, err := testQueries.Listtrains(context.Background(), 10)
	require.NoError(t, err)
	require.LessOrEqual(t, len(trains), 10)
	for _, train := range trains {
		require.NotZero(t, train.TrainNumber)
		require.NotEmpty(t, train)
	}
}
func TestListTrainsByRoutes(t *testing.T) {
	timetable := createRandomTimetable(t)
	trains, err := testQueries.Listtrainsbyroutes(context.Background(), ListtrainsbyroutesParams{
		Route: timetable.Route,
		Limit: 10,
	})
	require.NoError(t, err)
	require.LessOrEqual(t, len(trains), 10)
	for _, train := range trains {
		require.NotZero(t, train.TrainNumber)
		require.NotEmpty(t, train)
	}
}
func TestGetTrain(t *testing.T) {
	trains, errs := testQueries.Listtrains(context.Background(), 10)
	require.NoError(t, errs)

	for _, train := range trains {
		train1, err := testQueries.Gettrain(context.Background(), train.TrainNumber)
		require.NoError(t, err)
		require.NotZero(t, train1.TrainNumber)
		require.NotEmpty(t, train1)
	}
}
func TestDeleteTrain(t *testing.T) {
	trains, errs := testQueries.Listtrains(context.Background(), 10)
	require.NotEmpty(t, trains)

	require.NoError(t, errs)
	train1 := trains[rand.Intn(len(trains))]
	require.NotEmpty(t, train1)
	err := testQueries.Deletetrain(context.Background(), train1.TrainNumber)
	require.NoError(t, err)

	train2, err := testQueries.Gettrain(context.Background(), train1.TrainNumber)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, train2)

}
