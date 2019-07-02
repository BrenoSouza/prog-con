package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	number := 5
	ch := make(chan int)
	gateway(ch, number)
}

/**
Funçaõ que inicia num_replicas goroutines e cada thread goroutines a função request
*/
func gateway(ch chan int, num_replicas int) {
	for i := 0; i < num_replicas; i++ {
		go request(ch)
		fmt.Println(<-ch)
	}
}

/**
Sorteia números aleatórios entre 1 e 30, dormir pelo tempo do número e adiciona no channel
*/
func request(ch chan<- int) {
	number := generateContent()
	// time.Duration pra converter do tipo int pra duration
	time.Sleep(time.Duration(number) * time.Second)
	ch <- number
}

/**
Função criada para gerar um número inteiro aleatório entre 1 e 30
@return inteiro entre 1 e 30
*/
func generateContent() int {
	// Se usar só rand de int, ele gera sempre o msm número, pois topLevel functions compartilham um source que gera valores deterministicos
	// Desse jeito abaixo ele gera um número baseado no tempo em que foi executado
	rand.Seed(time.Now().UnixNano())
	// Tem que somar o mínimo, porque por default o valor mínimo gerado é 0
	min := 1
	// Será que o intervalo é aberto no limite superior?
	max := 30
	number := rand.Intn(max-min) + min

	return number
}
