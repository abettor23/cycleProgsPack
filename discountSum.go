package main

import "fmt"

func main() {
	fmt.Println("Введите стоимость товара:")
	var productPrice int
	fmt.Scan(&productPrice)

	for {
		fmt.Println("Введите размер скидки:")
		var discount int
		fmt.Scan(&discount)

		discountLimit := (productPrice / 100) * 30

		if discount <= discountLimit && discount <= 2000 {
			fmt.Println("Скидка успешно применена. Размер скидки:", discount)
			break
		} else {
			fmt.Println("Ошибка!Размер скидки не соответствует условиям.")
			fmt.Println("Скидка должна быть не больше 30% от цены товара и не больше 2000 рублей.")
			fmt.Println("Попробуйте снова.")
		}
	}
}
