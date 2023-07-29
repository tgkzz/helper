package test

import (
	"testing"

	"github.com/tgkzz/helper"
)

func TestToLower(t *testing.T) {
	got := helper.ToLower("Hello my name is KAMAL")

	want := "hello my name is kamal"

	if got != want {
		t.Errorf("Function1() = %s; expected %s", got, want)
	}
}
