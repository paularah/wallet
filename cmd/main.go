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
		log.Fatal("unable to load config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("unable to connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("unable to start server: ", err)
	}

}
