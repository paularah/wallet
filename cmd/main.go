package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	api "github.com/paularah/wallet/pkg/api"
	db "github.com/paularah/wallet/pkg/db/sqlc"
	"github.com/paularah/wallet/pkg/util"
)

// todo graceful shutdown and egde cases
func main() {

	config, err := util.LoadConfigFromEnv(".")

	if err != nil {
		log.Fatalf("%v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatalf("unable to connect to db: %v", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store, config)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("unable to start server: %v", err)
	}

}
