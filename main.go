package main

import (
	"curso/godesde0/ejercicios"
	"fmt"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	/*
		estado, texto := variables.ConviertoATexto(1573)
		fmt.Println(estado)
		fmt.Println(texto)

		if os := runtime.GOOS; os == "windows" {
			fmt.Println("Estás en Windows")
		} else if os == "linux" {
			fmt.Println("Estás en Linux")
		} else {
			fmt.Println("Sistema operativo no reconocido")
		}

		switch os := runtime.GOOS; os {
		case "windows":
			fmt.Println("Estás en Windows")
		case "linux":
			fmt.Println("Estás en Linux")
		default:
			fmt.Println("Sistema operativo no reconocido")
		}*/

	numero, texto := ejercicios.ConviertNumerico("500")

	fmt.Println(numero)
	fmt.Println(texto)
}
