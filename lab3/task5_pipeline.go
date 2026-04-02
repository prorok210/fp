package main

import "fmt"

// --- Вспомогательные функции высшего порядка ---

// Map (переиспользуется из лаб 2) — применяет f к каждому элементу среза.
func Map(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = f(v)
	}
	return result
}

// Filter — возвращает новый срез из элементов, для которых pred вернул true.
func Filter(nums []int, pred func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce — сворачивает срез в одно значение, применяя f последовательно.
// init — начальное значение аккумулятора.
func Reduce(nums []int, f func(int, int) int, init int) int {
	acc := init
	for _, v := range nums {
		acc = f(acc, v)
	}
	return acc
}

// --- Вспомогательные функции-предикаты и преобразования ---

// isEven — проверка на чётность.
func isEven(n int) bool { return n%2 == 0 }

// multiply — возвращает функцию, умножающую аргумент на factor.
func multiply(factor int) func(int) int {
	return func(n int) int { return n * factor }
}

// add — сложение двух чисел.
func add(a, b int) int { return a + b }

// mul — произведение двух чисел (для Reduce).
func mul(a, b int) int { return a * b }

// max — максимум из двух чисел (для Reduce).
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// --- SlicePipeline ---

type SlicePipeline struct {
	data []int
}

func NewSlicePipeline(data []int) *SlicePipeline {
	// Копируем срез, чтобы не изменять исходный
	cp := make([]int, len(data))
	copy(cp, data)
	return &SlicePipeline{data: cp}
}

func (p *SlicePipeline) Filter(pred func(int) bool) *SlicePipeline {
	p.data = Filter(p.data, pred)
	return p
}

func (p *SlicePipeline) Map(transform func(int) int) *SlicePipeline {
	p.data = Map(p.data, transform)
	return p
}

func (p *SlicePipeline) Reduce(f func(int, int) int, init int) int {
	return Reduce(p.data, f, init)
}

func (p *SlicePipeline) Result() []int {
	return p.data
}

func runTask5() {
	fmt.Println("=== Задание 5. Конвейеры (SlicePipeline) ===")

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Исходный срез:", nums)

	// Отфильтровать чётные → умножить на 3 → сумма
	sum := NewSlicePipeline(nums).
		Filter(isEven).
		Map(multiply(3)).
		Reduce(add, 0)
	fmt.Println("filter(чётные) → map(*3) → reduce(+):  ", sum)

	// Промежуточный результат: только фильтрация и map
	filtered := NewSlicePipeline(nums).Filter(isEven).Result()
	fmt.Println("filter(чётные):                          ", filtered)

	mapped := NewSlicePipeline(nums).Filter(isEven).Map(multiply(3)).Result()
	fmt.Println("filter(чётные) → map(*3):                ", mapped)

	fmt.Println()
	fmt.Println("--- Демонстрация Reduce ---")

	// Сумма
	sumAll := NewSlicePipeline(nums).Reduce(add, 0)
	fmt.Println("Сумма 1..10:          ", sumAll)

	// Произведение
	product := NewSlicePipeline(nums).Reduce(mul, 1)
	fmt.Println("Произведение 1..10:   ", product)

	// Максимум
	maximum := NewSlicePipeline(nums).Reduce(max, nums[0])
	fmt.Println("Максимум 1..10:       ", maximum)
}
