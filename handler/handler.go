package handler

import (
	"github.com/gin-gonic/gin"

	"banana/db"
)

type Handler struct {
	db *db.DB
}

func NewHandler(db *db.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) InitRoutes(r *gin.Engine) {
	r.POST("/register", BasicAuth, h.Register)
	r.POST("/login", h.Login)
	r.GET("/users/:uuid", h.GetUserByID)
	r.GET("/users", h.SearchUser)
}
