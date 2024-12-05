package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Password string

func (p *Password) UnmarshalJSON(b []byte) error {
	aux := ""
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}
	h := sha256.New()
	h.Write([]byte(aux))
	*p = Password(fmt.Sprintf("%x", h.Sum(nil)))
	return nil
}

func (p Password) String() string {
	return "****"
}

func (p Password) MarshalJSON() ([]byte, error) {
	var s string = "****"
	return json.Marshal(s)
}
