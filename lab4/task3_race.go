package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func runTask3() {
	fmt.Println("=== Задание 3. Проблемы конкурентного программирования ===")

	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение счётчика:", counter)
}
