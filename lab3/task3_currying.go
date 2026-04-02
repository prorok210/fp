package main

import "fmt"

// --- Каррирование ---

// volume — обычная функция: объём прямоугольного параллелепипеда.
func volume(l, w, h float64) float64 {
	return l * w * h
}

// volumeCurried — каррированная версия volume.
// Принимает аргументы по одному и возвращает следующую функцию.
func volumeCurried(l float64) func(float64) func(float64) float64 {
	return func(w float64) func(float64) float64 {
		return func(h float64) float64 {
			return l * w * h
		}
	}
}

// --- Частичное применение ---

// greet — собирает строку приветствия из трёх аргументов.
func greet(greeting, name, punctuation string) string {
	return greeting + ", " + name + punctuation
}

// partialGreet — фиксирует приветствие и знак препинания,
// возвращает функцию, ожидающую имя.
func partialGreet(greeting, punctuation string) func(string) string {
	return func(name string) string {
		return greet(greeting, name, punctuation)
	}
}

func runTask3() {
	fmt.Println("=== Задание 3. Каррирование и частичное применение ===")

	// Обычная функция
	fmt.Printf("volume(3, 4, 5)          = %.1f\n", volume(3, 4, 5))

	// Каррированная версия
	vol := volumeCurried(3)(4)(5)
	fmt.Printf("volumeCurried(3)(4)(5)   = %.1f\n", vol)

	// Частичное применение: фиксируем длину
	volumeWith3 := volumeCurried(3)
	fmt.Printf("volumeWith3(4)(5)        = %.1f\n", volumeWith3(4)(5))
	fmt.Printf("volumeWith3(2)(6)        = %.1f\n", volumeWith3(2)(6))

	fmt.Println()

	// greet
	g1 := greet("Привет", "Иван", "!")
	fmt.Println(g1)
	g2 := greet("Здравствуйте", "Иван Иванович", ".")
	fmt.Println(g2)

	fmt.Println()

	// partialGreet
	informalGreet := partialGreet("Привет", "!")
	fmt.Println(informalGreet("Иван"))
	fmt.Println(informalGreet("Мария"))

	formalGreet := partialGreet("Здравствуйте", ".")
	fmt.Println(formalGreet("Иван Иванович"))
}
