package helper

import (
	"testing"
)

func TestToUpper(t *testing.T) {

	type addTest struct {
		got, want string
	}

	addTests := []addTest{
		{"hello", "HELLO"},
		{"HELLO", "HELLO"},
		{"Hello My Name is KAMAL", "HELLO MY NAME IS KAMAL"},
	}

	for _, test := range addTests {
		got := ToUpper(test.got)

		if got != test.want {
			t.Errorf("Function1() = %s; expected %s", got, test.want)
		}
	}
}
