package main

import (
	"fmt"
	"time"
)

// fibonacciRecursive - обычная рекурсивная реализация (медленная).
func fibonacciRecursive(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

// fibonacciMemoized возвращает функцию, вычисляющую числа Фибоначчи
// с использованием мемоизации через замыкание.
func fibonacciMemoized() func(int) int {
	cache := make(map[int]int)
	var f func(int) int
	f = func(n int) int {
		if val, ok := cache[n]; ok {
			return val
		}
		if n <= 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		res := f(n-1) + f(n-2)
		cache[n] = res
		return res
	}
	return f
}

func runTask1() {
	fmt.Println("=== Задание 1. Мемоизация ===")

	n := 40

	// Проверка без мемоизации
	fmt.Println("Вычисление без мемоизации...")
	start := time.Now()
	res1 := fibonacciRecursive(n)
	fmt.Printf("fibonacciRecursive(%d) = %d (Время: %v)\n", n, res1, time.Since(start))

	// Проверка с мемоизацией
	memoized := fibonacciMemoized()

	fmt.Println("Первый вызов с мемоизацией (кеш заполняется)...")
	start = time.Now()
	res2 := memoized(n)
	fmt.Printf("memoized(%d) = %d (Время: %v)\n", n, res2, time.Since(start))

	fmt.Println("Второй вызов с мемоизацией (чтение из кеша)...")
	start = time.Now()
	res3 := memoized(n)
	fmt.Printf("memoized(%d) = %d (Время: %v)\n", n, res3, time.Since(start))
}
