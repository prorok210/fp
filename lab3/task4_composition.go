package main

import (
	"fmt"
	"math"
)

// sin и square — базовые функции для демонстрации композиции.
func sin(x float64) float64    { return math.Sin(x) }
func square(x float64) float64 { return x * x }

// compose — композиция функций: compose(f, g)(x) = f(g(x)).
// Сначала применяется g, затем f.
func compose(f, g func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		return f(g(x))
	}
}

// pipe — конвейер (слева направо): pipe(f, g)(x) = g(f(x)).
// Сначала применяется f, затем g.
func pipe(f, g func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		return g(f(x))
	}
}

func runTask4() {
	fmt.Println("=== Задание 4. Композиция и pipe ===")

	x := math.Pi / 4 // 45°

	// Через compose
	sinOfSquare := compose(sin, square) // sin ∘ square: sin(x²)
	squareOfSin := compose(square, sin) // square ∘ sin: (sin(x))²

	fmt.Printf("x = π/4 ≈ %.4f\n", x)
	fmt.Printf("compose(sin, square)(x) = sin(x²)   = %.6f\n", sinOfSquare(x))
	fmt.Printf("compose(square, sin)(x) = (sin(x))² = %.6f\n", squareOfSin(x))

	fmt.Println()

	// Через pipe (то же самое, порядок аргументов обратный)
	squareThenSin := pipe(square, sin) // сначала square, потом sin → sin(x²)
	sinThenSquare := pipe(sin, square) // сначала sin, потом square → (sin(x))²

	fmt.Printf("pipe(square, sin)(x)    = sin(x²)   = %.6f\n", squareThenSin(x))
	fmt.Printf("pipe(sin, square)(x)    = (sin(x))² = %.6f\n", sinThenSquare(x))
}
