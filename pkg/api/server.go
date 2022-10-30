package api

import (
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/paularah/wallet/pkg/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/wallets", server.createWallet)
	router.GET("/wallets", server.listWallet)
	router.GET("/wallets/:id", server.getWallet)
	router.GET("/wallets/transfers", server.createTransfer)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	log.Print(err)
	return gin.H{"error": err}

}
