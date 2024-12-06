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
	"github.com/google/uuid"

	"banana/apifront/db"
	"banana/apifront/model"
)

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
		payloadResponse, err := io.ReadAll(w.Body)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		log.Println("body:", string(payloadResponse))
		// Check if the response body is correct.
		// usualy we should use reflect.DeepEqual
		// but here we are sure that the password is hashed
		checkUserFields(t, payloadResponse, u)
	})
	t.Run("Register", func(t *testing.T) {

		// Create a new UserRegisterPayload value.
		payload := model.UserRegisterPayload{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane.test@doebibi.fr",
			Password:  "password",
		}
		urJane := &model.User{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane.test@doebibi.fr",
			Password:  "password",
		}

		// Create a new gin context.
		ctx, w := GetTestGinContext()

		// Mock the request with the payload value in the body.
		mockJsonPost(ctx, payload)

		// Call the Register method of the Handler value.
		handl.Register(ctx)

		// Check if the response status code is 200.
		payloadResponse, err := io.ReadAll(w.Body)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ctx.Writer.Status() != http.StatusOK {
			t.Errorf("expected: %v, got: %v", http.StatusOK, ctx.Writer.Status())
		}

		// Check if the response body is correct.
		log.Println("body:", string(payloadResponse))
		checkUserFields(t, payloadResponse, urJane)
	})
}

func mockJsonPost(ctx *gin.Context, content interface{}) {
	ctx.Request.Method = "POST" // or PUT
	ctx.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func checkUserFields(t *testing.T, data []byte, refUser *model.User) {
	var user model.User
	err := json.Unmarshal(data, &user)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if _, err = uuid.Parse(user.UUID); err != nil {
		t.Errorf("expected: a valid uuid got %v", user.UUID)
	}
	if user.FirstName != refUser.FirstName {
		t.Errorf("expected: %v, got: %v", refUser.FirstName, user.FirstName)
	}
	if user.LastName != refUser.LastName {
		t.Errorf("expected: %v, got: %v", refUser.LastName, user.LastName)
	}
	if user.Email != refUser.Email {
		t.Errorf("expected: %v, got: %v", refUser.Email, user.Email)
	}
}

func GetTestGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Request.Header.Set("Content-Type", "application/json")

	return ctx, w
}
