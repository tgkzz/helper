package helper

import (
	"testing"
)

func TestToLower(t *testing.T) {
	got := ToLower("Hello my name is KAMAL")

	want := "hello my name is kamal"

	if got != want {
		t.Errorf("Function1() = %s; expected %s", got, want)
	}
}
