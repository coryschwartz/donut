package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, i, j float64
	var k int
	var s = ".,-~:;=!*#$@"
	fmt.Printf("\x1b[2J")
	for {
		z := make([]float64, 1760)
		b := make([]byte, 1760)
		for bi := 0; bi < len(b); bi++ {
			b[bi] = byte(32)
		}
		for j = 0; j < 6.28; j += 0.07 {
			for i = 0; i < 6.28; i += 0.02 {
				c := math.Sin(i)
				d := math.Cos(j)
				e := math.Sin(A)
				f := math.Sin(j)
				g := math.Cos(A)
				h := d + 2
				D := 1 / (c*h*e + f*g + 5)
				l := math.Cos(i)
				m := math.Cos(B)
				n := math.Sin(B)
				t := c*h*g - f*e
				x := int(40 + 30*D*(l*h*m-t*n))
				y := int(12 + 15*D*(l*h*n+t*m))
				o := int(x + 80*y)
				N := int(8 * ((f*e-c*d*g)*m - c*d*e - f*g - l*d*n))
				if 22 > y && y > 0 && x > 0 && 80 > x && D > z[o] {
					z[o] = D
					if N > 0 {
						b[o] = s[N]
					} else {
						b[o] = s[0]
					}
				}
			}
		}
		fmt.Printf("\x1b[H")
		for k = 0; k < 1761; k++ {
			if k%80 != 0 {
				fmt.Printf(string(b[k]))
			} else {
				fmt.Printf("\n")
			}
			A += 0.00004
			B += 0.00002
		}
	}
}
