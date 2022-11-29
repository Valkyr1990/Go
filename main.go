package main

import "fmt"

func main() {
	//var b bool = 4 > 5 || 6 > 8   // false
	//var c bool = 3 == 5 || 10 > 8 // true
	//fmt.Println(c, b)           // false

	// Операция И
	// Операция ИЛИ
	// Операция НЕ

	//// ИЛИ
	//1 + 1 = 1
	//0 + 1 = 1
	//1 + 0 = 1
	//0 + 0 = 0
	//
	//// И
	//1 + 1 = 1
	//1 + 0 = 0
	//0 + 1 = 0
	//0 + 0 = 0
	//
	//// ИЛИ
	//true || true = true
	//false || true = true
	//true || false = true
	//false || false = false
	//
	//// И
	//true && true = true
	//true && false = false
	//false && true = false
	//false && false = false

	//var a int
	//
	//result, err := fmt.Scan(&a)
	//if err != nil {
	//	fmt.Println("ПАНИКА!!!!", err)
	//}

	secretDataCardNumber := "4272 2007 5435 8842"
	expiredDate := 04 / 27
	cvc := 769

	var login string
	var password string
	login = "valkyr1990"
	password = "qwerty"
	var b string
	var c string
	fmt.Println("Введите логин:")
	fmt.Scan(&b)
	fmt.Println("Введите пароль:")
	fmt.Scan(&c)

	if b == login && c == password {
		fmt.Println("Держи мою карту!")
		fmt.Println("НОМЕР:", secretDataCardNumber)
		fmt.Println("Дата:", expiredDate)
		fmt.Println("ЦВЦ:", cvc)
	} else {
		fmt.Println("Я НЕ ДАМ ТЕБЕ СВОЮ КАРТУ ЛЖИВОЕ ДЕРЬМО!!!!")
	}

	//fmt.Println(a, result)

}
