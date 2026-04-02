package main

import (
	"fmt"
	"time"
)

// --- Числа Фибоначчи ---

// fibonacciRecursive — рекурсивная реализация.
// Экспоненциальная сложность O(2^n), риск переполнения стека при больших N.
func fibonacciRecursive(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

// fibonacciIterative — итеративная реализация.
// Линейная сложность O(n), без риска переполнения стека.
func fibonacciIterative(n int) int {
	if n <= 0 {
		return 0
	}
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

// --- Бинарное дерево ---

// BinaryTree — рекурсивная структура данных.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// InorderRecursive — обход дерева в глубину (inorder: лево → корень → право) через рекурсию.
func (t *BinaryTree) InorderRecursive() []int {
	if t == nil {
		return nil
	}
	var result []int
	result = append(result, t.Left.InorderRecursive()...)
	result = append(result, t.Value)
	result = append(result, t.Right.InorderRecursive()...)
	return result
}

// InorderIterative — обход дерева в глубину (inorder) через цикл и стек.
func (t *BinaryTree) InorderIterative() []int {
	var result []int
	var stack []*BinaryTree
	current := t
	for current != nil || len(stack) > 0 {
		// Идём как можно левее
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		// Берём верхний узел стека
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Value)
		current = current.Right
	}
	return result
}

func runTask1() {
	fmt.Println("=== Задание 1. Рекурсия ===")

	// Сравнение времени выполнения
	n := 35
	start := time.Now()
	resR := fibonacciRecursive(n)
	durR := time.Since(start)

	start = time.Now()
	resI := fibonacciIterative(n)
	durI := time.Since(start)

	fmt.Printf("fibonacciRecursive(%d) = %d, время: %v\n", n, resR, durR)
	fmt.Printf("fibonacciIterative(%d) = %d, время: %v\n", n, resI, durI)
	fmt.Println("Рекурсия медленнее: каждый вызов порождает два новых (O(2^n)), много повторных вычислений.")
	fmt.Println("Итерация — O(n), без накладных расходов на вызовы функций.")

	// Приблизительный максимум для рекурсии без переполнения стека:
	// Go по умолчанию имеет растущий стек (goroutine stack), поэтому технически
	// переполнение стека наступит очень поздно, но экспоненциальное время
	// делает N>45 практически недостижимым по времени ожидания.
	fmt.Println("Максимально разумное N для рекурсии (по времени): ~40-45.")

	fmt.Println()

	// Бинарное дерево:
	//        4
	//       / \
	//      2   6
	//     / \ / \
	//    1  3 5  7
	tree := &BinaryTree{
		Value: 4,
		Left: &BinaryTree{
			Value: 2,
			Left:  &BinaryTree{Value: 1},
			Right: &BinaryTree{Value: 3},
		},
		Right: &BinaryTree{
			Value: 6,
			Left:  &BinaryTree{Value: 5},
			Right: &BinaryTree{Value: 7},
		},
	}

	fmt.Println("Inorder рекурсивно:", tree.InorderRecursive())
	fmt.Println("Inorder итеративно:", tree.InorderIterative())
	fmt.Println()
	fmt.Println("Для дерева рекурсия предпочтительнее: код короче и отражает рекурсивную природу структуры.")
	fmt.Println("Итеративный обход сложнее читать, но устраняет риск переполнения стека на очень глубоких деревьях.")
}
