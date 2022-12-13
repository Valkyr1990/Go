package main

import "fmt"

func main() {
	var x, a, result int
	q := 1
	fmt.Scan(&x)
	fmt.Scan(&a)
	for i := 10; i <= x*10; i *= 10 {
		t1 := (x % i) / (i / 10)
		if t1 != a {
			result = result + t1*q
			q = q * 10
		}
	}
	fmt.Print(result)
}
