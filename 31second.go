package main

import (
	"fmt"
	"net/http"
	"sync" // Імпортуємо пакет sync
	"time"
)

// checkSite тепер приймає WaitGroup для синхронізації
func checkSite(url string, wg *sync.WaitGroup) {
	// Зменшуємо лічильник WaitGroup, коли функція завершується
	defer wg.Done()

	start := time.Now()
	_, err := http.Get(url)
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("ПОМИЛКА: %s недоступний (час: %s)\n", url, duration)
		return
	}
	fmt.Printf("УСПІХ: %s доступний (час: %s)\n", url, duration)
}

func main() {
	sites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://dl.nure.ua",
		"https://github.com",
		"https://non-existent-site-123.com",
	}

	// Створюємо екземпляр WaitGroup
	var wg sync.WaitGroup

	startTime := time.Now()
	fmt.Println("--- Початок конкурентної перевірки ---")

	// Запускаємо перевірку конкурентно
	for _, site := range sites {
		// 1. Повідомляємо WaitGroup, що ми додаємо одне завдання
		wg.Add(1)

		// 2. Запускаємо checkSite у новій горутині
		go checkSite(site, &wg)
	}

	// 3. Чекаємо, доки лічильник WaitGroup не стане 0
	wg.Wait()

	// Виводимо загальний час
	fmt.Println("--- Конкурентну перевірку завершено ---")
	fmt.Printf("Загальний час: %s\n", time.Since(startTime))
}
