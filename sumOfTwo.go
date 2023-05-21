package main

import "fmt"

func main() {
	fmt.Println("Введите первое число:")
	var numOne int
	fmt.Scan(&numOne)

	fmt.Println("Введите второе число:")
	var numTwo int
	fmt.Scan(&numTwo)

	sum := numOne + numTwo

	for numOne < sum {
		numOne++
	}

	fmt.Println("Сумма чисел:", numOne)
}
