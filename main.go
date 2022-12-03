package main

import "fmt"

func main() {
	var b, c, a int
	for fmt.Scan(&b); b != 0; fmt.Scan(&b) {
		if b > c {
			a = 0
			c = b
		}
		if b == c {
			a += 1
		}

	}
	fmt.Println(a)
}
