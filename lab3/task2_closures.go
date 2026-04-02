package main

import "fmt"

// GetFibonacciGenerator возвращает замыкание, которое при каждом вызове
// возвращает следующее число Фибоначчи (0, 1, 1, 2, 3, 5, …).
// Замыкание захватывает переменные a и b из внешней функции и сохраняет
// состояние последовательности между вызовами.
func GetFibonacciGenerator() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func runTask2() {
	fmt.Println("=== Задание 2. Замыкания ===")

	nextFibonacci := GetFibonacciGenerator()
	for i := 0; i < 10; i++ {
		fmt.Printf("nextFibonacci() = %d\n", nextFibonacci())
	}

	fmt.Println()

	// Два независимых генератора — каждый хранит своё состояние
	gen1 := GetFibonacciGenerator()
	gen2 := GetFibonacciGenerator()
	fmt.Printf("gen1: %d %d %d\n", gen1(), gen1(), gen1())
	fmt.Printf("gen2: %d %d %d\n", gen2(), gen2(), gen2())
}
