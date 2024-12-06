package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"banana/apifront/db"
	"banana/apifront/model"
)

func GetTestGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	return ctx, w
}

func TestUserHandlers(t *testing.T) {

	// Create a new User value.
	u := model.NewUser("John", "Doe", "john@doe.fr", "password")
	log.Println("user:", u)

	// Create a new Handler value.
	dbMoke := db.NewMoke()
	handl := NewHandler(dbMoke)
	handl.db.SetUser(u)

	t.Run("GetUserByID", func(t *testing.T) {
		// Create a new gin context.
		ctx, w := GetTestGinContext()
		// create the params for the request
		ctx.AddParam("uuid", u.UUID)
		// Set the user in the database.
		// Call the GetUserByID method of the Handler value.
		handl.GetUserByID(ctx)
		// Check if the response status code is 200.
		if ctx.Writer.Status() != http.StatusOK {
			t.Errorf("expected: %v, got: %v", http.StatusOK, ctx.Writer.Status())
		}
		// Check if the response body is correct.
		urs, err := io.ReadAll(w.Body)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		log.Println("body:", string(urs))
		// Check if the response body is correct.
		// usualy we should use reflect.DeepEqual
		// but here we are sure that the password is hashed
		var user model.User
		err = json.Unmarshal(urs, &user)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if user.UUID != u.UUID {
			t.Errorf("expected: %v, got: %v", u.UUID, user.UUID)
		}
		if user.FirstName != u.FirstName {
			t.Errorf("expected: %v, got: %v", u.FirstName, user.FirstName)
		}
		if user.LastName != u.LastName {
			t.Errorf("expected: %v, got: %v", u.LastName, user.LastName)
		}
		if user.Email != u.Email {
			t.Errorf("expected: %v, got: %v", u.Email, user.Email)
		}
	})
	t.Run("Register", func(t *testing.T) {

		// Create a new UserRegisterPayload value.
		payload := model.UserRegisterPayload{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane@doe.fr",
			Password:  "password",
		}
		// Create a JSON representation of the UserRegisterPayload value.
		b, err := json.Marshal(payload)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// Create a new gin context.
		ctx, w := GetTestGinContext()
		// Set the request body.
		bod := bytes.NewReader(b)
		// bod.Write(b)
		ctx.Request.Body = bod
		handl.Register(ctx)

		urs, err := io.ReadAll(w.Body)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		log.Println("body:", string(urs))

	})
}
