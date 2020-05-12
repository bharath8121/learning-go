package main

import (
	"fmt"
	"strconv"
	"sort"
	"math"
	"strings"
	)


type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

type ListNode struct{
	Val int
	Next *ListNode
}

var decodeWaysMap map[string]string
var longestPallindromeSubSeq map[string]int = make(map[string]int)


func validateSubSequences(pushed, popped []int) bool{
	var stack []int
	j := 0
	for _, element := range(pushed){
		stack = append(stack, element)
		for len(stack) != 0 && stack[len(stack) - 1] == popped[j]{
			j++
			stack = stack[:len(stack) - 1]
		}
	}
	return len(popped) == j
}


func evalRPN(tokens []string) int{
	var stack []int
	for _, element := range(tokens){
		if n, err := strconv.Atoi(element); err == nil{
			stack = append(stack, n)
		} else{
			n1, n2, n3 := stack[len(stack) - 1], stack[len(stack) - 2], 0
			
			stack = stack[:len(stack) - 2]
			switch element {
			case "*":
				n3 = n2 * n1
			case "-":
				n3 = n2 - n1
			case "/":
				n3 = n2 / n1
			case "+":
				n3 = n2 + n1
			}
			stack = append(stack, n3)
		}
	}
	return stack[0]
}


func IsPresent(target string, list []string) bool{
	for _, b := range(list){
		if target == b{
			return true
		}
	}
	return false
}


func TrappingWater(height []int) int{
	var stack [][2]int
	amount := 0
	for i := range(height){
		if len(stack)!=0 && stack[0][1] < height[i]{
			min_tower := 0
			if height[i] < stack[0][1]{
				min_tower = height[i]
			} else{
				min_tower = stack[0][1]
			}
			total_coverage := min_tower * ( i - stack[0][0] - 1 )
			t := len(stack) - 1
			for t > 0 {
				total_coverage -= stack[t][1]
				t--
			}
			if stack[0][1] < height[i]{
				stack = [][2]int{}
			}
			amount += total_coverage
		}
		stack = append(stack, [2]int{i, height[i]})
		fmt.Println(stack)
		fmt.Println(amount)
	}
	fmt.Println(stack)
	return amount;
}

func LadderLength(beginWord, endWord string, wordList []string) int{
	var wordListMap map[string]int = make(map[string]int)
	var invertedIndex map[int][]string = make(map[int][]string)
	var level int = 0
	for _, word := range(wordList){
		for j := range(word){
			if !IsPresent(string([]byte{word[j]}), invertedIndex[j]){
				invertedIndex[j] = append(invertedIndex[j], string([]byte{word[j]}))
			}
		}
		wordListMap[word] = 1
	}

	fmt.Println(wordListMap)
	fmt.Println(invertedIndex)

	queue := []string{beginWord}
	var node string
	tempQueue, closedNodes := []string{}, []string{beginWord}

	for len(queue) != 0{
		for len(queue) != 0{
			node, queue = queue[0], queue[1:]
			for indx := range(node){
				for _, s := range(invertedIndex[indx]){
					temp := node[:indx]+s+node[indx+1:]
					
					if temp == endWord{
						return level+2
					}
					if _, ok := wordListMap[temp]; ok && !IsPresent(temp, closedNodes){
						closedNodes = append(closedNodes, temp)
						tempQueue = append(tempQueue, temp)
					}
				}
			}
		}
		queue = tempQueue
		tempQueue = []string{}
		level++
	}
	
	

	return 0;

}


func transferElements(indx int, nums []int) []int{
	newNums := []int{}
	for i := range(nums){
		if i != indx{
			newNums = append(newNums, nums[i])
		}
	}
	return newNums
}

func MaxCoins(nums []int) int{
	coins := 0
	fmt.Println(nums)

	if len(nums) == 1{
		return nums[0]
	}
	
	for indx := range(nums){
		val := nums[indx]
		if indx - 1 >= 0{
			val = val * nums[indx-1]
		}
		if indx + 1 <= len(nums)-1{
			val = val * nums[indx+1]
		}

		
		
		val += MaxCoins(transferElements(indx, nums))
		if coins < val{
			coins = val
		}
	}
	return coins
}

// This is not the actual solution. I understood the problem wrongly.
func LengthLongestPath(input string) int{
	lines := strings.Split(input, "\n")
	max := 0
	for i := range lines{
		max = int(math.Max(float64(strings.Count(lines[i], "\t")), float64(max)))
	}
	return max
}

func GenChildRecord(nodes, values *[]string, currLength int) []string{
	var storeMap map[string]int = make(map[string]int)
	var results []string
	var node string
	for len((*nodes)[0]) < currLength{
		node, *nodes = (*nodes)[0], (*nodes)[1:]
		for j := range (*values){
			insert := node + (*values)[j]
			if _, ok := storeMap[insert]; !ok  && strings.Count(insert, "A") < 2  &&  !strings.Contains(insert, "LLL"){
				*nodes = append((*nodes), insert)
				storeMap[insert] = 1
			}
		}
	}
	return results
}


func CheckRecord(n int) int{
	start := 0
	var values []string = []string{"A", "L", "P"}
	var results []string = []string{"A", "L", "P"}
	for start < n-1{
		GenChildRecord(&results, &values, start+2)
		start++
	}
	fmt.Println(results)
	return len(results)
}

func SplitListToParts(root *ListNode, k int) []*ListNode{
	counter := root
	length := 0
	var elements []int
	var result []*ListNode
	var temp *ListNode
	for counter != nil{
		length++
		counter = counter.Next
	}
	width, extras := length/k, length%k
	i := 0
	for i < k{
		elements = append(elements, width + int(math.Min(float64(1), float64(extras))))
		if extras > 0{ extras-- }
		i++
	}
	curr := root
	for i := range elements{
		result = append(result, curr)
		for elements[i] > 1{
			curr = curr.Next
			elements[i]--
		}
        if curr != nil{ 
            temp, curr = curr, curr.Next 
            temp.Next = nil
        }
	}
	return result
}

func TotalHammingDistance(nums []int)int{
	hammingDistance := 0
	for i := 0; i < len(nums)-1; i++{
		for j := i+1; j < len(nums); j++{
			xored := nums[i] ^ nums[j]
			for xored > 0{
				hammingDistance += xored & 1
				xored >>= 1
			}
		}
	}
	return hammingDistance
}

func PracticeGo(){
	decodeWaysMap = make(map[string]string)
	decodeWaysMap["bharath"] = "chandra"
	if decodeWaysMap["praneeth"] == ""{
		fmt.Println(decodeWaysMap["bharath"])
	}
}

func PruneTree(root *TreeNode) *TreeNode{
	if root == nil{
		return nil
	}
	root.Left = PruneTree(root.Left)
	root.Right = PruneTree(root.Right)
	if root.Left != nil || root.Right != nil || root.Val == 1{
		return root
	}
	return nil
}

func CountTriplets(A []int)int{
	count := 0
	for i:=0; i<len(A); i++{
		if A[i] == 0{
			count += 2 * (len(A))
			continue
		}
		for j:=0; j<len(A); j++{
			if A[i]&A[j] == 0{
				count += len(A)
				continue
			}
			for k:=0; k<len(A); k++{
				if A[i]&A[j]&A[k] == 0{
					count++
				}
			}
		}
	}
	return count
}

// func GenLadderChildren(word string, lookUpMap map[string]int) []string{
// 	var children []string

// 	for i := range word{
// 		for 
// 	}
// }

// func _FindLadders(word, endWord string, lookUpMap map[string]int, parent *[]string, result *[][]string){
// 	if word == endWord{
// 		fmt.Println("found the word")
// 		parent = append(parent, word)
// 		result = result(append, parent)
// 	}
// 	else{

// 	}
// }

// func FindLadders(beginWord string, endWord string, wordList []string)[][]string{
// 	var lookUpMap map[string]int = make(map[string]int)
// 	for i := range wordList{
// 		lookUpMap[wordList[i]] = 1
// 	}
// 	result := make([][]string)
// 	_FindLadders(beginWord, endWord, lookUpMap, &result)
// 	return result

// }
 


// Time limit exceeded
func LongestPallindromeSubSeq(s string)int{
	if (len(s) == 1){
		return 1
	}
	if (len(s) == 0){
		return 0
	}
	if val, ok := longestPallindromeSubSeq[s]; ok{
		return val
	}
	if s[0] == s[len(s)-1]{
		longestPallindromeSubSeq[s] = 2 + LongestPallindromeSubSeq(s[1:len(s)-1])
	} else{
		longestPallindromeSubSeq[s] = int(math.Max( float64(LongestPallindromeSubSeq(s[0:len(s)-1])), float64(LongestPallindromeSubSeq(s[1:len(s)])) ))
	}
	return longestPallindromeSubSeq[s]
}

// func ThirdMax(array []int){
	
// }



// func FindSubstringInWraproundString(p string) int{
// 	alphabets := "abcdefghijklmnopqrstuvwxyz"
// 	count, i, pointer, temp_count := 0, 0, 0, 0
// 	closedElements := make(map[string]int)
// 	for i < len(p){
// 		if alphabets[pointer] != p[i]{
// 			pointer = int(p[i]) - int('a')
// 			count += (temp_count * (temp_count + 1)) / 2
// 			temp_count = 0
// 		}
// 		if alphabets[pointer] == p[i]{
// 			pointer = (pointer + 1) % 26
// 			temp_count++
// 		}
// 		i++
// 	}
// 	return count + (temp_count * (temp_count + 1)) / 2
// }

// func Calculator(s string) int{

// }

// 784 question
// func HelperLetterCasePermutation(S string, i int) []string{
// 	S = S[0].upper
// 	HelperLetterCasePermutation()
// }

// func LetterCasePermutation(S string) []string{

// }


func NextLargestNodes(head *ListNode) []int{
	var listElements []int
	temp := head
	for temp != nil{
		listElements = append(listElements, temp.Val)
		temp = temp.Next
	}

	i := 1
	var stack []int = []int{listElements[0]}
	top := 0
	for i < len(listElements){
		if top > -1{
			if listElements[stack[top]] > listElements[i]{
				top++
				stack[top] = i
				i++
			} else{
				if listElements[stack[top]] < listElements[i]{
					listElements[stack[top]] = listElements[i]
					top--
				}
			}
		}
	}
	return listElements

}


// func BuildTree(node *TreeNode, array []int){
// 	if node.Left && node.Right{
// 		BuildTree(node.Left, array)
// 		BuildTree(node.Right, array)
// 	}
// }

// func GenerateTrees(n int)[] *TreeNode{
	
// }

// func MaxHeapWithArrays()

// 713 - easy problem look at code
func NumSubarrayProductLessThanK(nums []int, k int) int{
	start, end, result, product := 0, 1, 0, 1
	if len(nums) > 0{
		product = nums[0]
	}
	for end < len(nums){
		fmt.Println("start: ", nums[start])
		fmt.Println("end: ", nums[end])
		if product >= k{
			product /= nums[start]
			start++
		} else{
			if nums[end] < k{
				result++
			}
			product *= nums[end]
			result++
			end++
		}
	
		fmt.Println("reuslt: ", result)
	}
	for start < end-1{
		if product < k{
			result++
		}
		product /= nums[start]
		start++
	}
	return result
}

func FrequencySorts(s string) string{
	var Map map[byte]int = make(map[byte]int)
	for i := range s{
		Map[s[i]]++
	}
	var temp [][]int
	for i := range Map{
		temp = append(temp, []int{Map[i], int(i)})
	}
	sort.Slice(temp, func(i, j int) bool{
		if temp[i][0] == temp[j][0]{
			return temp[i][1] < temp[j][1]
		}
		return temp[i][0] > temp[j][0]
	})
	// fmt.Println(temp)
	var result string = ""
	for j := 0; j < len(temp) ; j++{
		for temp[j][0] > 0{
			result += string(temp[j][1])
			temp[j][0]--
		}
		
	}
	return result
	// var InvertMap map[int][]string
}

func TopKElements(nums []int, k int)[]int{
	var Map map[int]int = make(map[int]int)
	for i := range nums{
		Map[nums[i]]++
	}
	var temp [][]int
	for i := range Map{
		temp = append(temp, []int{Map[i], i})
	}
	sort.Slice(temp, func(i, j int) bool{
		if temp[i][0] == temp[j][0]{
			return temp[i][1] < temp[j][1]
		}
		return temp[i][0] > temp[j][0]
	})
	var result []int
	for i := 0; i < k; i++{
		result = append(result, temp[i][1])
	}
	return result
}

func DecodeWays(input string) int{
	if input == ""{
		// fmt.Println("returning 1")
		return 1
	}
	if input[0] != '0'{
		i := 0
		j := 0
		// fmt.Printf("Considering %c",input[0])
		i = DecodeWays(input[1:])
		if len(input) >= 2{
			// fmt.Printf("Considering ",input[:2])
			num, err := strconv.Atoi(input[:2])
			if err != nil{
				return 0
			}
			if num <= 26 && num >= 10{
				j = DecodeWays(input[2:])
			}
		}
		// fmt.Printf("%d\n", i+j)
		return i + j
	}
	return 0
}

func LongestParanthesis(input string) int{
	// var paran_stack map[string]
	return 0
}

func UniqueMorseRepresentations(words []string) int{
	uniqueelements := map[string]int{}
	representation := [26]string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	// fmt.Println(representation[0])
	for word := range words{
		var temp string = ""
		for char := range words[word]{
			temp += representation[int(words[word][char])-97]
		}
		uniqueelements[temp] = 1
	}
	return len(uniqueelements)
}

func MergeIntervals(matrix [][]int) [][]int{
	sort.Slice(matrix[:], func(i int, j int) bool{
		for x := range matrix[i]{
			if matrix[i][x] == matrix[j][x]{
				continue
			}
			return matrix[i][x] < matrix[j][x]
		}
		return false
	})
	start := 1
	curr := matrix[0]
	var final [][]int = make([][]int, 0)
	for start < len(matrix){
		if curr[1] >= matrix[start][0]{	
			curr[1] = int(math.Max(float64(curr[1]), float64(matrix[start][1])))
		} else{
			final = append(final, curr)
			curr = matrix[start]
		}
		start++
	}
	final = append(final, curr)
	return final
}

func PartitionLabels(input string) []int{
	var results []int
	var lastIndex map[byte]int = make(map[byte]int)
	for index := range input{
		lastIndex[input[index]] = index
	}
	// fmt.Println(lastIndex)
	start := 0
	curr := start
	end := lastIndex[input[curr]]
	for curr < len(input){
		// fmt.Printf("at: %s\n", string(input[curr]))
		end = int(math.Max(float64(end), float64(lastIndex[input[curr]])))
		// fmt.Printf("curr: %d, end: %d\n", curr,end)
		if curr ==  end{
			// fmt.Printf("Adding: %s\n", string(input[start:curr+1]))
			results = append(results, curr-start+1)
			curr++
			start = curr
		} else{
			curr++
		}		
	}
	return results	
}

// func SmallestRangeII(input []int, k int) []int{
// 	return 0
// }

func main(){
	// decodeWaysMap = make(map[string]string)	
	// number := "9371597631128776948387197132267188677349946742344217846154932859125134924241649584251978418763151253"
	// fmt.Println(DecodeWays(number))
	// PracticeGo()
	// var matrix [][]int = [][]int{
	// 					{1,4},
	// 					{4,5},
	// 				}
	// fmt.Println(MergeIntervals(matrix))
	// fmt.Println(PartitionLabels("ababcbacadefegdehijhklij"))
	
	// fmt.Println(UniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
	// fmt.Println(TopKElements([]int{1,1,2,3,4,4,1,3,5}, 2))
	// fmt.Println(FrequencySorts("Aabb"))
	// fmt.Println(NumSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	// fmt.Println(CountTriplets([]int{2,1,3}))
	// fmt.Println(LongestPallindromeSubSeq("euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew"))
	// fmt.Println(TotalHammingDistance([]int{4, 14, 2}))
	// fmt.Println(FindSubstringInWraproundString("abab"))
	// fmt.Println(CheckRecord(20))
	// fmt.Println(NextLargestNodes())

	// fmt.Println(LadderLength("teach", "place", []string{"peale","wilts","place","fetch","purer","pooch","peace","poach","berra","teach","rheum","peach"}))

	// fmt.Println(MaxCoins([]int{3,1,5,8}))
	fmt.Println(TrappingWater([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}