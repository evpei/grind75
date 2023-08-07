package invertbinarytree 

import "testing"

type TestData struct {
  input *TreeNode
  expected *TreeNode
}

func TestSolve(t *testing.T){

  tests := []TestData{
    {&TreeNode{Val: 4}, &TreeNode{Val: 0}},
  }
  for i, test := range tests {
    res := Solve(test.input)
    if res.Val != test.expected.Val {
      t.Fatalf("Test %d: Expected %d, got %d", i + 1, test.expected.Val, res.Val)
    }
  }
}

