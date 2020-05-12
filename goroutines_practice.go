package main

import (
	"fmt"
	"time"
	"strconv"
	"sync"
	)


func SetsWaitVariables(channel chan string, waitVar sync.WaitGroup){
	defer waitVar.Done()
	SetsVariables(channel)
}

func GetsWaitVariables(channel chan string, waitVar sync.WaitGroup){
	defer waitVar.Done()
	GetsVariables(channel)	
	// Insted of using defer, waitVar.Done() can also be called at the end of function
}

func GetsVariables(channel chan string){
	for str := range channel{
		fmt.Println(str)	
	}
}

func SetsVariables(channel chan string){
	for i:=0; i<10; i++{
		channel <- ("string_"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	close(channel)
}

func WaitBasedGoRoutines(channel chan string){
	var waitVar sync.WaitGroup
	waitVar.Add(2)
	go SetsWaitVariables(channel, waitVar)
	go GetsWaitVariables(channel, waitVar)
	waitVar.Wait()
}

func NoobGoRoutines(channel chan string){
	go SetsVariables(channel)
	go GetsVariables(channel)
	var input string
	fmt.Scanln(&input)
}

func main(){
	channel := make(chan string)
	WaitBasedGoRoutines(channel)
	// NoobGoRoutines(channel)
}
