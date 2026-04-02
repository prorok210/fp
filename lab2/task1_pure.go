package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Оригинальная (нечистая) функция.
// Не является чистой, так как:
//   - читает файл (побочный эффект, обращение к внешнему миру)
//   - выводит результат на экран (побочный эффект)
//   - возвращает разные результаты при изменении файла (недетерминирована вне кода)
func SumNumbersOriginal(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		sum += number
	}
	fmt.Println("Сумма чисел (оригинал):", sum)
}

// Чистая функция — принимает срез чисел, возвращает их сумму.
// Детерминирована и не имеет побочных эффектов.
func SumNumbers(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Функции ввода-вывода отделены от логики.

func ReadNumbers(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers, scanner.Err()
}

func runTask1() {
	fmt.Println("=== Задание 1. Чистые функции ===")

	// Создаём тестовый файл
	f, err := os.Create("numbers.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	for i := 1; i <= 1000000; i++ {
		fmt.Fprintln(f, i)
	}
	f.Close()

	// Оригинальная версия
	start := time.Now()
	SumNumbersOriginal("numbers.txt")
	fmt.Println("Время оригинала:", time.Since(start))

	// Рефакторинговая версия
	start = time.Now()
	numbers, err := ReadNumbers("numbers.txt")
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}
	sum := SumNumbers(numbers)
	fmt.Println("Сумма чисел (чистая):", sum)
	fmt.Println("Время чистой версии:", time.Since(start))

	// Пояснение
	fmt.Println()
	fmt.Println("SumNumbers — чистая функция:")
	fmt.Println("  - принимает []int, возвращает int")
	fmt.Println("  - не обращается к файловой системе или stdout")
	fmt.Println("  - при одинаковых входных данных всегда даёт одинаковый результат")
	fmt.Println("  - легко тестируется, переиспользуется в любом контексте")
}
