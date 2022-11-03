package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/paularah/wallet/pkg/db/sqlc"
	"github.com/paularah/wallet/pkg/util"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	testConfig := util.Config{AcessTokenDuration: time.Minute * 3, JwtSecretKey: util.RandomEmail()}

	server := NewServer(store, testConfig)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
