package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// For assim é um while true
	// for {
	// fmt.Println("funciona")
	// }
	ch := make(chan int)
	generateContent(ch)
}

//Sorteia números aleatórios entre 1 e 30, dormir pelo tempo do número e retorná-lo
// int request() {

// }

func generateContent(out chan int) {
	//Se usar só rand de int, ele gera sempre o msm número, pois topLevel functions compartilham um source que gera valores deterministicos
	// Desse jeito abaixo ele gera um número baseado no tempo que executei
	rand.Seed(time.Now().UnixNano())
	// Tem que somar o mínimo, pq por default o valor mínimo gerado é 0
	min := 1
	//Será que o intervalo é aberto no limite superior
	max := 30
	number := rand.Intn(max-min) + min
	// time.Duration pra converter do tipo int pra duration
	time.Sleep(time.Duration(number) * time.Second)
	fmt.Println(number)
}

//iniciar num_replicas goroutines, e cada thread goroutines a função request.
// int gateway(int num_replicas) {

// }
