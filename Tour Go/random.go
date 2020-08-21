package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(231)
	//siempre ejecuta el misma valor, para que sea distinto se debe cambiar la semilla
	fmt.Println("My favorite number is", rand.Intn(10))
}
