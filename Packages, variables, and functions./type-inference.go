/*
Ao declarar uma variável, sem especificar o seu tipo (usando var sem um tipo ou na sintaxe :=), o tipo da variável é inferida a partir do valor do lado direito.

Quando o lado direito da declaração é digitado, a nova variável é do mesmo tipo:

var i int
j := i // j é um int

*/

package main

import "fmt"

func main() {
	v := true 
	fmt.Printf("v is of type %T\n.", v)
}