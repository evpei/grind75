package validpalindrome

import (
	"testing"
)

type TestData struct {
	give string
	want bool
}

func TestSolve(t *testing.T) {
	tests := []TestData{
		{"0P", false},
		{" ", true},
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for index, test := range tests {
		res := Solve(test.give)
		if res != test.want {
			t.Fatalf(`Test #%d failed. Input: "%s". Expected "%v", but got "%v"`, index+1, test.give, test.want, res)
		}
	}
}
