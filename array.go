package main

import "fmt"

func main(){
	var i int
	x := [5] int { 2,4,6,8,10}
	
	var total int = 0 
	for i=0; i<len(x); i++{
		total += x[i]
	}
	
	fmt.Println(total/len(x))
}