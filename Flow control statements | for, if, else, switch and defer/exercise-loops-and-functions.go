//https://go-tour-br.appspot.com/flowcontrol/8

package main

import "fmt"

func main() {
	m := 144.0 // Atribuindo e inicializando um valor a variavel
	Sqrt(m)    // Chamando e passando um parametro para a função criada abaixo.
}

// z - = (z * z - x) / (2 * z)
//
func Sqrt(x float64) {
	// z = 1
	// 1 - = (1 * 1 - 16) / 2* 1
	// 1 - = -15/2
	// 1 - = -7.5
	z := 1.0

	for i := 1; i < 10; i++ {

		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
	}
}
