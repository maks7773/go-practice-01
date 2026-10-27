package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестен"

	}
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {

	if len(statuses) == 0 {
		return 0
	}
	count := 0
	for _, status := range statuses {
		if status == target {
			count++
		}
	}
	return count
}

func findStatus(statuses []string, target string) int {
	index := -1
	for i, status := range statuses {
		if status == target {
			index = i
			break
		}
	}
	return index
}

func completionPercent(statuses []string) float64 {
	if len(statuses) == 0 {
		return 0
	}
	var done int
	for _, status := range statuses {
		if status == "done" {
			done++
		}
	}
	return float64(100*done) / float64(len(statuses))
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}
	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d \n", counts["new"], counts["in_progress"], counts["done"])
	i := findStatus(statuses, "done")
	fmt.Println(i)
	done := completionPercent(statuses)
	fmt.Printf("Done: %.0f percents", done)
}
