package test

import (
	"testing"

	"github.com/tgkzz/helper"
)

func TestToUpper(t *testing.T) {
	got := helper.ToUpper("Hello my name is KAMAL")

	want := "HELLO MY NAME IS KAMAL"

	if got != want {
		t.Errorf("Function1() = %s; expected %s", got, want)
	}
}
