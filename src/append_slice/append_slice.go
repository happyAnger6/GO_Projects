package main

import (
	"fmt"
)

func main(){
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]

	for i, x := range slice {
		fmt.Printf("i:[%d] --- [%d]\n", i, x)
	}

	newSlice = append(newSlice, 60)

	for i, x := range slice {
		fmt.Printf("i:[%d] --- [%d]\n", i, x)
	}
}
