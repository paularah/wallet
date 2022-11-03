package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/paularah/wallet/pkg/db/sqlc"
	"github.com/paularah/wallet/pkg/jwt"
	"github.com/paularah/wallet/pkg/util"
)

type Server struct {
	config  util.Config
	store   db.Store
	router  *gin.Engine
	tokener jwt.Tokener
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func NewServer(store db.Store, config util.Config) *Server {
	server := &Server{store: store, config: config}

	tokener := jwt.NewTokener(config.JwtSecretKey)
	server.tokener = tokener
	server.buildRoutes()

	return server
}

func (server *Server) buildRoutes() {
	router := gin.Default()

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://foo.com"},
	// 	AllowMethods:     []string{"PUT", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUserWithEmail)
	router.POST("/users/auth/renew_token", server.renewAcessTokenFromRefreshToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokener))
	authRoutes.POST("/transfers", server.createTransfer)
	authRoutes.POST("/wallets", server.createWallet)
	authRoutes.GET("/wallets", server.listWallet)
	authRoutes.GET("/wallets/:id", server.getWallet)

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err}
}
