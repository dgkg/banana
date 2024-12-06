package model

import (
	"encoding/json"
	"testing"
)

func TestPassword(t *testing.T) {
	// TestPasswordUnmarshalJSON tests the behavior of the Marshal method
	// of the Password type.
	t.Run("TestPasswordUnmarshalJSON", func(t *testing.T) {
		// Create a new Password value.
		p := Password("password")
		// Create a JSON representation of the Password value.
		b, err := json.Marshal(p)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if string(b) != "\"****\"" {
			t.Errorf("expected: %v, got: %v", "\"****\"", string(b))
		}
	})
	// TestPasswordString tests the behavior of the String method of the Password
	// type.
	t.Run("TestPasswordString", func(t *testing.T) {
		// Create a new Password value.
		p := Password("password")
		// Check if the string representation of the Password value is correct.
		if p.String() != "****" {
			t.Errorf("expected: %v, got: %v", "****", p.String())
		}
	})
	// TestPasswordUnmarshal tests the behavior of the Unmarshal method of the
	// Password type.
	t.Run("TestPasswordUnmarshal", func(t *testing.T) {
		// Create a new Password value.
		data := []byte(`{"pass": "password"}`)
		myStruct := struct {
			Pass Password `json:"pass"`
		}{}

		err := json.Unmarshal(data, &myStruct)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		testpass := Password("5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8")
		if myStruct.Pass != testpass {
			t.Errorf("expected: %v, got: %v", "password", myStruct.Pass)
		}
	})
}
