package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/thnam4500/simple_bank/db/sqlc"
	"github.com/thnam4500/simple_bank/util"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type getUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
}

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:     req.Username,
		HashPassword: hashPassword,
		FullName:     req.Fullname,
		Email:        req.Email,
	}
	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := createUserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}
	ctx.JSON(http.StatusOK, resp)
}
