package main

import (
	"fmt"
	"container/heap"
	"time"
		)

type State struct{
	state [][]int
	parent *State
}

type Element struct{
	node *State
	moves int
	cost int
	index int
}

type Heap []*Element



// Heap implementation

func (h Heap) Len() int{
	return len(h)
}

func CalculateHofX(state [][]int) int{
	hLength := 0
	for m := 0; m < len(state)-1; m++{
		hLength += len(state[m])
	}
	return hLength
}

func (h Heap) Less(i, j int) bool{
	return h[i].cost < h[j].cost	
}

func (h Heap) Swap(i, j int){
	// fmt.Printf("swapping: %s and %s\n", h[i].word, h[j].word)
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *Heap) Push(x interface{}){
	// fmt.Println("Pushing")
	element := x.(*Element)
	element.index = len(*h)
	*h = append(*h, element)
	heap.Fix(h, element.index)
}

func (h *Heap) Pop() interface{}{
	// fmt.Println("Popping")
	element := (*h)[len(*h)-1]
	element.index = -1
	*h = (*h)[0:len(*h)-1]
	return element
}

func (h *Heap) update(element *Element, node *State){
	// fmt.Println("Updating")
	element.node = node
	heap.Fix(h, element.index)
}


func CompareStates(state1, state2 [][]int) bool{
	for i := range state1{
		if len(state1[i]) == len(state2[i]){
			for j := range state1[i]{
				if state1[i][j] != state2[i][j]{
					return false
				}
			}
		} else{
			return false
		}
	}
	return true
}

func StateClosed(state *State, closedNodes []*State) bool{
	for i := range closedNodes{
		if CompareStates(closedNodes[i].state, state.state){
			return true
		}
	}
	return false
}

func GenChildren(currentNode *State, currentNodes []*State) []*State{
	
	// fmt.Println("Current Node: ", currentNode)
	for i := 0; i<len(currentNode.state); i++{
		for j:=0; j<len(currentNode.state); j++{
			// fmt.Println(i,j)
			var tmpNode State
			for m := 0; m < len(currentNode.state); m++{
				var t []int
				for n := 0; n < len(currentNode.state[m]); n++{
					t = append(t, currentNode.state[m][n])
				}
				tmpNode.state = append(tmpNode.state, t)
			}
			tmpNode.parent = currentNode
			// fmt.Println(tmpNode)
			var element int
			if len(currentNode.state[i]) != 0 && (len(currentNode.state[j]) == 0 || currentNode.state[i][len(currentNode.state[i])-1] < currentNode.state[j][len(currentNode.state[j])-1]){
				element, tmpNode.state[i] = tmpNode.state[i][len(tmpNode.state[i])-1], tmpNode.state[i][:len(tmpNode.state[i])-1]
				tmpNode.state[j] = append(tmpNode.state[j], element)
				currentNodes = append(currentNodes, &tmpNode)
				// fmt.Println("Temp Node: ", tmpNode)
				// fmt.Println("Current Node: ", currentNode)
			}
		}
	}
	return currentNodes
}

func BFSTowersofHanoi(finalState State, currentNodes []*State) *State{
	var currentNode *State
	var closedNodes []*State
	nodesTraversed := 0
	for len(currentNodes) > 0{
		nodesTraversed += 1
		currentNode, currentNodes = currentNodes[0], currentNodes[1:]
		if CompareStates(finalState.state, currentNode.state){
			// fmt.Println("Reached the state: ", currentNode)
			// fmt.Println("Current Nodes: ", currentNodes)
			fmt.Println("Number of Nodes Traversed: ", nodesTraversed)
			return currentNode
		}
		if StateClosed(currentNode, closedNodes){
			continue
		} else{
			closedNodes = append(closedNodes, currentNode)
		}
		// fmt.Println("Current Node: ", currentNode)
		currentNodes = GenChildren(currentNode, currentNodes)
		// break
	}
	// fmt.Println("Current Nodes: ", currentNodes)
	return nil
}



func CheckIntersection(forwardCurrentNodes, forwardClosedNodes, backwardCurrentNodes, backwardClosedNodes []*State) [2]*State{
	// Calculating Forward Current Nodes => Backward Closed Nodes
	for i := range forwardCurrentNodes{
		for j := range backwardClosedNodes{
			if CompareStates(forwardCurrentNodes[i].state, backwardClosedNodes[j].state){
				return [2]*State{forwardCurrentNodes[i], backwardClosedNodes[j]}
			}
		}
	}
	for i := range backwardCurrentNodes{
		for j := range forwardClosedNodes{
			if CompareStates(backwardCurrentNodes[i].state, forwardClosedNodes[j].state){
				return [2]*State{forwardClosedNodes[j], backwardCurrentNodes[i]}
			}
		}
	}
	// Calculating Backward Current Nodes => Foward Closed Nodes
	return [2]*State{}
}

func BiDirectionalTowersofHanoi(forwardCurrentNodes, backwardCurrentNodes []*State, finalState, initialState State) *State{
	var forwardClosedNodes []*State
	var backwardClosedNodes []*State
	var currentBackwardNode *State
	var currentForwardNode *State
	numNodes := 0

	for true{
		numNodes += 1
		equalNodes := CheckIntersection(forwardCurrentNodes, forwardClosedNodes, backwardCurrentNodes, backwardClosedNodes)
		// fmt.Println("one")
		if equalNodes[0] != nil && equalNodes[1] != nil{
			// fmt.Println(equalNodes[0].state, equalNodes[1].state)
			newParent, child := equalNodes[0], equalNodes[1].parent
			for child != nil{
				temp := child.parent
				child.parent = newParent
				newParent = child
				child = temp
			}
			// fmt.Println(newParent.state)
			// fmt.Println("found similar states")
			fmt.Println("Number of Nodes Traversed: ", numNodes*2)
			return newParent
		}
		currentForwardNode, forwardCurrentNodes = forwardCurrentNodes[0], forwardCurrentNodes[1:]
		if CompareStates(finalState.state, currentForwardNode.state){
			// fmt.Println("Reached the state: ", currentNode)
			// fmt.Println("Current Nodes: ", currentNodes)
			return currentForwardNode
		}
		if !StateClosed(currentForwardNode, forwardClosedNodes){
			forwardClosedNodes = append(forwardClosedNodes, currentForwardNode)
			forwardCurrentNodes = GenChildren(currentForwardNode, forwardCurrentNodes)
		}

		currentBackwardNode, backwardCurrentNodes = backwardCurrentNodes[0], backwardCurrentNodes[1:]
		if CompareStates(initialState.state, currentBackwardNode.state){
			return currentBackwardNode
		}
		if !StateClosed(currentBackwardNode, backwardClosedNodes){
			backwardClosedNodes = append(backwardClosedNodes, currentBackwardNode)
			backwardCurrentNodes = GenChildren(currentBackwardNode, backwardCurrentNodes)
		}

	}
	return nil
}

func AStarTowersofHanoi(h Heap, finalState State) *State{
	var currentNodes []*State
	var closedNodes []*State
	numNodes := 0
	for len(h) > 0{
		currentElement := heap.Pop(&h).(*Element)
		currentNode := currentElement.node
		numNodes += 1
		if CompareStates(finalState.state, currentNode.state){
			// fmt.Println("Found the path")
			// fmt.Println(currentNode)
			fmt.Println("Number of Nodes: ", numNodes)
			return currentNode
		}

		if StateClosed(currentNode, closedNodes){
			continue
		} else{
			closedNodes = append(closedNodes, currentNode)
		}

		// fmt.Println(currentNode)
		currentNodes := GenChildren(currentNode, currentNodes)
		for i := range currentNodes{
			newElement := Element{
				node: currentNodes[i],
				moves: currentElement.moves + 1,
				cost: CalculateHofX(currentNodes[i].state) + currentElement.moves + 1, 
			}
			heap.Push(&h, &newElement)
		}
	}
	return nil	
}

func GetInitialState(numTowers, numRings int) State{
	var initialState State
	var temp []int
	for i := numRings; i>0; i--{
		temp = append(temp, i)
	}
	initialState.state = append(initialState.state, temp)
	for i := 1; i < numTowers; i++{
		initialState.state = append(initialState.state, make([]int, 0))
	}
	initialState.parent = nil
	return initialState
}

func GetFinalState(numTowers, numRings int) State{
	var finalState State
	for i := 0; i < numTowers; i++{
		finalState.state = append(finalState.state, make([]int, 0))
	}
	for i := numRings; i > 0; i--{
		finalState.state[len(finalState.state)-1] = append(finalState.state[len(finalState.state)-1], i)
	}
	return finalState
}

func PrintResult(result *State){
	moves := -1
	for result != nil{
		fmt.Println(result.state)
		moves += 1
		result = result.parent
	}
	fmt.Println("Total number of moves: ", moves)
}

func HandleBFSTowersofHanoi(numTowers, numRings int){
	fmt.Println("BFS ALGORITHM")
	start := time.Now().Unix()
	var currentNodes []*State
	initialState := GetInitialState(numTowers, numRings)
	currentNodes = append(currentNodes, &initialState)
	// fmt.Println("Current Queue: ", currentNodes)
	finalState := GetFinalState(numTowers, numRings)
	// fmt.Println(TowersofHanoi(finalState))
	result := BFSTowersofHanoi(finalState, currentNodes)
	PrintResult(result)
	fmt.Println("Total Time Taken: ", time.Now().Unix() - start)
}

func HandleBiDirectionalTowersofHanoi(numTowers, numRings int){
	fmt.Println("BI-DIRECTIONAL ALGORITHM")
	start := time.Now().Unix()
	var forwardCurrentNodes []*State
	var backwardCurrentNodes []*State

	initialState := GetInitialState(numTowers, numRings)
	finalState := GetFinalState(numTowers, numRings)

	forwardCurrentNodes = append(forwardCurrentNodes, &initialState)
	backwardCurrentNodes = append(backwardCurrentNodes, &finalState)

	result := BiDirectionalTowersofHanoi(forwardCurrentNodes, backwardCurrentNodes, finalState, initialState)
	PrintResult(result)
	fmt.Println("Total Time Taken: ", time.Now().Unix() - start)

}

func HandleAStartTowersofHanoi(numTowers, numRings int){
	fmt.Println("A-STAR ALGORITHM")
	start := time.Now().Unix()
	var h Heap

	initialState := GetInitialState(numTowers, numRings)
	finalState := GetFinalState(numTowers, numRings)
	newElement := Element{
		node: &initialState,
		moves: 0,
		cost: CalculateHofX(initialState.state),
	}
	h = append(h, &newElement)
	heap.Init(&h)

	result := AStarTowersofHanoi(h, finalState)
	PrintResult(result)
	fmt.Println("Total Time Taken: ", time.Now().Unix() - start)

}

func main(){
	var numTowers int
	var numRings int
	fmt.Scanf("%d\n", &numTowers)
	fmt.Scanf("%d\n", &numRings)
	HandleBFSTowersofHanoi(numTowers, numRings)
	HandleBiDirectionalTowersofHanoi(numTowers, numRings)
	HandleAStartTowersofHanoi(numTowers, numRings)
	
}