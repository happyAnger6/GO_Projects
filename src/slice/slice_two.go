package main

import (
	_ "fmt"
)

func main() {

	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	
	for x range slice {
		fmt.Printf("[%d]\n",x)	
	}
	

}
