package main

import "fmt"

//Imovel representa um imóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := new(Imovel)
	fmt.Printf("Casa é:  %+v\r\n", casa)

	chacara := Imovel{17, 28, "Chacara Linda", 280000}
	fmt.Printf("Chacara é: %p - %v\r\n", &chacara, chacara)

	mudaImovel(&chacara)
	fmt.Printf("Chacara é: %p - %v\r\n", &chacara, chacara)

}

func mudaImovel(imovel *Imovel) {

	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5

}
