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

func main() {
	number := 5
	gateway(number)
}

/**
Funçaõ que inicia num_replicas goroutines e cada thread goroutines a função request
*/
func gateway(num_replicas int) {
	ch := make(chan int)

	for i := 0; i < num_replicas; i++ {

		fmt.Printf("%d goroutine começou\n", i+1)

		go request(ch, i)
	}

	fmt.Println(<-ch)

	// select {
	// case req := <-ch:
	// 	fmt.Println(req)
	// }

}

/**
Sorteia números aleatórios entre 1 e 30, dormir pelo tempo do número e adiciona no channel o número
*/
func request(ch chan<- int, index int) {
	number := generateRandomNumber()
	// time.Duration pra converter do tipo int pra duration
	time.Sleep(time.Duration(number) * time.Second)

	fmt.Printf("%d goroutine terminou\n", index+1)

	ch <- number
}

/**
Função criada para gerar um número inteiro aleatório entre 1 e 30
@return inteiro entre 1 e 30
*/
func generateRandomNumber() int {
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
