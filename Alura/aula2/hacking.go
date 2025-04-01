package main

import (
	"Alura/contas"
	"fmt"
)

// Importa os pacotes fmt e os

func main() {
	contaDaSilvia := contas.ContaCorrente{
		Titular:       "Silvia",
		NumeroAgencia: 0,
		NumeroConta:   0,
		Saldo:         300,
	}
	contaDoGustavo := contas.ContaCorrente{
		Titular:       "Gustavo",
		NumeroAgencia: 0,
		NumeroConta:   0,
		Saldo:         100,
	}
	status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	fmt.Println(status)
	fmt.Println("Conta do Gustavo:", contaDoGustavo)
	fmt.Println("Conta da Silvia:", contaDaSilvia)
}
