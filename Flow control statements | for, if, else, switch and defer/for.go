// Go tem apenas uma estrutura de laço, o for.
/*
O laço for básico tem três componentes separados por ponto e vírgula:

	a instrução inicial: executada antes da primeira iteração
	a expressão de condição: avaliada antes de cada iteração
	a pós-instrução: executado no final de cada iteração
*/

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum+= i
	}
	fmt.Println(sum)
}