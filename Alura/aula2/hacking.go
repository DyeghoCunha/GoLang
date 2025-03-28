package main

import "fmt"

// Importa os pacotes fmt e os

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	var contaDoGuilherme ContaCorrente = ContaCorrente{
		titular:       "Guilherme",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}
	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	fmt.Println(contaDoGuilherme, contaDaBruna)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"

	contaDaValter := &ContaCorrente{}
	contaDaValter.titular = "Valter"

	fmt.Println(contaDaValter)

}
