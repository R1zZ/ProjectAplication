package main
import "fmt"

func main(){
	var i int
	
	for i = 1; i <=100; i++{
		if i % 3 == 0{
			fmt.Print(i,", ")
		}
	}
	fmt.Println("")
}
