package handler

import (
	"banana/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

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
	// create a jwt token
	jwtValue, err := h.newJWTToken(u.UUID, "secret")
	if err != nil {
		respError(ctx, "user", err)
		return
	}
	// user is authorized
	ctx.JSON(200, gin.H{"ok": true, "jwt": jwtValue})
}

func (h *Handler) Register(ctx *gin.Context) {
	var payload model.UserRegisterPayload
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

func (h *Handler) UpdateUser(ctx *gin.Context) {
	uuidParam := ctx.Param("uuid")
	_, err := uuid.Parse(uuidParam)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var payload model.UserUpdatePayload
	err = ctx.Bind(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = validator.New().Struct(payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	usr, err := h.db.GetUserByID(uuidParam)
	if err != nil {
		respError(ctx, "user", err)
		return
	}
	if payload.FirstName != nil {
		usr.FirstName = *payload.FirstName
	}
	if payload.LastName != nil {
		usr.LastName = *payload.LastName
	}
	if payload.Email != nil {
		usr.Email = *payload.Email
	}
	ctx.JSON(http.StatusOK, usr)
}
