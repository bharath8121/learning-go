package main

import (
		"fmt"
		"math"
	)

func fountain(array []int, index int, mostRecent int, covered int) float64{
	
	fmt.Printf("At index: %d and Most Recent: %d\n", index, mostRecent)

	if index >= len(array){
		fmt.Printf("Returning Inf\n")
		return math.Inf(-8)
	}

	if index == len(array)-1{
		if mostRecent == index{
			if covered + 1 == len(array){
				fmt.Println("Returning 1")
				return 1
			} else{
				fmt.Println("Returning Infinite")
				return math.Inf(-8)
			}
		} else if mostRecent < index{
			if covered + int(math.Min(float64(index - mostRecent), float64(array[index]))) + array[index] + 1 == len(array){
				fmt.Println("Node Covered: ", covered + int(math.Min(float64(index - mostRecent), float64(array[index]))) + array[index] + 1)
				return 1
			} else{
				fmt.Println("Returning Infinite")
				return math.Inf(-8)
			}
		}
		fmt.Println("Returning Infinite")
		return math.Inf(-8)
	}

	// Calculating covered = backwards + forwards
	newCovered := covered + int(math.Min(float64(index - mostRecent), float64(array[index]))) + array[index] + 1
	fmt.Println("Node Covered: ", newCovered)

	result := math.Min( 1 + fountain(array, index+array[index]+1, index+array[index]+1, newCovered), fountain(array, index+1, index, covered) )
	fmt.Println("Result: ",result)
	return result
}

func main(){
	fmt.Println(fountain([]int{4,1,0,0,0,5}, 0, 0, 0))
}