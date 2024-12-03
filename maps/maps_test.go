package maps

import "testing"

func TestCreateAndAdd(t *testing.T) {
	m := New()
	m.Add("key", "value")
	m.Add("toto", "titi")
	if m.Get("key") != "value" {
		t.Error("Expected value to be 'value' got", m.Get("key"))
	}
	if m.Get("toto") != "titi" {
		t.Error("Expected value to be 'titi' got", m.Get("toto"))
	}
}
