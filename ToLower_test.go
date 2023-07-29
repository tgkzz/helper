package helper

import (
	"testing"
)

func TestToLower(t *testing.T) {

	type addTest struct {
		got, want string
	}

	addTests := []addTest{
		{"hello", "hello"},
		{"HELLO", "hello"},
		{"Hello My Name is KAMAL", "hello my name is kamal"},
	}

	for _, test := range addTests {
		got := ToLower(test.got)

		if got != test.want {
			t.Errorf("Function1() = %s; expected %s", got, test.want)
		}
	}
}
