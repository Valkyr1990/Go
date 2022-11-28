package main

import "fmt"

func main() {
	var a int
	//c = 7
	a = 8
	//b = 91
	//c = a * b
	fmt.Println("Введите любое целое число:")
	fmt.Scan(&a)
	fmt.Println("Вы ввели число:")
	fmt.Println(a)
}
