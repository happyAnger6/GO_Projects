package main

import (
	"fmt"
)

func main() {

	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	
	for _, x := range slice {
		fmt.Printf("[%d]\n",x)	
	}

	newSlice[1] = 100
	for _, x := range slice {
		fmt.Printf("[%d]\n",x)	
	}
	
}
