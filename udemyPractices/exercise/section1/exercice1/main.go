/*
Dentro de nuestro código ya tenemos definido un array, debemos reccorrerlo e incrementar todos sus valores en 20.
Al finalizar el programa se debera visualizar el array con los valores incrementados

Ejemplo:

Si ingresamos un vector con los valores [4, 2, 5, 6, 7] se debe imprimir el siguiente mensaje:
Los valores del array son:  [24 22 25 26 27]
*/

package main

import "fmt"

func main() {

	array := [5]int{4, 2, 5, 6, 7}

	for i := 0; i < len(array); i++ {
		array[i] += 20
	}

	fmt.Println("Los valores del array son: ", array)
}
