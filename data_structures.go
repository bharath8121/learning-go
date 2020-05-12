package main

import (
	"fmt"
	)


type Stack struct{
	stack []int
	top int
}


func StackPush(stack *Stack, element int){
	stack.top++
	stack.stack = append(stack.stack, element)
}

func StackPop(stack *Stack) int{
	var element int
	element, stack.stack = stack.stack[stack.top], stack.stack[:stack.top]
	stack.top--
	return element	
}


func main(){

	stack := Stack{
		stack: []int{},
		top: -1,
	}

	StackPush(&stack, 5)
	fmt.Printf("%d\n", stack.top)
	fmt.Println(stack.stack)
	fmt.Println("Deleted: ", StackPop(&stack))
	fmt.Println(stack.stack)
	fmt.Println(stack.top)
}