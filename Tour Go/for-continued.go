package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}


/*
valor inicial y ejecucion final son opcionales
esto es como un while

*/
