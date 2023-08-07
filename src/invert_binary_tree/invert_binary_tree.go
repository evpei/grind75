package invertbinarytree

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func Solve(root *TreeNode) *TreeNode {
  return &TreeNode{Val: -1}
}
