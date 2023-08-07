package maxprofit

import (
  "testing"
)


type TestData struct {
  input []int
  want int
}

func TestMaxProfit(t *testing.T) {

  tests := []TestData{
    {[]int{7,1,5,3,6,4}, 5},
    {[]int{7,6,4,3,1}, 0},
  }

  for _, test := range tests {
    if got := Solve(test.input); got != test.want {
      t.Errorf("maxProfit(%v) = %v", test.input, got)
    }
  }
}
