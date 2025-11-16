package main

import (
	"fmt"
	"time"
)

// heavyTask - симуляція важкої роботи для CPU
func heavyTask(id int) int {
	// Симуляція "важкої" роботи
	time.Sleep(50 * time.Millisecond)
	result := id * 2
	fmt.Printf("Завдання %d виконано, результат %d\n", id, result)
	return result
}

func main() {
	const numJobs = 20

	startTime := time.Now()
	fmt.Println("--- Початок послідовної обробки ---")

	var results []int
	for i := 1; i <= numJobs; i++ {
		result := heavyTask(i)
		results = append(results, result)
	}

	fmt.Println("--- Послідовну обробку завершено ---")
	fmt.Printf("Загальний час: %s\n", time.Since(startTime))
}
