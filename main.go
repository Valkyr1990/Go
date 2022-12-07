package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите размерность:")
	value, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(value[:len(value)-1])
	if err != nil {
		panic(err)
	}
	mas := make([][]float64, n)
	for i := 0; i < n; i++ {
		fmt.Println("Массив ", i+1, ":")
		for j := 0; j < n; j++ {
			fmt.Println("Введите ", j+1, " число:")
			value, _ = reader.ReadString('\n')
			fnum, err := strconv.ParseFloat(value[:len(value)-1], 64)
			if err != nil {
				panic(err)
			}
			mas[i] = append(mas[i], fnum)
		}
	}

	fmt.Println(mas)
}
