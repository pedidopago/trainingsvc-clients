package utils

import "testing"

func TestSecureID(t *testing.T) {
	if len(SecureID().String()) != 26 {
		t.Fail()
	}
	if !IsIDValid(SecureID().String()) {
		t.Fail()
	}
	if IsIDValid("asdasdasdasdqweqweqwe") {
		t.Fail()
	}
}
