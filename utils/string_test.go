package utils

import "testing"

func TestNonEmptyString(t *testing.T) {
	x := NonEmptyString("", "", "a", "b")
	if x != "a" {
		t.Fail()
	}
	x = NonEmptyString("", "", "")
	if x != "" {
		t.Fail()
	}
}
