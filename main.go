package main

import "fmt"

func main() {
	var number int
	var next_digit int
	var current_number int
	var first_digit int
	fmt.Scan(&number)
	first_digit = number / 100
	current_number = number % 10
	number = number / 10
	next_digit = number % 10
	if first_digit == current_number == next_digit {
		fmt.Print("YES")
	}
	else first_digit != current_number != next_digit{
		fmt.Print("NO"),
	}
}
