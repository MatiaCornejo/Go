package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(12))
}

/*
para funciones cortas se define el retorno en el inicio de la funcion
este tipo de retorno se denomina naked
*/
