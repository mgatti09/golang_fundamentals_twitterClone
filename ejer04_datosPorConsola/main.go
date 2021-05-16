package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 Bufio paquete muy completo para lectura de archivos, por teclado. Alternativa al Scan de fmt
*/
var numero1, numero2, resultado int
var leyenda string

func main(){
	fmt.Println("Ingrese numero 1: ")
	fmt.Scanln(&numero1)
	
	fmt.Println("Ingrese numero 2: ")	
	fmt.Scanln(&numero2)

	fmt.Println("Descripción: ")

	//Aca se declara con bufio un nuevo scanner donde se indica que la entrada sera el teclado
	//standarInput siempre es el teclado, el standarOutput es la pantalla
	scanner := bufio.NewScanner(os.Stdin)
	//Se evalua	si ejecutando el metodo Scan se obtiene algo
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	resultado = numero1+numero2
	fmt.Println(leyenda,resultado)
}