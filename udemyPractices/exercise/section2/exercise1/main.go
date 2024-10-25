/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren un package llamado matrix que:

> Tenga una función New que reciba una serie de valores de punto flotante e inicialice los en una estructura Matrix
> Que la estructura Matrix tenga un metodo Print que imprima por pantalla la matriz, las dimensiones de la matriz y si es cuadrática
> Todas las columnas deben tener el mismo tamaño, sino la funcion New debera retornar un error
> Todas las filas deben tener el mismo tamaño, sino la funcion New debera retornar un error

(la dimension de la matriz es la cantidad de columnas y filas que tiene, y si el numero de filas y columnas son iguales, es cuadratica)

La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho y si es cuadrática.

# Y una vez ejecutado debería devolver lo siguiente por pantalla

[ 2 7 8 ]
[ 4 4 7 ]
[ 5 6 1 ]
3 x 3
Cuadratic:  true
*/
package main

import (
	"fmt"

	"github.com/SanGameDev/GoProjects/udemyPractices/exercise/section2/exercise1/matrix"
)

func main() {
	m, err := matrix.New([]float64{2, 7, 8}, []float64{4, 4, 7}, []float64{5, 6, 1})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m.Print()
}
