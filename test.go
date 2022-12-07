package main

import "fmt"

func main() {
	var workArray [10]uint8
	var a, q, w, e, r, t, y uint8
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		workArray[i] = a
	}
	fmt.Scan(&q, &w, &e, &r, &t, &y)
	workArray[q], workArray[w] = workArray[w], workArray[q]
	workArray[e], workArray[r] = workArray[r], workArray[e]
	workArray[t], workArray[y] = workArray[y], workArray[t]
	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
}
