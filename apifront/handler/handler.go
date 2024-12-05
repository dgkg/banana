package handler

import (
	"github.com/gin-gonic/gin"

	"banana/apifront/db"
)

type Handler struct {
	db db.DB
}

func NewHandler(db db.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) InitRoutes(r *gin.Engine) {
	authWithJWT := VerifyJWTToken("secret")

	r.POST("/register", BasicAuth, h.Register)
	r.POST("/login", h.Login)
	r.GET("/users/:uuid", authWithJWT, h.GetUserByID)
	r.GET("/users", authWithJWT, h.SearchUser)
}
