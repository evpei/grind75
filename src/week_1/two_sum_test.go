package week1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type test_data struct {
	nums   []int
	target int
	want   []int
}

func TestSolve(t *testing.T) {
	for index, data := range data() {
		result := Solve(data.nums, data.target)
		if !cmp.Equal(data.want, result) {
			t.Fatalf(`Test #%d failed. Expected "%v", but got "%v"`, index+1, data.want, result)
		}
	}
}

func data() []test_data {
	return []test_data{
		test_data{
			nums:   []int{2, 5, 5, 11},
			target: 10,
			want:   []int{1, 2},
		},
		test_data{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		test_data{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		test_data{
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

}
