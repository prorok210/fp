package main

import "fmt"

// TransformImperative — императивный стиль: умножает каждый элемент на 2 через цикл.
// Исходный срез не изменяется.
func TransformImperative(nums []int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = v * 2
	}
	return result
}

// Map — функция высшего порядка: применяет f к каждому элементу среза.
// Исходный срез не изменяется.
func Map(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = f(v)
	}
	return result
}

func runTask3() {
	fmt.Println("=== Задание 3. Императивный vs функциональный Map ===")

	original := []int{1, 2, 3, 4, 5}

	// Императивный подход
	resultImperative := TransformImperative(original)
	fmt.Println("Исходный срез:              ", original)
	fmt.Println("TransformImperative (*2):   ", resultImperative)

	// Функциональный подход
	resultMap := Map(original, func(n int) int { return n * 2 })
	fmt.Println("Map(n*2):                   ", resultMap)

	fmt.Println()

	// Пункт 6: смена задачи — возвести в квадрат.
	// Императивно: нужно изменить тело TransformImperative (result[i] = v*v).
	// Функционально: достаточно передать другую анонимную функцию — код Map не трогаем.
	resultSquare := Map(original, func(n int) int { return n * n })
	fmt.Println("Map(n^2):                   ", resultSquare)

	fmt.Println()
	fmt.Println("Вывод:")
	fmt.Println("  TransformImperative — для смены операции нужно менять тело функции.")
	fmt.Println("  Map — операция передаётся снаружи; сама функция остаётся неизменной.")
}
