package main

import (
	"Alura/contas"
	"fmt"
)

// Importa os pacotes fmt e os
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaCorrente{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.ObterSaldo())
	contaDaLuisa := contas.ContaPoupanca{}
	contaDaLuisa.Depositar(100)
	PagarBoleto(&contaDaLuisa, 60)
	fmt.Println(contaDaLuisa.ObterSaldo())
}
