package main

import "fmt"

//arreglo de enteros inicializado con valor 0 en sus 10 posiciones
var tabla [10]int
func main(){
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	//Otra manera de declarar un arreglo e inicializarlo
	tabla2 := [10]int{5,7,58,4,65,8,25,10,15,14}
	fmt.Println(tabla2)

	for i:= 0; i < len(tabla2); i++{
		fmt.Println(tabla2[i])
	}

	//Matrices
	var matriz[5][7]int 
	matriz[3][5] = 1
	fmt.Println(matriz)

	//Slice: Vectores dinamicos. Ampliar la longitud de los vectores en tiempo real.
	matriz2 := []int{2,5,4}
	fmt.Println(matriz2)
	variante2()
	variante3()
	variante4()
}

func variante2(){
	elementos := [5]int{1,2,3,4,5}
	//GO permite crear un slice con base a un vector
	porcion := elementos[3:]
	fmt.Println(porcion)
}

func variante3(){
	//Crear slices para poder adicionar elementos durante la ejecución
	//Primer argumento: el tipo de vector
	//Segundo argumento: Longitud inicial del vector
	//Tercer argumento: espacio máximo reservado en memoria. Capacidad de un slice
	elementos := make([]int,5,20)
	//La función cap me permite conocer la capacidad del vector
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func variante4(){
	nums := make([]int, 0, 0)
	for i:= 0; i<100; i++ {
		//El append resuelve el aumentar la capacidad sin embargo tiene un costo computacional
		//El programador debe monitorear la capacidad y cuando se vea que se va a rebasar entonces usar
		//append
		nums = append(nums, i)
	}
	//Cuando se supera la capacidad la incrementa en 2^n
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}