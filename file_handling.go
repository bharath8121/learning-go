package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	)


func main(){
	data, err := ioutil.ReadFile("file.txt")

	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Printf("%s", data)
	}

	
}