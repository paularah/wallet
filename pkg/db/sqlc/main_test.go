package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/paularah/wallet/pkg/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfigFromEnv("../../..")

	if err != nil {
		log.Fatal("unable to load config")
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("unable to connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
