package main
import "fmt"

func main(){
var i int 

	for i =1; i <=100; i++ {
		if i % 3 == 0 && i % 5 == 0{
			fmt.Println(i, "Lipat 3 dan Lipat 5")
		}else if i % 3 == 0{
			fmt.Println(i,"Lipat 3")		
		}else if i % 5 == 0{
			fmt.Println(i,"Lipat 5")
		}else{
			fmt.Println(i)
		}
	}
}
