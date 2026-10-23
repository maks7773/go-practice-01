package main

import "fmt"

func main() {
	var name string = "Максим"
	var group string = "Учебная-01"
	const course = "Бэкенд-разработка на языке Go"
	completed := 1
	var interes bool = true
	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Println("Курс: ", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бэкенду: %t\n", interes)
}
