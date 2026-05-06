package main

import (
	"fmt"
	"sync"
)

func runTask2() {
	fmt.Println("=== Задание 2. Знакомство с конкурентным программированием ===")

	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Println("Hello", val)
		}(i)
	}

	wg.Wait()
	fmt.Println("end")
	fmt.Println("Объяснение: Без wg.Wait() главная горутина завершалась раньше, чем успевали стартовать и выполниться дочерние горутины, поэтому ничего не выводилось.")
}
