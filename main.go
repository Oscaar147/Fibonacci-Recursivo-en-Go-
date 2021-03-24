package main

import "fmt"

func fibonacci(pos int) int {
	if pos > 1 {

		return fibonacci(pos-1) + fibonacci(pos-2)

	} else {
		if pos == 0 {
			return 0

		}
		if pos == 1 {
			return 1
		}
	}

	return 1
}

func main() {
	var p int
	fmt.Println("Ingrese posicion de fibonacci")
	fmt.Scan(&p)
	fmt.Println("El numero de la serie de fibonacci es: ", fibonacci(p))

}
