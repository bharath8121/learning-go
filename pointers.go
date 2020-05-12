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

func del_node(data int, head **Node){
	var ptp **Node = head
	for *ptp != nil{
		if((*ptp).data == data){
			*ptp = (*ptp).next
		} else {
			ptp = &(*ptp).next
		}
	}
}


func main(){
	var node *Node = new_node(10)
	var temp *Node = node
	for i:=2; i<10; i++{
		temp.next = new_node(i*10)
		temp = temp.next
	}

	temp = node
	var del_data int
	fmt.Printf("enter data to delete")
	fmt.Scanln(&del_data)
	del_node(del_data, &temp)
	for temp!=nil{
		fmt.Println(temp.data)
		temp = temp.next
	}
}

