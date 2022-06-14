/*
Structs
A struct é uma coleção de campos.
*/

package main

import "fmt"

type Vertex struct {
	X	string
	Y	int
}

func main () {
	fmt.Println(Vertex{"Rafaella", 23})
}