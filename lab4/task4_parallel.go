package main

import (
	"fmt"
	"math"
	"runtime/debug"
	"sync"
	"time"
)

// MapParallel разбивает срез data на K частей и обрабатывает их параллельно.
// Поскольку каждая горутина записывает результат по уникальному индексу,
// состояние гонки (data race) не возникает (ресурсы разделены логически).
func MapParallel(data []float64, K int, f func(float64) float64) []float64 {
	N := len(data)
	result := make([]float64, N)
	var wg sync.WaitGroup

	if K <= 0 {
		K = 1
	}

	chunkSize := (N + K - 1) / K

	for w := 0; w < K; w++ {
		start := w * chunkSize
		if start >= N {
			break
		}
		end := start + chunkSize
		if end > N {
			end = N
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				result[i] = f(data[i])
			}
		}(start, end)
	}

	wg.Wait()
	return result
}

func sinCos2(x float64) float64 {
	return math.Pow(math.Sin(x), 2) + math.Pow(math.Cos(x), 2)
}

func runTask4() {
	fmt.Println("=== Задание 4. Параллельная обработка данных ===")

	// Отключаем GC для чистых тестов
	debug.SetGCPercent(-1)

	N := 33554432 // 2^25 элементов (~268 МБ)
	data := make([]float64, N)
	for i := 0; i < N; i++ {
		data[i] = float64(i)
	}

	fmt.Printf("Объём данных: %d элементов\n", N)
	fmt.Println("Кол-во горутин | Время (мс) | Ускорение")
	fmt.Println("---------------------------------------")

	var baseTime time.Duration

	// Тестируем от 1 до 32 горутин
	goroutines := []int{1, 2, 4, 8, 16, 32}
	for _, K := range goroutines {
		start := time.Now()
		_ = MapParallel(data, K, sinCos2)
		elapsed := time.Since(start)

		if K == 1 {
			baseTime = elapsed
			fmt.Printf("%14d | %10d | %8.2fx\n", K, elapsed.Milliseconds(), 1.0)
		} else {
			speedup := float64(baseTime) / float64(elapsed)
			fmt.Printf("%14d | %10d | %8.2fx\n", K, elapsed.Milliseconds(), speedup)
		}
	}

	// Возвращаем сборщик мусора
	debug.SetGCPercent(100)
}
