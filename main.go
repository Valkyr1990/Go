package main

import "fmt"

func main() {
	ExampleScope1()

}

func ExampleScope1() {
	var v int = 1

	fmt.Println(v)

	/*
	 * В примере мы объявили в отдельных лексических блоках переменные с одним именем и
	 * разными типами, а затем напечатали значения этих переменных.
	 */

	// Output:
	// 2
	// 1
}
