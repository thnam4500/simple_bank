package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/thnam4500/simple_bank/db/sqlc"
)

// Server serves HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PUT("/accounts/", server.updateAccountBalance)
	router.PUT("/accounts/add-balance", server.addBalance)

	server.router = router
	return server
}
