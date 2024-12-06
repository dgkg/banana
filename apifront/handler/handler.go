package handler

import (
	"github.com/gin-gonic/gin"

	"banana/apifront/db"
	"banana/concert"
)

type Handler struct {
	db      db.DB
	concert *concert.SDKAPI
}

func NewHandler(db db.DB, concert *concert.SDKAPI) *Handler {
	return &Handler{
		db:      db,
		concert: concert,
	}
}

func (h *Handler) InitRoutes(r *gin.Engine) {
	authWithJWT := VerifyJWTToken("secret")

	r.POST("/register", BasicAuth, h.Register)
	r.POST("/login", h.Login)
	r.GET("/users/:uuid", authWithJWT, h.GetUserByID)
	r.GET("/users", authWithJWT, h.SearchUser)
}
