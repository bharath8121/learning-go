package main

import (
	"fmt"
	"sort"
)

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

var smallestFromLeaf string = string('z'+1)

// func BSTFromPreorder(preorder []int) *TreeNode{
	
// }

func ProfitableSchemesRecursion(G, p, P, i int, schemes *int){
	if G == 0 && p >= P{

	}
}

func ProfitableSchemes(G int, P int, group []int, profit []int) int{
	var schemes int = 0

}

func DFSDiameterOfBinaryTree(root *TreeNode, max *int) int{
	if root == nil{
		return 0
	}
	left, right := 0, 0
	left = DFSDiameterOfBinaryTree(root.Left, max)
	right = DFSDiameterOfBinaryTree(root.Right, max)
	*max = int(math.Max(float64(*max), float64(left+right+1)))
	return int(math.Max(float64(left), float64(right))) + 1
}

func DiameterOfBinaryTree(root *TreeNode)int{
	var max int = 0
	DFSDiameterOfBinaryTree(root, &max)
	return int(math.Max(float64(max-1), float64(0)))
}

// func TrimBST(root *TreeNode, L int, R int) *TreeNode{
// 	queue := []*TreeNode{root}
// 	for len(queue) > 0{

// 	}
// }


func InorderIncreasingBST(node *TreeNode, result *[]*TreeNode){
	if node != nil{
		InorderIncreasingBST(node.Left, result)
		*result = append(*result, node)
		InorderIncreasingBST(node.Right, result)
	}
}


func IncreasingBST(root *TreeNode) *TreeNode{

	var result []*TreeNode
	InorderIncreasingBST(root, &result)
	result = append(result, nil)
	for i := range result{
		if result[i] != nil{
			result[i].Right = result[i+1]
			result[i].Left = nil
		}
	}
	return result[0]

}

func _SumRootToLeaf(node *TreeNode, result *int, path []int) int{
	if node != nil{
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil{
			curr := 0
			for i:=0; i < len(path); i++{ curr += (path[i] * int(math.Pow(float64(2), float64(len(path)-1 - i)))) }
			*result = (*result) + curr
		} else{
				_SumRootToLeaf(node.Left, result, path)
				_SumRootToLeaf(node.Right, result, path)
		}
	}
}

func DFSAllPossibleFBT(node *TreeNode, N int, result []*TreeNode){
	if N == 0{
		
	}
}

func AllPossibleFBT(N int) []*TreeNode{
	root := &TreeNode{
		Val: 0,
		Left: nil,
		Right: nil,
	}
	var result []*TreeNode
	DFSAllPossibleFBT(root, N-1, &result)
}

func _FindTilt(node *TreeNode, result *int) int{
	if node == nil{
		return 0
	}
	leftTilt := _FindTilt(node.Left, result)
	rightTilt := _FindTilt(node.Right, result)
	*result = (*result) + int(math.Abs(float64(leftTilt) - float64(rightTilt)))
	return node.Val + leftTilt + rightTilt
}

func FindTilt(root *TreeNode) int{
	var result int = 0
	_FindTilt(root, &result)
	return result
}

func SumRootToLeaf(root *TreeNode) int{
	var result int = 0
	_SumRootToLeaf(root, &result, []int{})
	return result
}

func _ConvertBST(node *TreeNode, prev int) int{
	if node != nil{
		nodeTemp := node.Val
		right := _ConvertBST(node.Right, prev)
		node.Val += right + prev
		left := _ConvertBST(node.Left, node.Val)
		return nodeTemp + right + left
	}
	return 0
}

func ConvertBST(root *TreeNode) *TreeNode{
	_ConvertBST(root, 0)
	return root
}

func DFSDelNodes(root *TreeNode, result *[]*TreeNode, storeMap *map[int]int){
	if root != nil{
		DFSDelNodes(root.Left, result, storeMap)
		DFSDelNodes(root.Right, result, storeMap)
        if root.Left != nil{
			if _, ok := (*storeMap)[root.Left.Val]; ok{
				root.Left = nil
			}
		}
        if root.Right != nil{
			if _, ok := (*storeMap)[root.Right.Val]; ok{
				root.Right = nil
			}
		}
        
		if _, ok := (*storeMap)[root.Val]; ok{
			if root.Left != nil{
				*result = append(*result, root.Left)
			}
			if root.Right != nil{
				*result = append(*result, root.Right)
			}
		}	
	}
} 

func DelNodes(root *TreeNode, to_delete []int) []*TreeNode{
	var result []*TreeNode
	var storeMap map[int]int = make(map[int]int)
	for i := range to_delete{
		storeMap[to_delete[i]] = 1
	}
	dummy := &TreeNode{
		Val: 0,
		Left: root,
		Right: nil,
	}
	DFSDelNodes(dummy, &result, storeMap)
	if dummy.Left != nil{
		result = append(result, root)
	}
	return result	
}

func DFSBinaryTree(node *TreeNode) []int{
	var stack []*TreeNode = []*TreeNode{node}
	var currentNode *TreeNode
	var result []int
	for len(stack) != 0{
        currentNode, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if currentNode.Right != nil{
			stack = append(stack, currentNode.Right)
		}
		if currentNode.Left != nil{
			stack = append(stack, currentNode.Left)
		}
		if currentNode.Left == nil && currentNode.Right == nil{
			result = append(result, currentNode.Val)
		}
	}
	return result
}


func LeafSimilar(root1, root2 *TreeNode) bool{
	return DFSBinaryTree(root1) == DFSBinaryTree(root2)
}


func InsertNode(root, newNode *TreeNode) bool{
	temp := root
	for temp != nil{
		if temp.Val < newNode.Val{
			if temp.Right != nil{
				temp = temp.Right
			} else{
				temp.Right = newNode
				return true
			}
			
		} else{
			if temp.Left != nil{
				temp = temp.Left
			} else{
				temp.Left = newNode
				return true
			}
		}
	}
	return false
}

func DeleteNode(root, node *TreeNode) bool{
	return false
}

func VerticalTraversal(root *TreeNode) [][]int{
	storeMap  := make(map[int]map[int][]int)
	queue := []*TreeNode{root}
	var node *TreeNode
	var coordinates []int
	queue_coordinates := [][]int{{0,0}}
	for len(queue) > 0{
		
		node, queue = queue[0], queue[1:]
		coordinates, queue_coordinates = queue_coordinates[0], queue_coordinates[1:]
		// fmt.Println("Entering with queue: ", node)
		// fmt.Println("The coordinates: ", coordinates)
		if len(storeMap[coordinates[0]]) == 0{
			// fmt.Println("init for x: ", coordinates[0])
			storeMap[coordinates[0]] = make(map[int][]int)
		}
		storeMap[coordinates[0]][coordinates[1]] = append(storeMap[coordinates[0]][coordinates[1]], node.Val)
		if node.Left != nil{
			queue_coordinates = append(queue_coordinates, []int{ coordinates[0] - 1, coordinates[1] - 1 })
			queue = append(queue, node.Left)
		}
		if node.Right != nil{
			queue_coordinates = append(queue_coordinates, []int{ coordinates[0] + 1, coordinates[1] - 1 })
			queue = append(queue, node.Right)
		}
		// fmt.Println("Exiting with queue: ", queue)
	}
	// fmt.Println(storeMap)
	var sortedArray []int
	for key := range storeMap{
  		sortedArray = append(sortedArray, key)
	}
	sort.Ints(sortedArray)
	// fmt.Println("SortedX: ", sortedArray)
	var result [][]int
	for x := range sortedArray{
		sortedY := []int{}
		for y := range storeMap[sortedArray[x]]{
			sortedY = append(sortedY, y)
		}
		sort.Ints(sortedY)
		i := len(sortedY) - 1
		temp := []int{}
		for i >= 0{
			// fmt.Printf("For x: %d, y: %d\n", sortedArray[x], sortedY[i])
			sort.Ints(storeMap[sortedArray[x]][sortedY[i]])
			for j := range storeMap[sortedArray[x]][sortedY[i]]{
				temp = append(temp, storeMap[sortedArray[x]][sortedY[i]][j])
			}
			i--
		}
		result = append(result, temp[:])
	}
	return result
}

func CommonAncestor(root, p, q *TreeNode) *TreeNode{
	var storeMap map[*TreeNode]*TreeNode = make(map[*TreeNode]*TreeNode)
	storeMap[root] = nil
	var queue []*TreeNode = []*TreeNode{root}
	var currentNode *TreeNode
	for len(queue) != 0{
		currentNode, queue = queue[0], queue[1:]
		fmt.Println("current node: ", currentNode)
		if currentNode.Left != nil{
			storeMap[currentNode.Left] = currentNode
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil{
			
			queue = append(queue, currentNode.Right)
			storeMap[currentNode.Right] = currentNode
		}
	}
	fmt.Println(storeMap)
	var ancestors map[*TreeNode]int = make(map[*TreeNode]int)
	for p != nil{
		fmt.Println(p)
		ancestors[p] = 1
		p = storeMap[p]
	}
	fmt.Println(ancestors)
	for q != nil{
		if _, ok := ancestors[q]; ok{
			return q
		}
		q = storeMap[q]
	}
	return nil
	
}

func BSTLowestCommonAncestor(root, p, q *TreeNode) *TreeNode{
	var currentNode *TreeNode = root
	for currentNode != nil{
		if (p.Val <= currentNode.Val && q.Val >= currentNode.Val) || (p.Val >= currentNode.Val && q.Val <= currentNode.Val){
			return currentNode
		} else{
			if (p.Val < currentNode.Val && q.Val < currentNode.Val){
				currentNode = currentNode.Left
			} else{
				currentNode = currentNode.Right
			}
		}
	}
	return nil
}

func _SmallestFromLeaf(root *TreeNode, s string){
	curVal := string(root.Val+97)+s
	if root.Left == nil && root.Right == nil{
        if curVal < SmallestFromLeaf{
            SmallestFromLeaf = curVal
        }
        fmt.Println("The value: ", curVal)
	} else{
		if root.Left != nil{
			_SmallestFromLeaf(root.Left, curVal)
		} 
        if root.Right != nil{
			_SmallestFromLeaf(root.Right, curVal)
		}
	}
}

func SmallestFromLeaf(root *TreeNode) string{
	_SmallestFromLeaf(root, "")
	return smallestFromLeaf
}


func main(){

	root := &TreeNode{
		Val: 3,
	}
	fmt.Println(root.Val)
	InsertNode(root, &TreeNode{Val:9,})
	InsertNode(root, &TreeNode{Val:20,})
	InsertNode(root, &TreeNode{Val:15,})
	InsertNode(root, &TreeNode{Val:7,})

	temp := root
	for temp != nil{
		fmt.Println(temp.Val)
		temp = temp.Right
	}

	// fmt.Println(VerticalTraversal(root))
	fmt.Println(CommonAncestor(root, root.Right.Right, root.Right))

}