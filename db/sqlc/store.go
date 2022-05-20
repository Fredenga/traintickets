package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/fredenga/traintickets/util"
)

//TODO: EXECUTE DATABASE TRANSACTION INSTEAD OF INDIVIDUAL DATABASE QUERY
//!Provides all functions to execute db queries individually
type Store struct {
	*Queries 
	db *sql.DB
/* //!composition: embedding structs in other structs, all functionality is extended
	//!requires db connection 
	*/
	
}
//TODO: Create a new global store(db is determinant)
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}
//TODO: executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil) //BEGIN
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q) //query function
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr) //ROLLBACK
		}
		return err
	}
	return tx.Commit()//COMMIT
}

type BookTrainTxParams struct {
	Route       string 		   `json:"route"`
	TripDate    time.Time      `json:"trip_date"`
	Email       string         `json:"email"`
	CreditCardNumber int32     `json:"credit_card_number"`
}
type BookTrainTxResult struct {
	Ticket 		Ticket 		`json:"ticket"`
	Timetable 	Timetable	`json:"timetable"`
	Train		Train		`json:"train"`
}

var txKey = struct{}{}//Must Not inbuilt type

func (store *Store) BookTrainTx(ctx context.Context, arg BookTrainTxParams) (BookTrainTxResult, error) {
	var result BookTrainTxResult
	err := store.execTx(ctx, func(q *Queries) error{
		var err error
		txName := ctx.Value(txKey)//Available

		_, err = q.Getuser(ctx, arg.Email)
		if err != nil {
			return err
		}
		fmt.Println(txName, "Get timetable using route")
		result.Timetable, err = q.Gettimetable(ctx, arg.Route)
		if err != nil {
			return err
		}

		fmt.Println(txName, "Get all trains travelling via this route")

		var trains []Train
		trains, err = q.Listtrainsbyroutes(ctx, ListtrainsbyroutesParams{
			Route: result.Timetable.Route,
			Limit: 4,
		})
		if err != nil {
			return err
		}
		
		fmt.Println(txName, "Generate train ticket")
		result.Train = trains[0]
		result.Ticket, err = q.Createticket(ctx, CreateticketParams{
			Route: result.Train.Route,
			TrainNumber: result.Train.TrainNumber,
			CoachNumber: util.RandomNum(),
			SeatNumber: util.RandomInt(1, int64(result.Train.MaxPassengerNo)),
			BookingDate: time.Now(),
			TripDate: arg.TripDate,
			Fare: result.Timetable.Price,
			Email: arg.Email,
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "Process train ticket payment")
		q.Createpayment(ctx, CreatepaymentParams{
			TicketID: result.Ticket.TicketID,
			Amount: result.Ticket.Fare,
			CreditCardNumber: arg.CreditCardNumber,
		})
		return nil

	})
	return result, err
}