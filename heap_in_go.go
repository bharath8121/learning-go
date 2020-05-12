package main

import (
		"fmt"
		"container/heap"
	)

type Element struct{
	value int
	word string
	index int
}



type Heap []*Element

func (h Heap) Len() int{
	return len(h)
}

func (h Heap) Less(i, j int) bool{
	if h[i].value == h[j].value{
		return h[i].word < h[j].word
	}
	// fmt.Printf("checking: %s > %s\n", h[i].word, h[j].word)
	return h[i].value > h[j].value
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

func (h *Heap) update(element *Element, word string, value int){
	fmt.Println("Updating")
	element.word = word
	element.value = value
	heap.Fix(h, element.index)
}

func TopKFrequentsWord(words []string, k int)[]string{
	var Map map[string]int = make(map[string]int)
	for i := range words{
		Map[words[i]]++
	}

	var h Heap
	j := 0
	for i := range Map{
		fmt.Printf("%s:%d\n",i,Map[i])
		h = append(h, &Element{
			word: i,
			value: Map[i],
			index: j,
		})
		j++
	}
	heap.Init(&h)
	var result []string
	fmt.Println(len(h))
	for num:=0;num < k;num++{
		result = append(result, heap.Pop(&h).(*Element).word)
	}
	return result
}

func main(){

	fmt.Println(TopKFrequentsWord([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}
