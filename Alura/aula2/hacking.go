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
	//var contaDoGuilherme ContaCorrente = ContaCorrente{
	//	titular:       "Guilherme",
	//	numeroAgencia: 589,
	//	numeroConta:   123456,
	//	saldo:         125.5,
	//}

	//contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	//fmt.Println(contaDoGuilherme, contaDaBruna)g
	//
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	fmt.Println(contaDaCris == contaDaCris2)
	fmt.Println(contaDaCris)
	fmt.Println(contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)
	fmt.Println(&contaDaCris)
	fmt.Println(&contaDaCris2)
	fmt.Println(&contaDaCris == &contaDaCris2)
	//
	//contaDaValter := &ContaCorrente{}
	//contaDaValter.titular = "Valter"
	//
	//fmt.Println(contaDaValter)

}
