package main

import "fmt"

// --- Задание 3. Монады ---

// Option — контейнер, который может хранить значение или быть пустым.
type Option[T any] struct {
	value T
	valid bool
}

// Some — помещает значение в контейнер (Pure)
func Some[T any](v T) Option[T] {
	return Option[T]{value: v, valid: true}
}

// None — создаёт пустой контейнер
func None[T any]() Option[T] {
	return Option[T]{valid: false}
}

// Get — возвращает содержимое контейнера и флаг успешности
func (o Option[T]) Get() (T, bool) {
	return o.value, o.valid
}

// MapOption — Map для контейнера Option (Функтор)
func MapOption[A, B any](opt Option[A], f func(A) B) Option[B] {
	if !opt.valid {
		return None[B]()
	}
	return Some(f(opt.value))
}

// FlatMap — применяет f: A -> Option[B] к Option[A], не создавая вложенности. (Монада)
func FlatMap[A, B any](val Option[A], f func(A) Option[B]) Option[B] {
	if !val.valid {
		return None[B]()
	}
	return f(val.value)
}

// --- Новые операции Дробить и Плавить, которые могут завершиться неудачно ---

type ToughStone struct {
	Name      string
	IsTooHard bool
}

// CrushSafe завершается неудачно, если камень слишком твёрдый
func (s ToughStone) CrushSafe() Option[Sand] {
	if s.IsTooHard {
		fmt.Println("Ошибка: Камень", s.Name, "слишком твёрдый для дробления!")
		return None[Sand]()
	}
	fmt.Println("Успех: Камень", s.Name, "раздроблен.")
	return Some[Sand](QuartzSand{Name: "Песок из " + s.Name})
}

// MeltSafe завершается неудачно, если температура недостаточна.
// Для упрощения передаём флаг как поле, хотя в реальности это мог бы быть аргумент среды.
type MeltableSand struct {
	Name        string
	Temperature int
}

func (s MeltableSand) MeltSafe() Option[Glass] {
	if s.Temperature < 1000 {
		fmt.Println("Ошибка: Температуры", s.Temperature, "недостаточно для плавления", s.Name)
		return None[Glass]()
	}
	fmt.Println("Успех: Песок", s.Name, "расплавлен.")
	return Some[Glass](SheetGlass{Name: "Стекло из " + s.Name})
}

func runTask3() {
	fmt.Println("\n=== Задание 3. Монады ===")

	// Сценарий 1: Всё проходит успешно
	fmt.Println("Сценарий 1: Успешная цепочка преобразований")
	stone1 := Some(ToughStone{Name: "Мягкий известняк", IsTooHard: false})

	// FlatMap(Crush)
	sandOpt1 := FlatMap(stone1, func(s ToughStone) Option[MeltableSand] {
		// Оборачиваем Sand в MeltableSand для следующего шага
		sandResult := s.CrushSafe()
		return FlatMap(sandResult, func(sand Sand) Option[MeltableSand] {
			qs := sand.(QuartzSand)
			return Some(MeltableSand{Name: qs.Name, Temperature: 1200}) // Успешная температура
		})
	})

	// FlatMap(Melt)
	glassOpt1 := FlatMap(sandOpt1, func(s MeltableSand) Option[Glass] {
		return s.MeltSafe()
	})

	if val, ok := glassOpt1.Get(); ok {
		fmt.Printf("Итог 1: %T (%+v)\n", val, val)
	}

	fmt.Println("\nСценарий 2: Ошибка на этапе дробления")
	stone2 := Some(ToughStone{Name: "Алмаз", IsTooHard: true})
	sandOpt2 := FlatMap(stone2, func(s ToughStone) Option[MeltableSand] {
		sandResult := s.CrushSafe()
		return FlatMap(sandResult, func(sand Sand) Option[MeltableSand] {
			qs := sand.(QuartzSand)
			return Some(MeltableSand{Name: qs.Name, Temperature: 1200})
		})
	})
	glassOpt2 := FlatMap(sandOpt2, func(s MeltableSand) Option[Glass] {
		return s.MeltSafe()
	})

	if _, ok := glassOpt2.Get(); !ok {
		fmt.Println("Итог 2: Преобразование прервано (None)")
	}

	fmt.Println("\nСценарий 3: Ошибка на этапе плавления")
	stone3 := Some(ToughStone{Name: "Обычный камень", IsTooHard: false})
	sandOpt3 := FlatMap(stone3, func(s ToughStone) Option[MeltableSand] {
		sandResult := s.CrushSafe()
		return FlatMap(sandResult, func(sand Sand) Option[MeltableSand] {
			qs := sand.(QuartzSand)
			return Some(MeltableSand{Name: qs.Name, Temperature: 500}) // Температуры недостаточно
		})
	})
	glassOpt3 := FlatMap(sandOpt3, func(s MeltableSand) Option[Glass] {
		return s.MeltSafe()
	})

	if _, ok := glassOpt3.Get(); !ok {
		fmt.Println("Итог 3: Преобразование прервано (None)")
	}
}
