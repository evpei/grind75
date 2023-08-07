package mergetwo

import (
	"testing"

	ds "github.com/evpei/grind-75/src/kit/datastructures"
	"github.com/google/go-cmp/cmp"
)

type test_data struct {
	list1 ds.LinkedList[int]
	list2 ds.LinkedList[int]
	want  ds.LinkedList[int]
}

func TestSolve(t *testing.T) {
	for index, data := range data() {
		result := Solve(data.list1, data.list2).Head
		testData := data.want.Head

		for testData != nil || result != nil {
			if result == nil {
				t.Fatalf("Result is nil")
			}
			if testData == nil {
				t.Fatalf("Result is not nil")
			}
			if !cmp.Equal(testData.Val, result.Val) {
				t.Fatalf(`Test #%d failed. Expected "%v", but got "%v"`, index+1, testData.Val, result.Val)
			}
			result = result.Next
			testData = testData.Next
		}
	}
}

func data() []test_data {
	return []test_data{
		{
			list1: ds.LinkedList[int]{}.New(1, 2, 4),
			list2: ds.LinkedList[int]{}.New(1, 3, 4),
			want:  ds.LinkedList[int]{}.New(1, 1, 2, 3, 4, 4),
		},
		{
			list1: ds.LinkedList[int]{},
			list2: ds.LinkedList[int]{},
			want:  ds.LinkedList[int]{},
		},
	}
}
