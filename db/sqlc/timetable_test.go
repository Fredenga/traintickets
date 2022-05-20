package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/fredenga/traintickets/util"
	"github.com/stretchr/testify/require"
)

func createRandomTimetable(t *testing.T) Timetable {
	arg := CreatetimetableParams{
		Route: util.RandomRoute(),
		Departure: time.Now(),
		Arrival: time.Now().Add(time.Hour*2),
		Distance: util.RandomInt(3, 10),
		Price: util.RandomMoney(),
	}
	timetable, err := testQueries.Createtimetable(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, timetable)
	require.Equal(t, arg.Route, timetable.Route)
	require.Equal(t, arg.Distance, timetable.Distance)
	require.Equal(t, arg.Price, timetable.Price)
	require.WithinDuration(t, arg.Arrival, timetable.Arrival, time.Second)
	require.WithinDuration(t, arg.Departure, timetable.Departure, time.Second)
	return timetable
}

func TestCreateTimeTable(t *testing.T){
	createRandomTimetable(t)
}
func TestGetTimeTable(t *testing.T){
	timetable1 := createRandomTimetable(t)
	timetable2, err := testQueries.Gettimetable(context.Background(), timetable1.Route)
	require.NoError(t, err)
	require.NotEmpty(t, timetable2)
	require.Equal(t, timetable1.Arrival, timetable2.Arrival)
	require.Equal(t, timetable1.Departure, timetable2.Departure)
	require.Equal(t, timetable1.Route, timetable2.Route)
	require.Equal(t, timetable1.Distance, timetable2.Distance)
	require.Equal(t, timetable1.Price, timetable2.Price)
}

func TestListTimeTables(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomTimetable(t)
	}
	timetables, err := testQueries.Listtimetables(context.Background(), 5)
	require.NoError(t, err)
	require.LessOrEqual(t, len(timetables), 5)
	for _, timetable := range timetables {
		require.NotEmpty(t, timetable)
	}
}

func TestDeleteTimeTable(t *testing.T) {
	timetable1 := createRandomTimetable(t)
	err := testQueries.Deletetimetable(context.Background(), timetable1.Route)
	require.NoError(t, err)
	timetable2, err := testQueries.Gettimetable(context.Background(), timetable1.Route)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, timetable2)
}