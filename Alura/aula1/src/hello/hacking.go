package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	leSitesDoArquivo()

	for {
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
	fmt.Println("0- Sair do Programa \n")
}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//	sites := []string{"https://www.alura.com.br",
	//	"https://random-status-code.herokuapp.com/",
	//	"https://www.uol.com.br",
	//	"https://www.caelum.com.br"}
	sites := leSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			//fmt.Println(site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("-------------------------------------\n")
	}
}
func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	switch resp.StatusCode {
	case 200:
		fmt.Println("Site", site, "foi carregado com sucesso")
	default:
		fmt.Println("Site", site, "foi carregado com falha. Status Code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	//	arquivo, err := os.Open("sites.txt")
	//	arquivo, err := ioutil.ReadFile("C:\\Users\\Dyegho\\Documents\\GitHub\\GoLang\\Alura\\aula1\\src\\hello\\sites.txt")

	arquivo, err := os.Open("C:\\Users\\Dyegho\\Documents\\GitHub\\GoLang\\Alura\\aula1\\src\\hello\\sites.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo", err)
	}
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		//fmt.Println(linha)
		sites = append(sites, linha)

		if err == io.EOF {

			break
		}

	}
	fmt.Println(sites)
	return sites
}
