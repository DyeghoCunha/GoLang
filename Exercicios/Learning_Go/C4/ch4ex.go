package main

import "fmt"

type Animal struct {
	Name  string
	Idade int
	isSom bool
}

func (c *Animal) setName(name string) {
	c.Name = name
}
func (c *Animal) setIdade(idade int) {
	c.Idade = idade
}
func (c Animal) emiteSom() bool {
	return c.isSom
}

type Late interface {
	late() string
}

type Mia interface {
	mia() string
}

func (c Animal) late() {
	fmt.Println("Late")
}
func (c Animal) mia() {
	fmt.Println("Mia")
}

func main() {

	cachorro := Animal{
		Name:  "Cachorro",
		Idade: 10,
		isSom: true,
	}
	gato := Animal{
		Name:  "Gato",
		Idade: 5,
		isSom: true}

	borboleta := Animal{
		Name:  "Borboleta",
		Idade: 1,
		isSom: false,
	}

	fmt.Println(cachorro)
	fmt.Println(gato)
	fmt.Println(borboleta)

	cachorro.setName("Rex")
	cachorro.setIdade(7)
	fmt.Println(cachorro)
	gato.setName("Mia")
	gato.setIdade(3)
	fmt.Println(gato)
	borboleta.setName("Borbo")
	borboleta.setIdade(2)
	fmt.Println(borboleta)

	if cachorro.emiteSom() {
		fmt.Println("O cachorro emite som.")
		cachorro.late()
	} else {
		fmt.Println("O cachorro não emite som.")

	}
	if gato.emiteSom() {
		fmt.Println("O gato emite som.")
		gato.mia()
	} else {
		fmt.Println("O gato não emite som.")
	}
	if borboleta.emiteSom() {
		fmt.Println("A borboleta emite som.")
	} else {
		fmt.Println("A borboleta não emite som.")
	}
}
