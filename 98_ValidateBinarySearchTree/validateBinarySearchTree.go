package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	return isValid(root, nil, nil)
}

func isValid(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= *min {
		return false
	}
	if max != nil && root.Val >= *max {
		return false
	}

	return isValid(root.Left, min, &(root.Val)) && isValid(root.Right, &(root.Val), max)
}

func main() {
	leftNode := TreeNode{
		Val: 1,
	}
	rightNode := TreeNode{
		Val: 3,
	}
	root := TreeNode{
		Val:   2,
		Left:  &leftNode,
		Right: &rightNode,
	}

	println(isValidBST(&root))
}
