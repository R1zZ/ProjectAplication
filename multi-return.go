/*
(int, int) pada fungsi ditujukan untuk fungsi returns 2 int.
Disini menggunakan 2 nilai berbeda utk return.
*/

package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}
func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	/*
	   Gunakan indetifier kosong (_), jika anda ingin mem-bypass return value
	   tertentu.
	*/
	_, c := vals()
	fmt.Println(c)

	d, _ := vals()
	fmt.Println(d)
}
