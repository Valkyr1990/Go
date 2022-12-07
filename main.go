package main

import "fmt"

func main() {
	// put your code here
	var a, b float64
	fmt.Scan(&a)
	b = a * a
	if a <= 0 {
		fmt.Printf("число %2.2f не подходит", a)
	}
	if b > 10000 {
		fmt.Printf("%e", a)
	}
	if (b > 0 && b < 10000) && a > 0 {
		fmt.Printf("%.4f", b)
	}
}
