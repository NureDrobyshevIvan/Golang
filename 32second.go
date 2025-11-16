package main

import (
	"fmt"
	"sync"
	"time"
)

// heavyTask - та сама симуляція важкої роботи
func heavyTask(id int) int {
	time.Sleep(50 * time.Millisecond)
	result := id * 2
	return result
}

// worker - це горутина-"робітник"
// Вона читає завдання з каналу 'jobs' і пише результати в канал 'results'
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 'for range' по каналу 'jobs'
	// Цей цикл автоматично завершиться, коли канал 'jobs' буде закрито
	for j := range jobs {
		fmt.Printf("Робітник %d почав завдання %d\n", id, j)
		result := heavyTask(j)
		fmt.Printf("Робітник %d завершив завдання %d, результат %d\n", id, j, result)
		results <- result // Відправляємо результат у канал
	}
}

func main() {
	const numJobs = 20
	const numWorkers = 4 // Кількість робітників (наприклад, за кількістю ядер)

	// Створюємо буферизовані канали
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	startTime := time.Now()
	fmt.Println("--- Початок конкурентної обробки ---")

	// 1. Запускаємо 'numWorkers' горутин-робітників
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 2. Відправляємо всі завдання в канал 'jobs'
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// Закриваємо канал 'jobs', сигналізуючи робітникам, що завдань більше не буде
	close(jobs)

	// 3. Чекаємо, доки всі робітники завершать роботу
	wg.Wait()

	// 4. Закриваємо канал 'results' (оскільки в нього вже ніхто не пише)
	close(results)

	// 5. Збираємо результати (необов'язково для заміру часу, але для повноти)
	var allResults []int
	for r := range results {
		allResults = append(allResults, r)
	}

	fmt.Println("--- Конкурентну обробку завершено ---")
	fmt.Printf("Загальний час: %s\n", time.Since(startTime))
	// fmt.Println("Всі результати зібрано:", len(allResults))
}
