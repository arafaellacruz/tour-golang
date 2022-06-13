/*
For é o "while" de Go
Nesse ponto, você pode tirar as vírgulas: o while do C é escrito com for em Go.
*/

package main

import "fmt"

func main() {
	sum := 7
	for sum < 590 {
		sum += sum
	}
	fmt.Println(sum)
}