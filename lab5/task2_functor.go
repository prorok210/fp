package main

import "fmt"

// --- Задание 2. Функторы ---

// Box — обобщённая структура-контейнер
type Box[T any] struct {
	value T
}

// NewBox — конструктор контейнера
func NewBox[T any](v T) Box[T] {
	return Box[T]{value: v}
}

// Get — возвращает содержимое контейнера
func (b Box[T]) Get() T {
	return b.value
}

// Map — функция для функтора. Применяет f к содержимому Box[A] и возвращает Box[B].
func Map[A, B any](box Box[A], f func(A) B) Box[B] {
	result := f(box.Get())
	return NewBox(result)
}

func runTask2() {
	fmt.Println("\n=== Задание 2. Функторы ===")

	// Создаём ящик с камнем
	stoneBox := NewBox[Stone](Quartzite{Name: "Большой кварцит"})
	fmt.Printf("Создан ящик: %+v\n", stoneBox)

	// Применяем морфизм Дробить через Map (превращаем Box[Stone] в Box[Sand])
	sandBox := Map(stoneBox, func(s Stone) Sand {
		return s.Crush()
	})
	fmt.Printf("После Map(Crush) ящик содержит: %T (%+v)\n", sandBox.Get(), sandBox.Get())

	// Применяем морфизм Плавить через Map (превращаем Box[Sand] в Box[Glass])
	glassBox := Map(sandBox, func(s Sand) Glass {
		return s.Melt()
	})
	fmt.Printf("После Map(Melt) ящик содержит: %T (%+v)\n", glassBox.Get(), glassBox.Get())
}
