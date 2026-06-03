package main

import (
	"fmt"
	"sync"
)

func worker(id int, resultados chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	resultados <- fmt.Sprintf("Resultado de tarea %d", id)
}

func main() {
	resultados := make(chan string)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, resultados, &wg)
	}

	go func() {
		wg.Wait()
		close(resultados)
	}()

	for res := range resultados {
		fmt.Println(res)
	}
}
