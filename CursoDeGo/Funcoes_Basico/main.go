package main

import (
	"fmt"

	"github.com/edsonlourenco/CursoDeGo/Funcoes_Basico/matematica"
)

func main() {
	resultado := matematica.Calculo(Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)
	resultado = matematica.Soma(2, 2)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
	resultado = matematica.Divisor(2, 2)
	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("O resultado da divisão foi %d e o resto da divisão é %d", resultado, resto)
}

//Multiplicador multiplica um número * y
func Multiplicador(x int, y int) int {
	return x * y
}
