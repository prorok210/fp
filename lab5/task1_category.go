package main

import "fmt"

// --- Задание 1. Категории ---

// Интерфейсы моделируют объекты категории через исходящие морфизмы.

type Stone interface {
	Crush() Sand
	Identity() Stone
}

type Sand interface {
	Melt() Glass
	Identity() Sand
}

type Glass interface {
	Identity() Glass
}

// --- Конкретные реализации объектов ---

type Quartzite struct {
	Name string
}

func (q Quartzite) Crush() Sand {
	fmt.Println("Дробим", q.Name, "в песок...")
	return QuartzSand{Name: "Кварцевый песок из " + q.Name}
}

func (q Quartzite) Identity() Stone {
	return q
}

type QuartzSand struct {
	Name string
}

func (s QuartzSand) Melt() Glass {
	fmt.Println("Плавим", s.Name, "в стекло...")
	return SheetGlass{Name: "Листовое стекло из " + s.Name}
}

func (s QuartzSand) Identity() Sand {
	return s
}

type SheetGlass struct {
	Name string
}

func (g SheetGlass) Identity() Glass {
	return g
}

func runTask1() {
	fmt.Println("=== Задание 1. Категории ===")
	var stone Stone = Quartzite{Name: "Крепкий кварцит"}

	// Морфизм Дробить
	sand := stone.Crush()

	// Морфизм Плавить
	glass := sand.Melt()

	// Тождественные морфизмы
	fmt.Printf("Результат: %T, Значение: %+v\n", glass.Identity(), glass.Identity())
}
