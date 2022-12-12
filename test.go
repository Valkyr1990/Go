package main

import "fmt"

func main() {
	var x, q, c, a, w int
	fmt.Scan(&x)
	for i := 10; i <= x*10; i *= 10 {
		t1 := (x % i) / (i / 10)
		q = q + t1
		for q > 9 {
			c = q / 10
			a = q % 10
		}
		w = a + c
	}
	fmt.Print(w)
}
