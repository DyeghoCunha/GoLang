package main

import (
	"fmt"
	"net/http"
)
import "os"

func main() {
	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {

	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

}

func exibeIntroducao() {
	var versao float32 = 1.1
	var nome string = "Dyegho"
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp.Status)
}
