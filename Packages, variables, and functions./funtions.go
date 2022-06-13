/*
Funções
A função pode ter zero ou mais argumentos.

Neste exemplo, adicione (add) dois parâmetros do tipo int.

Observe que o tipo vem após o nome da variável.
*/

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main () {
	fmt.Println(add(42, 13))
}