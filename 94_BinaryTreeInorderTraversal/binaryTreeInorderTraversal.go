package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type IntStack struct {
	List []*TreeNode
}

func (stack *IntStack) Push(val *TreeNode) {
	stack.List = append(stack.List, val)
}

func (stack *IntStack) Pop() *TreeNode {
	reVal := stack.List[len(stack.List)-1]
	stack.List = stack.List[:len(stack.List)-1]

	return reVal
}

var reList []int

func inorderTraversal(root *TreeNode) []int {

	getInorderTraversal(root)

	return reList
}

func getInorderTraversal(root *TreeNode) {
	if root != nil {
		getInorderTraversal(root.Left)
		reList = append(reList, root.Val)
		getInorderTraversal(root.Right)
	}
}

func inorderTraversal1(root *TreeNode) []int {
	reIntList := []int{}
	nodeStack := IntStack{}
	for root != nil || len(nodeStack.List) > 0 {
		for root != nil {
			nodeStack.Push(root)
			root = root.Left
		}

		if len(nodeStack.List) > 0 {
			root = nodeStack.Pop()
			reIntList = append(reIntList, root.Val)
			root = root.Right
		}
	}

	return reIntList
}

func main() {
	leftNode := TreeNode{
		Val: 2,
	}
	rightNode := TreeNode{
		Val: 3,
	}
	root := TreeNode{
		Val:   1,
		Left:  &leftNode,
		Right: &rightNode,
	}

	fmt.Println(inorderTraversal(&root))

	fmt.Println(inorderTraversal1(&root))
}
