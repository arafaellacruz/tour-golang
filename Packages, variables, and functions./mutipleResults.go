/*
Resultados Múltiplos
Uma função pode retornar qualquer número de resultados.

A função swap retorna duas strings.
*/

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("when im away from you", "im happier than ever")
	fmt.Println(a, b)
}
