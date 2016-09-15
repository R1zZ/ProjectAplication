package main

import "fmt"

func main(){
	slice := []int {1,2,3}
	
	slice1 := append (slice, 4,5)
	fmt.Println(slice,slice1)
	
	slice2 := [] int {1,2,3}
	slice4 := make([]int,3)
	
	fmt.Println(slice2,slice4)
	
	copy(slice4,slice2)
	fmt.Println(slice2,slice4)
	
	
}
