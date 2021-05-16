package main

import (
	"fmt"
	"strconv"
)	

var numero int
var texto string
var status bool

func main(){
	num5, num6, num7, text := 2 , 5, 68, "hola"
	var moneda int64 = 654654
	num5 = int(moneda)

	//Para convertir enteros en texto
	//Se debe castear a entero genero para poder usar la funcion Itoa
	text = strconv.Itoa(int(moneda))	

	fmt.Println(num5)
	fmt.Println(num7)
	fmt.Println(num6)
	fmt.Println(text)
	fmt.Println(status)
}