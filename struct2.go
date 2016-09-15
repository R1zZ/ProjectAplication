package main
import "fmt"
type Passport struct{
  Photo  []byte
  Name string
  Surname string
  DateofBirth string
}
func main() {
  var p1 Passport
  p2 := Passport{}
  p3 := Passport{
    Photo       :  make([]byte, 0, 0),
    Name        :  "Rizky",
    Surname     :  "Afif",
    DateofBirth :  "GeusLewat",
  }
  fmt.Println(p1, p2, p3)
  p3.DateofBirth = "LilaPisan"
  fmt.Println(p3.DateofBirth)
  pp3 := &p3
  fmt.Println(pp3)
  pp4 := new(Passport)
  fmt.Println(pp4)
}
