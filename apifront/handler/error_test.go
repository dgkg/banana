package handler

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"banana/apifront/db"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewErrorValidation(t *testing.T) {
	err := errors.New("validation error")
	validationErr := NewErrorValidation("TestEntity", "Test message", err)

	assert.Equal(t, "TestEntity", validationErr.Entity)
	assert.Equal(t, "Test message", validationErr.Message)
	assert.Equal(t, err, validationErr.Err)
	assert.Equal(t, 400, validationErr.StatusCode)
}

func TestNewErrorAutorization(t *testing.T) {
	authErr := NewErrorAutorization("TestEntity", "1234")

	assert.Equal(t, "TestEntity 1234", authErr.Entity)
	assert.Equal(t, "user not authorized", authErr.Message)
	assert.Nil(t, authErr.Err)
	assert.Equal(t, 401, authErr.StatusCode)
}

func TestRespError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		entity         string
		err            error
		expectedStatus int
		expectedBody   gin.H
	}{
		{
			name:           "Validation Error",
			entity:         "TestEntity",
			err:            NewErrorValidation("TestEntity", "Test message", errors.New("validation error")),
			expectedStatus: 400,
			expectedBody:   gin.H{"success": false, "error": "Test message"},
		},
		{
			name:           "Authorization Error",
			entity:         "TestEntity",
			err:            NewErrorAutorization("TestEntity", "1234"),
			expectedStatus: 401,
			expectedBody:   gin.H{"success": false, "error": "user not authorized"},
		},
		{
			name:           "Database Error",
			entity:         "TestEntity",
			err:            &db.ErrorDB{Message: "database error", StatusCode: 500},
			expectedStatus: 500,
			expectedBody:   gin.H{"success": false, "error": "database error"},
		},
		{
			name:           "Unknown Error",
			entity:         "TestEntity",
			err:            errors.New("unknown error"),
			expectedStatus: 500,
			expectedBody:   gin.H{"success": false, "error": "TestEntity unknown error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			respError(c, tt.entity, tt.err)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, toJSON(tt.expectedBody), w.Body.String())
		})
	}
}

func toJSON(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
