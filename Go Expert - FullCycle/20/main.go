package main

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	soma := Soma(m)

	println(soma)
}
