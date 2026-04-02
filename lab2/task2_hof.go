package main

import (
	"errors"
	"fmt"
	"math"
)

// apply — обычная функция, выполняет арифметическое действие по строковому оператору.
// НЕ является функцией высшего порядка.
func apply(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("действие не поддерживается: %q", operator)
	}
}

// applyFunc — функция высшего порядка: принимает функцию f как аргумент.
// Является функцией высшего порядка, потому что принимает другую функцию (f) в качестве аргумента.
func applyFunc(a, b float64, f func(float64, float64) float64) float64 {
	return f(a, b)
}

func runTask2() {
	fmt.Println("=== Задание 2. Функции высшего порядка ===")

	// Проверка apply
	res, err := apply(3, 5, "+")
	fmt.Printf("apply(3, 5, \"+\") = %.1f, err=%v\n", res, err)

	res, err = apply(7, 10, "*")
	fmt.Printf("apply(7, 10, \"*\") = %.1f, err=%v\n", res, err)

	res, err = apply(3, 5, "#")
	fmt.Printf("apply(3, 5, \"#\") = %.1f, err=%v\n", res, err)

	fmt.Println()

	// applyFunc с анонимными функциями
	add := func(a, b float64) float64 { return a + b }
	sub := func(a, b float64) float64 { return a - b }
	mul := func(a, b float64) float64 { return a * b }
	div := func(a, b float64) float64 { return a / b }

	fmt.Printf("applyFunc(3, 5, add) = %.1f\n", applyFunc(3, 5, add))
	fmt.Printf("applyFunc(10, 4, sub) = %.1f\n", applyFunc(10, 4, sub))
	fmt.Printf("applyFunc(3, 5, mul) = %.1f\n", applyFunc(3, 5, mul))
	fmt.Printf("applyFunc(10, 4, div) = %.1f\n", applyFunc(10, 4, div))

	fmt.Println()

	// applyFunc поддерживает операции, не предусмотренные в apply:
	// возведение в степень
	pow := func(a, b float64) float64 { return math.Pow(a, b) }
	fmt.Printf("applyFunc(2, 10, pow) = %.1f\n", applyFunc(2, 10, pow))

	// «конкатенация» строк через числовое представление невозможна напрямую,
	// но можно показать любую произвольную операцию, например гипотенуза:
	hypotenuse := func(a, b float64) float64 { return math.Sqrt(a*a + b*b) }
	fmt.Printf("applyFunc(3, 4, hypotenuse) = %.1f\n", applyFunc(3, 4, hypotenuse))

	fmt.Println()
	fmt.Println("Вывод:")
	fmt.Println("  apply    — расширяется только изменением кода (жёсткий switch).")
	fmt.Println("  applyFunc — расширяется без изменения самой функции: передаём новую операцию как аргумент.")
}
