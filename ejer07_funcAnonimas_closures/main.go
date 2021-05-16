package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 +7 = %d \n", Calculo(5, 7))

	//Vamos a redefinir Calculo
	Calculo = func(i1, i2 int) int {return i1-i2}	
	fmt.Printf("Resta 6 - 4 = %d \n", Calculo(6, 4))

	//Vamos a redefinir Calculo
	Calculo = func(i1, i2 int) int {return i1/i2}	
	fmt.Printf("Dividir 12 / 3 = %d \n", Calculo(12, 3))

	Operaciones()

	/* CLOSURES */
	tablaDel := 2
	fmt.Printf("\nMi tabla del %d \n", tablaDel)
	miTabla := Tabla(tablaDel)
	for i := 1; i<=10; i++ {
		fmt.Println(miTabla())
	}

	tablaDel = 3
	fmt.Printf("\nMi tabla del %d \n", tablaDel)
	miTabla = Tabla(tablaDel)
	for i := 1; i<=10; i++ {
		fmt.Println(miTabla())
	}

}

func Operaciones(){
	resultado := func() int {
		var a int = 23
		var b int = 14
		return a + b
	}

	fmt.Println(resultado())
}

/*
Closures: Funciones anonimas que tiene que ver con la protección e isolación de código
Pueden acceder a variables declaradas por fuera de la función original
*/

// Funcion tabla que devuelve una funcion de tipo entero
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	// La declaración de lo que se devuelve debe coincidir con la declaración que devuelve la función madre	
	return func() int {
		secuencia++
		return numero*secuencia
	}
}
