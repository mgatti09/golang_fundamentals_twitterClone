package main

/* CHANNELS similar al await de NodeJS
Son unos canales que permiten que una GoRoutine envíe información hacia otra función como el main u otra
GoRoutine.
Para que nos permite enviar información, precisamente para que cada ejecución paralela que se este
ejecutando en el procesador pueda dialogar con otra porque sino tendremos islas y no tendremos control
por ejemplo puede terminar la ejecución de un programa sin que se complete una GoRoutine y esto genera
un descontrol al no tener control del programa

Un CHANNEL es un espacio de memoria de dialogo en el que van a dialogar distintas rutinas. Los canales
transportan un tipo de dato.
*/

import (
	"fmt"
	"time"
)

func main(){
	canal1 := make(chan time.Duration)	
	go bucle(canal1)
	fmt.Println("Llegué hasta aquí")

	//Para que esto sea como un await hay que decirle al programa que hay una routine corriendo
	//Para decir que estoy escuchando al canal 1 se hace lo siguiente	
	msg := <- canal1
	// hasta que msg no tiene el valor de canal1, el sistema se detiene, similar al await
	
	fmt.Println(msg)
}

func bucle(canal1 chan time.Duration){
	inicio := time.Now()
	for i:= 0; i < 10000000000; i++ {		
	}
	final:= time.Now()

	//Sub es una función que retorna la duración y que esta duración a su vez es un tipo de dato
	// es por ello que el canal se declaró de tipo time.Duration
	// <- con esto se asigna un valor al canal
	canal1 <- final.Sub(inicio)
}