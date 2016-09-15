package main
import "time"
import "fmt"
func main() {
  channel1 := make(chan string) //membuat channel1
  channel2 := make(chan string) //membuat channel2
  go func() {
    time.Sleep(time.Second * 4)
    channel1 <- "satu"
  }() /*menjalankan function go yang dimana time.second akan dikalikan 4 lalu
  channel1 yang akan mengeluarkan comment satu */
  go func() {
    time.Sleep(time.Second * 2)
    channel2 <- "dua"
  }()/*sama seperti channel1 ini akan mengalikan time.second namun disini di kali
  2*/
  for i := 0; i < 2; i++{
    select{
    case msg1 := <-channel1:
      fmt.Println("menerima",msg1)
    case msg2 := <-channel2:
      fmt.Println("menerima",msg2)
    }
  }
}
