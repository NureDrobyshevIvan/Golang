package main

import (
	"fmt"
	"time"
)

// Звичайна функція
func sayHello() {
	fmt.Println("Привіт!")
}

func main() {
	// Запуск функції у новій горутині
	go sayHello()

	// Основна функція (сама є горутиною) продовжує виконання
	fmt.Println("Я в main!")

	// Невелика пауза, щоб горутина встигла виконатися
	time.Sleep(50 * time.Millisecond)
}
