package main

import (
	"errors"
	"fmt"
)

type Stack []interface {}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) IsEmpty() bool  {
	return len(stack) == 0
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack *Stack) Push(value interface{})  {
	*stack = append(*stack, value)
}

func (stack Stack) Top() (interface{}, error)  {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack) - 1], nil
}

func (stack *Stack) Pop() (interface{}, error)  {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack) - 1]
	*stack = theStack[:len(theStack) - 1]
	return value, nil
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	if len(heights)==0{
		return 0
	}
	heights = append(heights,0)
	stack := &Stack{}
	maxRec := heights[0]
	for idx,val := range heights{
		for top,_:=stack.Top();!stack.IsEmpty()&& heights[top.(int)]>=val;{
			h := heights[top.(int)]
			stack.Pop()
			length:=idx
			if !stack.IsEmpty(){
				top,_=stack.Top()
				length=idx-top.(int)-1
			}
			maxRec=max(maxRec, h * length)
		}
		stack.Push(idx)
	}

	return maxRec
}

//noStack
func getTop(stack []int) int {
	if len(stack)>0{
		return stack[len(stack)-1]
	}
	return -1
}

func largestRectangleArea2(heights []int) int {
	if len(heights)==0{
		return 0
	}
	heights = append(heights,0)
	heightStack := []int{}
	maxRec := heights[0]
	for idx,val := range heights{
		for top:=getTop(heightStack);len(heightStack)>0 && heights[top]>=val;{
			h := heights[top]
			heightStack=heightStack[:len(heightStack)-1]
			length:=idx
			if len(heightStack)>0{
				top=getTop(heightStack)
				length=idx-top-1
			}
			maxRec=max(maxRec, h * length)
		}
		heightStack=append(heightStack,idx)
	}

	return maxRec
}

/**
	数组中数字n代表 底为1高为n的矩形，并排排列后 求最大矩形面积
 */
func main()  {
	heights:=[]int{2,1,5,6,2,3}
	fmt.Println(largestRectangleArea(heights))
}