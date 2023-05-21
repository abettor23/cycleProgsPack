package main

import "fmt"

func main() {
	fmt.Println("Заполним яблоками три корзины.")

	fmt.Println("Укажите вместимость первой корзины:") //узнаем у пользователя объем (cap, capacity) трёх корзин
	var basketOneCap int
	fmt.Scan(&basketOneCap)

	fmt.Println("Укажите вместимость второй корзины:")
	var basketTwoCap int
	fmt.Scan(&basketTwoCap)

	fmt.Println("Укажите вместимость третьей корзины:")
	var basketThreeCap int
	fmt.Scan(&basketThreeCap)

	basketOneFull := 0 //фактическая наполненность (full, fullness) корзин на данный момент
	basketTwoFull := 0
	basketThreeFull := 0

	for {
		if basketOneFull != basketOneCap {
			basketOneFull++
			continue
		}

		if basketTwoFull != basketTwoCap {
			basketTwoFull++
			continue
		}

		if basketThreeFull != basketThreeCap {
			basketThreeFull++
			continue
		}

		break
	}

	fmt.Println("Первая корзина наполнена. Объём:", basketOneFull)
	fmt.Println("Вторая корзина наполнена. Объём:", basketTwoFull)
	fmt.Println("Третья корзина наполнена. Объём:", basketThreeFull)
}
