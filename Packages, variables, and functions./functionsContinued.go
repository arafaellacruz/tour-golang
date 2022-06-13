/*
Funções continuação
Quando dois ou mais parâmetros nomeados consecutivos de uma função compartilham o mesmo tipo, você pode omitir o tipo de todos menos o último. 
Neste exemplo, vamos encurtar: x int, y int
para: x, y int
*/

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}