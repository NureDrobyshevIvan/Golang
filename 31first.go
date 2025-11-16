package main

import (
	"fmt"
	"net/http"
	"time"
)

// checkSite виконує HTTP GET запит до сайту
func checkSite(url string) {
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
		"https://lab141dsd.com", // Неіснуючий сайт
	}

	// Заміряємо загальний час
	startTime := time.Now()
	fmt.Println("--- Початок послідовної перевірки ---")

	// Запускаємо перевірку послідовно, один за одним
	for _, site := range sites {
		checkSite(site)
	}

	// Виводимо загальний час
	fmt.Println("--- Послідовну перевірку завершено ---")
	fmt.Printf("Загальний час: %s\n", time.Since(startTime))
}
