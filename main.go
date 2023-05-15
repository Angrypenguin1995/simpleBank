package main

import (
	"database/sql"
	"log"

	"github.com/angrypenguin1995/simple__bank/api"
	db "github.com/angrypenguin1995/simple__bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	serverAddress = "0.0.0.0:8080"
	dbDriver      = "postgres"
	// dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
	dbSource = "postgresql://root:password@localhost:54000/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
