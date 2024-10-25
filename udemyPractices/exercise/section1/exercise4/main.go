/*
Se debe implementar un archivo de go, en el cual al ejecutarlo se deben ir agregando codigos por consola, los codigos son los siguientes:

"10" : "notebook"
"15" : "tv"
"21" : "heladera"
"27" : "monitor"
"35" : "camara"

Cada vez que voy agregando codigos por consola se debe ir agregando la descripcion a un array, si agrego un codigo que no existe se debe agregar el texto "No encontrado" al array.
Al tipear 0 (cero) se debe finalizar la ejecución.

Por ejemplo, si ejecuto el archivo e ingreso los siguientes valores por pantalla:
> go run main.go
> 10
> 5
> 15
> 0
El programa deberá mostrar por pantalla los valores del array
Los valores del array son:  [notebook No encontrado tv]
*/

package main

import "fmt"

func main() {
	var userInputCode int
	var description []string
	code := map[int]string{
		10: "notebook",
		15: "tv",
		21: "heladera",
		27: "monitor",
		35: "camera",
	}

	for {
		fmt.Println("Input Code")
		fmt.Scanln(&userInputCode)

		e, ok := code[userInputCode]

		if userInputCode == 0 {
			fmt.Println(description)
			break
		}

		if ok {
			description = append(description, e)
		} else {
			description = append(description, "No encontrado")
		}
	}
}
