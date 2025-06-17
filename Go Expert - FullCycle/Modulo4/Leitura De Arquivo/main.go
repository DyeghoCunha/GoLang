package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//f, err := os.Create("arquivo.txt")
	//if err != nil {
	//	panic(err)
	//}
	//tamanho, err := f.WriteString("Ola mundo")
	//tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Tamanho do arquivo: %d\n", tamanho)
	//f.Close()
	//	Leitura
	//a, err := os.ReadFile("arquivo.txt")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(a))

	// leitura de linha	a linha
	arquivo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 10)
	all := make([]byte, 0)
	for {
		n, err := reader.Read(buffer)
		fmt.Println(n)
		if err != nil {
			break
		}
		all = append(all, buffer[:n]...)
		fmt.Println(string(buffer[:n]))
		fmt.Println(string(all))
	}
	fmt.Println(string(all))
}
