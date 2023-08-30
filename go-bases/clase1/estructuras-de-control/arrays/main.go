package main

import "fmt"

func main() {
	var word = "Perro"
	/*cantidad := len(word)
	fmt.Println(cantidad)*/
	for i, c := range word {
		fmt.Println(string(c))
		if i == len(word)-1 {
			fmt.Println("la cantidad es:", i+1)
		}
	}
}
