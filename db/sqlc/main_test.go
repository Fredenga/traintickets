package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB
//! One global testQuery
//!contains DBTX which can be a connection or transaction
//!It calls the individual database functions in the unit test
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5434/traintickets?sslmode=disable"
)
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	testQueries = New(testDB)//!New defined in db.go
	//TODO: connect unit tests to the database
	os.Exit(m.Run()) //!Report exit code status to test runner
}
//!The main entry point for all unit tests inside a go package
//TODO: install golang lib/pq, sql is generic therefore it requires a database driver