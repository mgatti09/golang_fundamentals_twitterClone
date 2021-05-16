package main

/* Manejo de excepciones y errores
en Go no hay try catch

DEFER: instrucción que siempre se ejecuta al final de una función ya sea si
hace un return, un error o llega al final. Similar a un Finally

PANIC: forzar un error, similar a un Throw. El Panic aborta el programa. Si se desea controlar el PANIC
se hace uso del RECOVER

RECOVER: se ejecuta cuando detecta que hay un PANIC y toma el mensaje del PANIC

*/

import (
	"fmt"
	"log"
	"os"
)

func main(){
	archivo := "prueba.txt"
	f, err := os.Open(archivo)

	// esto se ejecutará al final y garantiza que cierra el archivo independiente si hay falla o no
	defer f.Close()

	if err != nil {
		fmt.Println("error abriendo el archivo")
		//os.Exit(1)
	}

	ejemploPanic()
}

func ejemploPanic(){
	//defer ejecuta una sola instrucción, si se desean varias instrucciones entonces se usa una función
	// por lo general anonima
	defer func() {
		//recover me trae el resultado del ultimo panic, si no hubo panic entonces devuelve un nil
		reco := recover()
		//Grabando en log en caso de haber ocurrido un PANIC
		if reco != nil{
			/*log FatalF realiza lo siguiente
			 - Incluye fecha y hora adicional del mensaje del error
			 - Muestra el texto por pantalla con la fecha y hora
			 - Hace un exit status 1 para finalizar el programa
			*/
			log.Fatalf("ocurrio un error que generó Panic \n %v", reco)
		}
	}() //estos parentesis del final para ejecutar de inmediato la función

	a:=1
	if a ==1 {
		panic("Se encontro el valor de 1")
	}
}