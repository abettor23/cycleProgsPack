package main

import "fmt"

func main() {
	varOne := 0
	varTwo := 0
	varThree := 0

	for {
		if varOne != 10 {
			varOne++
			continue
		}

		if varTwo != 100 {
			varTwo++
			continue
		}

		if varThree != 1000 {
			varThree++
			continue
		}

		break
	}

	fmt.Println("Первая переменная наполнена. Значение:", varOne)
	fmt.Println("Вторая переменная наполнена. Значение:", varTwo)
	fmt.Println("Третья переменная наполнена. Значение:", varThree)
}
