package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}


/*
Para declarar funciones se escribe func
datos y tipo y el tipo de retorno
ej datos

x int
p *int
a [3]int
FUENTE:https://blog.golang.org/declaration-syntax
*/
