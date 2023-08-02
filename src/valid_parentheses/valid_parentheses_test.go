package validparentheses

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type test_data struct {
	input string
	want  bool
}

func TestSolve(t *testing.T) {
	for index, data := range data() {
		result := Solve(data.input)
		if !cmp.Equal(data.want, result) {
			t.Fatalf(`Test #%d failed. Expected "%v", but got "%v"`, index+1, data.want, result)
		}
	}
}

func data() []test_data {
	return []test_data{
		test_data{
			input: "()",
			want:  true,
		},
		test_data{
			input: "()[]{}",
			want:  true,
		},
		test_data{
			input: "(]",
			want:  false,
		},
		test_data{
			input: "([)]",
			want:  false,
		},
	}

}
