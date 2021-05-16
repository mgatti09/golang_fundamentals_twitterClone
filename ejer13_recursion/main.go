package main

import "fmt"

func main(){
	exponente(2)
}

func exponente(numero int){
	if numero > 1000000000000000 {
		return
	}
	fmt.Println(numero)
	exponente(numero*2)
}