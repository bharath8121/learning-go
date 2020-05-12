package main

import "fmt"

type Node struct{
	data int
	next *Node
}

func new_node(data int) *Node{
	newNode := new(Node)
	newNode.data = data
	return newNode
}

func RemoveDuplicates(head *Node){
	for head.next != nil{
		if head.next.data == head.data {
			head.next = head.next.next
		} else {
			head = head.next
		}
	}
}

func NthNodeFromEnd(head *Node, n int) int{
	element, end := head, head

	for i:=0; i<n-1; i++{
		end = end.next
	}

	for end.next != nil{
		element = element.next
		end = end.next
	}
	return element.data
}

// func CounterClockWise(head *Node, k int){
// 	temp := head
// 	for i:=0; i<k; i++{
// 		temp := temp.next
// 	}
// 	end := temp
// 	for end.next != nil{
// 		end := end.next
// 	}
// 	end.next = head
// 	temp.next = nil
// }

func SwapElements(head *Node) *Node{
	dummy := new_node(-1)
	dummy.next = head
	ptr1, ptr2, ptr3 := dummy, dummy.next, dummy.next.next
	for ptr3 != nil{
		ptr1.next = ptr3
		ptr2.next = ptr3.next
		ptr3.next = ptr2
		ptr3, ptr2 = ptr2, ptr3
		fmt.Println("after swap")
		fmt.Println(ptr1.data,ptr2.data,ptr3.data)
		if(ptr3.next != nil && ptr3.next.next != nil){
			ptr1 = ptr1.next.next
			ptr2 = ptr2.next.next
			ptr3 = ptr3.next.next
			fmt.Println("after shift")
			fmt.Println(ptr1.data,ptr2.data,ptr3.data)
		}else{
			break
		}
	}
	return dummy
}

func main(){
	var node *Node = new_node(10)
	var temp *Node = node
	for i:=2; i<10; i++{
		temp.next = new_node(i*10)
		temp = temp.next
	}
	// RemoveDuplicates(temp)
	// fmt.Println(NthNodeFromEnd(node, 3))
	dummy := SwapElements(node)

	temp = dummy.next
	for temp != nil{
		fmt.Println(temp.data)
		temp = temp.next
	}
}