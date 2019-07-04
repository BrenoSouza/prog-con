package main

/**
A função request deve sortear um número aleatório entre 1 e 30, dormir por uma quantidade de segundos dada pelo
número aleatório sorteado, e retornar o valor do número. Por sua vez, a função gateway deve iniciar num_replicas
goroutines, e cada thread goroutines a função request.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

const MIN_NUMBER = 1
const MAX_NUMBER = 30

func main() {

	num_replicas := 5
	result := gateway(num_replicas)
	fmt.Printf("Valor da primeira goroutine: %d", result)
}

/**
Funçaõ que inicia num_replicas goroutines e cada thread goroutines a função request
*/
func gateway(num_replicas int) int {
	ch := make(chan int)

	for i := 0; i < num_replicas; i++ {

		fmt.Printf("%d goroutine começou\n", i+1)

		// Passei o index da goroutine por motivos de visualização
		go request(ch, i+1)
	}

	return <-ch
}

/**
Sorteia números aleatórios entre 1 e 30, dormir pelo tempo do número e adiciona no channel o número
*/
func request(ch chan<- int, index int) {
	number := generateRandomNumber(MIN_NUMBER, MAX_NUMBER)
	// time.Duration pra converter do tipo int pra duration
	fmt.Printf("%d goroutine dormindo por %d segundos\n", index, number)

	time.Sleep(time.Duration(number) * time.Second)

	fmt.Printf("%d goroutine terminou\n", index)

	ch <- number
}

/**
Função criada para gerar um número inteiro aleatório entre 1 e 30
@return inteiro entre 1 e 30
*/
func generateRandomNumber(min int, max int) int {
	// Se usar só rand de int, ele gera sempre o msm número, pois topLevel functions compartilham um source que gera valores deterministicos
	// Desse jeito abaixo ele gera um número baseado no tempo em que foi executado
	rand.Seed(time.Now().UnixNano())
	// Tem que somar o mínimo, porque por default o valor mínimo gerado é 0
	// Será que o intervalo é aberto no limite superior?
	number := rand.Intn(max-min) + min

	return number
}
