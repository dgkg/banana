package testhandler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"banana/apifront/model"
)

type MokeTests struct{}

func (MokeTests) PostRequest(ctx *gin.Context, content interface{}) {
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

func (MokeTests) CheckUserFields(t *testing.T, data []byte, refUser *model.User) {
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

func (MokeTests) NewContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Request.Header.Set("Content-Type", "application/json")

	return ctx, w
}

func (MokeTests) CheckResponseError(t *testing.T, data []byte, errRef error) {
	// {"error":"not found","success":false}
	var errResponse ErrorForTestResponse
	err := json.Unmarshal(data, &errResponse)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(errResponse, errRef) {
		t.Errorf("expected: %v, got: %v", errRef, errResponse)
	}
}

type ErrorForTestResponse struct {
	Err     string `json:"error"`
	Success bool   `json:"success"`
}

func (ErrorForTestResponse) Error() string {
	return "error"
}
