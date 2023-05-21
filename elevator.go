package main

import "fmt"

func main() {
	highestFloor := 24
	elevatorPosition := 1
	downgrade := false

	needStopTen := true
	needStopSeven := true
	needStopFour := true

	passengersAmount := 3
	elevatorFullness := 0
	elevatorLimit := 2

	fmt.Println("Пассажиров ожидающих лифт:", passengersAmount)

	for passengersAmount > 0 || elevatorFullness > 0 {
		if !downgrade {
			if elevatorPosition < highestFloor {
				elevatorPosition++
				fmt.Println("Лифт на", elevatorPosition, "этаже.")
			} else {
				downgrade = true
				fmt.Println("Лифт закончил подъем.")
			}
			continue
		} else {
			if elevatorPosition != 10 || elevatorPosition != 7 || elevatorPosition != 4 {
				elevatorPosition--
				fmt.Println("Лифт на", elevatorPosition, "этаже.")
			}
			if elevatorPosition == 10 && elevatorFullness < elevatorLimit && needStopTen {
				fmt.Println("Пассажир зашел в лифт.")
				passengersAmount--
				fmt.Println("Осталось пассажиров:", passengersAmount)
				elevatorFullness++
				fmt.Println("Осталось места в лифте:", elevatorLimit-elevatorFullness)
				needStopTen = false
				elevatorPosition--
			}
			if elevatorPosition == 7 && elevatorFullness < elevatorLimit && needStopSeven {
				fmt.Println("Пассажир зашел в лифт.")
				passengersAmount--
				fmt.Println("Осталось пассажиров:", passengersAmount)
				elevatorFullness++
				fmt.Println("Осталось места в лифте:", elevatorLimit-elevatorFullness)
				needStopSeven = false
				elevatorPosition--
			}
			if elevatorPosition == 4 && elevatorFullness < elevatorLimit && needStopFour {
				fmt.Println("Пассажир зашел в лифт.")
				passengersAmount--
				fmt.Println("Осталось пассажиров:", passengersAmount)
				elevatorFullness++
				fmt.Println("Осталось места в лифте:", elevatorLimit-elevatorFullness)
				needStopFour = false
				elevatorPosition--
			}
			if elevatorPosition == 1 {
				fmt.Println("Осталось пассажиров:", passengersAmount)
				elevatorFullness = 0
				fmt.Println("Пассажиры вышли. Осталось места:", elevatorLimit-elevatorFullness)
				downgrade = false
				fmt.Println("Лифт закончил спуск.")
			}
		}
	}
}
