package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Meu nome é"
	a[1] = "Rafaella Cruz"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [7]int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(primes)
}
