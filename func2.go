/*
untuk memahami return pada program ini, silahkan merujuk pada file
return.go dan multi-return.go
*/

package main

import (
	"fmt"
	"math"
)

func perhitungan(a float64, b float64, c float64) (int, float64, float64) {
	D := b*b - 4*a*c

	switch {
	case D > 0:
		return 2, (-b + math.Sqrt(D)) / (2 * a), (-b - math.Sqrt(D)) / (2 * a)
	case D == 0:
		return 1, -b / (2 * a), 0
	default:
		return 0, 0, 0
	}
}

func main() {
	var a, b, c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)

	fmt.Println(perhitungan(a, b, c))

}
