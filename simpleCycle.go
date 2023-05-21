package main

import "fmt"

func main() {
	fmt.Println("Напишите произвольное число:")
	var number int
	fmt.Scan(&number)

	for i := 0; i <= number; i++ {
		fmt.Println(i)
	}
}
