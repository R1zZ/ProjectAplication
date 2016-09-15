package main
import "fmt"
type orang struct {
  nama string
  umur int
}
func main() {
  fmt.Println(orang{"cinta",20})
  fmt.Println(orang{nama: "anti",umur : 30})
  fmt.Println(orang{nama: "rangga"})
  fmt.Println(&orang {nama: "beti", umur: 40})
  s := orang{nama:"budi",umur: 50}
  fmt.Println(s.nama)
  sp := &s
  fmt.Println(sp.umur)
  sp.umur= 51
  fmt.Println(sp.umur)
}
