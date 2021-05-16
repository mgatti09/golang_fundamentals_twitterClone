package main

import (
	"fmt"
)

var estado bool

func main(){
	estado = true
	if estado {
		fmt.Println(estado)
	} else {
		fmt.Println("Es falso")
	}

	//Dentro del if se pueden realizar asignaciones antes de la evaluacion	
	if estado=false; estado {
		fmt.Println(estado)
	} else {
		fmt.Println("Es falso")
	}

	//Switch. De igual manera se puede realizar asignaciones previo a la evaluación
	//A diferencia de otros lenguajes no es necesario colocar el break por cada case
	switch numero := 5; numero {
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println(3)
		case 4:
			fmt.Println(4)
		case 5:
			fmt.Println(5)
		default:
			fmt.Println("Mayor a 5")		
	}


}