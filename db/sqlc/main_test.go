package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/angrypenguin1995/simple__bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	// testDB, err := sql.Open(dbDriver, dbSource) cause "nil pointer reference error", because both err and testDB have been declared earlier
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
