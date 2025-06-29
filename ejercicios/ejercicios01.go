package ejercicios

import "strconv"

func ConviertNumerico(texto string) (int, string) {
	num, err := strconv.Atoi(texto)

	if err != nil {
		return 0, "Error al convertir el texto a número: " + err.Error()
	}

	if num > 100 {
		return num, "El número es mayor que 100"
	} else {
		return num, "El número es menor o igual a 100"
	}
}
