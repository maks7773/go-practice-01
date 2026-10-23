package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	var n3 int

	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)
	var summa int = n1 + n2 + n3
	var average float64 = float64(n1+n2+n3) / 3
	fmt.Printf("Сумма: %d\n", summa)
	fmt.Printf("среднее: %.1f\n", average)
}
