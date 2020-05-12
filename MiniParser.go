package main

import (
	"fmt"
)


type NestedInteger struct{

}

func Deserialize(s string) *NestedInteger{
	var stack []NestedInteger = []NestedInteger{new(NestedInteger)}
	var tempNode NestedInteger
	for indx, element := range(s){
		switch element{
		case "[":
			stack = append(stack, new(NestedInteger))
		case "]":
			tempNode, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack[len(stack)-1].add(tempNode)
		case ",":
			tempNode, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack[len(stack)-1].add(tempNode)
		}
	}
}


func main(){

}