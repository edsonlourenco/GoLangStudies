package main

import (
	"fmt"

	"github.com/edsonlourenco/CursoDeGo/Pacotes/operadora"
	"github.com/edsonlourenco/CursoDeGo/Pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário...
var NomeDoUsuario = "Edson Lourenço"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da Operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste: %s\r\n", prefixo.TesteComPrefixo)
}
