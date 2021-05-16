package main

/* Middlewares son INTERCEPTORES
que permiten ejecutar instrucciones comunes a varias funciones que reciben y devuelven los mismos
tipos de variables
Condición las funciones encadenadas por medio del Middleware deben recibir los mismos tipos de datos
y retornar los mismos tipos de datos.
*/

import "fmt"

var result int

func main(){
	fmt.Println("inicio")
	result = operacionesMidd(sumar)(2,3)
	fmt.Println(result)
	
	result = operacionesMidd(restar)(10,6)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(2,4)
	fmt.Println(result)
}

func sumar(a,b int) int {
	return a+b
}

func restar(a,b int) int {
	return a+b
}

func multiplicar(a,b int) int {
	return a*b
}

//Lo que recibe y lo que devuelve el middleware debe ser exactamente igual
func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(i1, i2 int) int {
		fmt.Println("Inicio de operación")
		return f(i1,i2)
	}
}