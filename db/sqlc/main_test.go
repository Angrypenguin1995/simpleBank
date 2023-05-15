package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	// dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
	dbSource = "postgresql://root:password@localhost:54000/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	// testDB, err := sql.Open(dbDriver, dbSource) cause "nil pointer reference error", because both err and testDB have been declared earlier
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
