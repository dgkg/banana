package handler

import (
	"io"
	"net/http"
	"testing"

	"banana/apifront/db"
	"banana/apifront/handler/testhandler"
	"banana/apifront/model"
)

func TestUserHandlers(t *testing.T) {
	// Create a new User value.
	u := model.NewUser("John", "Doe", "john@doe.fr", "password")
	// Create a new Handler value.
	dbMoke := db.NewMoke()
	handl := NewHandler(dbMoke)
	handl.db.SetUser(u)

	m := testhandler.MokeTests{}
	t.Run("GetUserByID", func(t *testing.T) {
		// Create a new gin context.
		ctx, w := m.NewContext()
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
		// Check if the response body is correct.
		// usualy we should use reflect.DeepEqual
		// but here we are sure that the password is hashed
		m.CheckUserFields(t, payloadResponse, u)
	})
	t.Run("GetUserByID undefined", func(t *testing.T) {
		// Create a new gin context.
		ctx, w := m.NewContext()
		// create the params for the request
		ctx.AddParam("uuid", "f0835e25-f332-4d20-9a01-7a7d31eb06fc")
		// Set the user in the database.
		// Call the GetUserByID method of the Handler value.
		handl.GetUserByID(ctx)
		// Check if the response status code is 200.
		if ctx.Writer.Status() != http.StatusNotFound {
			t.Errorf("expected: %v, got: %v", http.StatusOK, ctx.Writer.Status())
		}
		// Check if the response body is correct.
		payloadResponse, err := io.ReadAll(w.Body)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// Check if the response body is correct.
		errNotFound := testhandler.ErrorForTestResponse{
			Err:     "not found",
			Success: false,
		}

		m.CheckResponseError(t, payloadResponse, errNotFound)
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
		ctx, w := m.NewContext()

		// Mock the request with the payload value in the body.
		m.PostRequest(ctx, payload)

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
		m.CheckUserFields(t, payloadResponse, urJane)
	})
}
