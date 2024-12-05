package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"

	"banana/db"
	"banana/model"
)

type Handler struct {
	db *db.DB
}

func NewHandler(db *db.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Login(ctx *gin.Context) {
	// bin from body payload
	var payload model.UserLoginPayload
	err := ctx.Bind(&payload)
	if err != nil {
		err := NewErrorValidation("user", "invalid payload", err)
		respError(ctx, "user", err)
		return
	}
	// get user in db
	u, err := h.db.GetUserByEmail(payload.Email)
	if err != nil {
		respError(ctx, "user", err)
		return
	}
	// verify if the user has the good credentials
	if u.Password != payload.Password {
		err := NewErrorAutorization("user", u.UUID)
		respError(ctx, "user", err)
		return
	}
	// user is authorized
	ctx.JSON(200, gin.H{"ok": true})
}

func (h *Handler) Register(ctx *gin.Context) {
	var payload model.UserRegister
	err := ctx.Bind(&payload)
	// data, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	ctx.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }
	// log.Println("data in body:", string(data))
	// err = json.Unmarshal(data, &payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = validator.New().Struct(payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.Printf("UserRegisterPayload : %#v", payload)
	usr := model.NewUser(payload.FirstName, payload.LastName, payload.Email, payload.Password)
	err = h.db.SetUser(usr)
	if err != nil {
		respError(ctx, "user", err)
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

func (h *Handler) GetUserByID(ctx *gin.Context) {
	uuidParam := ctx.Param("uuid")
	log.Println("HandlerGetUserByID: uuidParam", uuidParam)
	_, err := uuid.Parse(uuidParam)
	log.Println("HandlerGetUserByID: err", err)
	if err != nil {
		log.Println("validation uuid", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	u, err := h.db.GetUserByID(uuidParam)
	log.Println("HandlerGetUserByID: GetUserByID:", u, err)
	if err != nil {
		respError(ctx, "user", err)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func (h *Handler) SearchUser(ctx *gin.Context) {
	name := ctx.Query("name")
	if len(name) > 0 {
		u, err := h.db.GetUserByName(name)
		if err != nil {
			respError(ctx, "user", err)
			return
		}
		ctx.JSON(http.StatusOK, u)
		return
	}
	u, err := h.db.GetAllUser()
	if err != nil {
		respError(ctx, "user", err)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func respNotfound(ctx *gin.Context, entity string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"error":   entity + " not found",
	})
}

func respError(ctx *gin.Context, entity string, err error) {
	log.Println(err.Error())
	switch err := err.(type) {
	case *ErrorHandlerValidation:
		ctx.JSON(err.StatusCode, gin.H{
			"success": false,
			"error":   err.Message,
		})
	case *db.ErrorDB:
		ctx.JSON(err.StatusCode, gin.H{
			"success": false,
			"error":   err.Message,
		})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   entity + " " + err.Error(),
		})
	}
}
