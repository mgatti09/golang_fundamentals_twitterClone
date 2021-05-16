package main

import "fmt"

func main(){
	fmt.Println(multiplicarPorDos(8))

	numero, estado := dos(1)
	fmt.Println(numero,estado)
	
	//Note que no se ponen los dos puntos dado que ya fue declarada anteriormente con el :=
	numero, estado = dos(3)
	fmt.Println(numero,estado)

	//Invocando la función que recibe multiples valores
	fmt.Println(Calculo(5,46))
	fmt.Println(Calculo(1,2,3,4,5,6,7,8,9,10))
}

func multiplicarPorDos(numero int) (z int) {
	z = numero * 2
	return 
}

func dos(numero int) (int, bool) {
	if numero == 1{
		return  5, false
	}else {
		return 10, true
	}
}

//Pasando multiples parametros en una funcion
//Calculo podra recibir una cantidad indefinida de int
func Calculo(numero ...int) int{
	total:=0
	//range siempre devuelve dos valores. Toma un un vector y permite iterar por cada uno de sus elementos
	//devuelve el indice y el valor del arreglo. Como no requiero almacenar el indice uso el _
	for _, num:= range numero {
		total = total + num
	}
	return total
}