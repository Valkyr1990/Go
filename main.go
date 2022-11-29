package main

import "fmt"

func main() {
	var b bool = 4 > 5 || 6 > 8   // false
	var c bool = 3 == 5 || 10 > 8 // true
	fmt.Println(b, c)             // false
}
