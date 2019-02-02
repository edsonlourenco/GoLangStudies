package main

import "fmt"

//Tipos mais comuns em Go Lang
var (
	//Endereco é um valor importante e tem de ser publico
	Endereco string
	//Telefone é um valor importante para nossa aula
	Telefone            string  // Variável Endereco é publica, telefone é privada
	quantidade, estoque int     //quantidade = 0
	comprou             bool    //comprou = false
	valor               float64 //valor = 0.00
	palavras            rune
)

func main() {

	teste := "Valor de testes"
	fmt.Println("Endereco: %s\r\n", Endereco)
	fmt.Println("Quantidade: %d\r\n", quantidade)
	fmt.Println("Comprou: %v\r\n", comprou)
	fmt.Println("palavras: %v\r\n", palavras)
	fmt.Println("O valor de teste é: %s\r\n", teste)

}
