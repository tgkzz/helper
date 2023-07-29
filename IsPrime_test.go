package helper

import (
	"testing"
)

func TestIsPrime(t *testing.T) {

	type addTest struct {
		got  int
		want bool
	}

	addTests := []addTest{
		{2, true},
		{3, true},
		{8, false},
		{11, true},
		{97, true},
		{703, false},
	}

	for _, test := range addTests {
		got := IsPrime(test.got)

		if got != test.want {
			t.Errorf("IsPrime() = %t; expected %t", got, test.want)
		}
	}
}
