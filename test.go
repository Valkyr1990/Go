package main

import "fmt"

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

func sumInt (args ...int)(int,int) {
	var length int
	length = len(args )
	sum := 0
	for j := range args  {
		sum += j
	}
	return length, sum
}
