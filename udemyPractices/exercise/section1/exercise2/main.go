/*
Debemos implementar un programa en go para ir ingresando valores por consola hasta que se agrega cero y finaliza el programa.

Todos los valores ingresados por consola se deben agregar a un array e imprimirlo por pantalla al finalizar.

Por ejemplo:

Ejecutamos el archivo e ingresamos los siguientes valores por consola
> go run main.go
> 3
> 4
> 5
> 0

El programa deberia devolverme
Los valores del array son:  [3 4 5]

Una vez terminado podes revisarlo en el repositorio que estamos trabajando, dentro del folder  exercise/section1/exercise2/main.go
*/

package main

import (
	"fmt"
)

func main() {
	var userInput int
	var myArray []int

	for {
		fmt.Scanln(&userInput)
		if userInput != 0 {
			myArray = append(myArray, userInput)
		} else {
			fmt.Println("Array: ", myArray)
			break
		}
	}

}
