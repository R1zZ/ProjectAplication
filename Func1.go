/*
Fungsi
Deklarasi fungsi memili format seperti berikut :
func namafungsi (variabel type_data) { scope }

Memungkinkan untuk menggunakan fungsi lebih dari satu, seperti contoh dibawah
terdapat 2 fungsi, yaitu perhitungan dan main. perhitungan mencakup proses
penyelesaian persoalan aritmatika, sementara main digunakan untuk memberikan
input dan keluaran ke user.

Selain itu karena menggunakan perhitungan aritmatika, dalam contoh ini juga
menyertakan package math untuk memanggil built in fungsi aritmatika.
*/

package main

import (
	"fmt"
	"math"
)

//a float64, b float64, c float64 adalah pendefisian variable berikut tipe data.

func perhitungan(a float64, b float64, c float64) {
	D := b*b - 4*a*c

	switch {
	case D > 0:
		fmt.Println("X1=", (-b+math.Sqrt(D))/(2*a))
		fmt.Println("X2=", (-b-math.Sqrt(D))/(2*a))
	case D == 0:
		fmt.Print("X=", -b/(2*a))
	default:
		fmt.Println("Tidak ada akar")
	}
}

func main() {

	/*
	   karena variable dalam fungsi perhitungan berada dalam scope fungsi perhitungan,
	   sehingga pada fungsi main kita perlu mendefenisikan ulang variable a, b, c.
	*/

	var a, b, c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)

	perhitungan(a, b, c)

}
