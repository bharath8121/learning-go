package main

import (
	"fmt"
	"math"
	"strings"
	"time"
	)


type Node struct{
	state [3][3]string
	row int
	column int
	score int
}

func EvaluateFunction(currentState [3][3]string)int{
	finish, winner := CheckForFinish(currentState)	
	X_number := [2]int{0, 0}
	O_number := [2]int{0, 0}
	if finish{
		if winner == "X"{
			return -1
		}
		if winner == "O"{
			return 1
		}
	} else{
		for row := range currentState{
			// Code to add strings in a row
			finalString := ""
			for col := range currentState[row]{
				finalString += currentState[row][col]
			}
			x_count := strings.Count(finalString, "X")
			o_count := strings.Count(finalString, "O")
			if x_count > 0{
				X_number[x_count-1]++
			}
			if o_count > 0{
				O_number[o_count-1]++
			}
		}
	}
	// From the AI Text Book
	return 3 * X_number[1] + X_number[0] - ( 3 * O_number[1] + O_number[0] )
}

// func ProductWithoutSelf(array []int){
// 	frontProduct := make([]int, len(array))
// 	backProduct := make([]int, len(array))
// }

func GenerateChildren(currentState [3][3]string, move string) []Node{
	var children []Node
	for i:=0; i<len(currentState); i++{
		for j:=0; j<len(currentState[0]); j++{
			if currentState[i][j] == ""{
				var temp [3][3]string
				for m:=0; m<len(currentState); m++{
					for n:=0; n<len(currentState[0]); n++{
						temp[m][n] = currentState[m][n]
					}
				}
				temp[i][j] = move
				newNode := Node{
					state: temp,
					row: i,
					column: j,
				}
				children = append(children, newNode)
			}
		}
	}
	return children
}

func Max(currentNode Node, depth int) (Node, Node){
	time.Sleep(10000 * time.Millisecond)
	var resultNode Node
	if depth == 2{
		currentNode.score = EvaluateFunction(currentNode.state)
	} else{
		children := GenerateChildren(currentNode.state, "O")
		fmt.Println("Possibilities: ", len(children))
		max := math.Inf(-8)
		for i := range children{
			children[i], _ = Min(children[i], depth+1)
			if max < float64(children[i].score){
				max = float64(children[i].score)
				resultNode = children[i]
			}
		}
		currentNode.score = int(max)
	}
	fmt.Println("Returning for state: ", currentNode.state, currentNode.score)
	return currentNode, resultNode
}

func Min(currentNode Node, depth int) (Node, Node){
	time.Sleep(10000 * time.Millisecond)
	var resultNode Node
	if depth == 2{
		currentNode.score = EvaluateFunction(currentNode.state)
	} else{
		children := GenerateChildren(currentNode.state, "X")
		fmt.Println("Possibilities: ", len(children))
		min := math.Inf(8)
		for i := range children{
			children[i], _ = Max(children[i], depth+1)
			min = math.Min(min, float64(children[i].score))
			if min > float64(children[i].score){
				min = float64(children[i].score)
				resultNode = children[i]
			}
		}
		currentNode.score = int(min)
	}
	fmt.Println("Returning for state: ", currentNode.state, currentNode.score)
	return currentNode, resultNode
}


func CheckForFinish(currentState [3][3]string) (bool, string){
	cross_check := ""
	rev_cross_check := ""
	remaining := 0
	for i := 0; i<3; i++{
		row_check := ""
		column_check := ""
		for j := 0; j<3; j++{
			if currentState[i][j] == ""{
				remaining++
			}
			if i == j{
				cross_check += currentState[i][j]
				rev_cross_check += currentState[i][2-j]
			}
			row_check += currentState[i][j]
			column_check += currentState[j][i]
		}
		if row_check == "XXX" || column_check == "XXX"{
			return true, "X"
		} else if column_check == "OOO" || row_check == "OOO"{
			return true, "O"
		}
	}
	if cross_check == "XXX" || rev_cross_check == "XXX"{
		return true, "X"
	} else if cross_check == "OOO" || rev_cross_check == "OOO"{
		return true, "O"
	} else if remaining > 0{
		return false, ""
	} else if remaining == 0{
		return true, ""
	}
	return false, ""
}

func GameHost(currentState [3][3]string){
	var row, column int
	for true{
		fmt.Scanf("%d\n", &row)
		fmt.Scanf("%d\n", &column)
		currentState[row][column] = "X"
		finish, winner := CheckForFinish(currentState)
		if finish{
			fmt.Printf("%s WON THE MATCH", winner)
		}

		currentNode := Node{
			state: currentState,
			row: row,
			column: column,
		}
		_, resultNode := Max(currentNode, 0)
		row, column = resultNode.row, resultNode.column
		currentState[row][column] = "O"
		fmt.Println(resultNode)
		time.Sleep(30 * time.Millisecond)
		finish, winner = CheckForFinish(currentState)
		if finish{
			fmt.Printf("%s WON THE MATCH", winner)
			break
		}
	}
	// fmt.Printf("%d %d", row, column)
}


func main(){
	var initialState [3][3]string
	GameHost(initialState)
	// finished, winner := CheckForFinish([3][3]string{{"X", "O", "O"}, {}, {"X"}})
	// if finished{
	// 	fmt.Println(winner)
	// } else{
	// 	fmt.Println("not won")
	// }
}