package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/thnam4500/simple_bank/db/sqlc"
	"github.com/thnam4500/simple_bank/token"
	"github.com/thnam4500/simple_bank/util"
)

// Server serves HTTP requests
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can't create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupServer()
	return server, nil
}

func (s *Server) setupServer() {
	router := gin.Default()
	router.POST("/accounts", s.createAccount)
	router.GET("/accounts/:id", s.getAccount)
	router.GET("/accounts", s.listAccount)
	router.DELETE("/accounts/:id", s.deleteAccount)
	router.PUT("/accounts/", s.updateAccountBalance)
	router.PUT("/accounts/add-balance", s.addBalance)

	router.POST("/transfer", s.createTransfer)

	router.POST("/users", s.createUser)
	router.POST("/users/login", s.loginUser)
	s.router = router
}
