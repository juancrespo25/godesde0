package main

import (
	"curso/godesde0/variables"
	"fmt"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()

	estado, texto := variables.ConviertoATexto(1573)
	fmt.Println(estado)
	fmt.Println(texto)
}
