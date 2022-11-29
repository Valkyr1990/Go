package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a > 360 {
		a = 360
	}
	if a < 0 {
		a = 0
	}
	var c int
	c = a / 30
	var b int
	b = 2 * (a % 30)
	fmt.Println("It is", c, "hours", b, "minutes.")

}
