package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000 ; { // A instrução inicial e a pós-instrução são opcionais.
		sum += sum
	}
	fmt.Println(sum)
}