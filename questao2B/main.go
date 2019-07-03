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
	ch := make(chan int, num_replicas)
	sum := 0

	for i := 0; i < num_replicas; i++ {

		fmt.Printf("%d goroutine começou\n", i+1)

		go request(ch, i+1)
	}

	timeout := time.Tick(16 * time.Second)

	for try := num_replicas; try > 0; try-- {
		select {
		case <-timeout:
			sum = -1
			break
		case req := <-ch:
			sum += req
		}
	}
	fmt.Printf("Soma total %d\n", sum)

}

/**
Sorteia números aleatórios entre 1 e 30, dormir pelo tempo do número e adiciona no channel o número
*/
func request(ch chan<- int, index int) {
	number := generateRandomNumber()
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
func generateRandomNumber() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Tem que somar o mínimo, porque por default o valor mínimo gerado é 0
	min := 10
	// Será que o intervalo é aberto no limite superior?
	max := 30
	number := r.Intn(max-min) + min

	return number
}
