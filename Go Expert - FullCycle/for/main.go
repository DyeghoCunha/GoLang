package main

func main() {
	//No Go temos apenas o for

	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}

	// for com while
	i := 0
	for i < 10 {
		println(i)
		i++
	}
	//loop infinito
	for {
		println("loop infinito")
		break
	}
}
