package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generarAlertas(alertas chan<- int) {
	for i := 1; i <= 30; i++ {
		select {
		case alertas <- i:
			// ya se envio el dato
		default:
			fmt.Println("Alerta descartada por saturación de canal")
		}
	}
	close(alertas)
}

func worker(id int, alertas <-chan int, reportes_limpios chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for alerta := range alertas {
		tiempo := time.Duration(rand.Intn(251)+50) * time.Millisecond
		time.Sleep(tiempo)

		reporte := fmt.Sprintf("Alerta %d validada por Worker %d", alerta, id)
		reportes_limpios <- reporte
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	alertas := make(chan int, 10)
	reportes_limpios := make(chan string)

	var wg sync.WaitGroup

	go generarAlertas(alertas)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, alertas, reportes_limpios, &wg)
	}

	go func() {
		wg.Wait()
		close(reportes_limpios)
	}()

	for reporte := range reportes_limpios {
		fmt.Println(reporte)
	}

	fmt.Println("Procesamiento finalizado correctamente.")
}
