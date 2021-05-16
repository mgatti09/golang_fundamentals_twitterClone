package main

/*
Mapa es una estructura que se puede serializar clave valor. Similar a un Diccionario en Python
*/

import (
	"fmt"
)


func main(){
	//El segundo parametro es opcional para reservar el espacio en memoria. Si no se coloca
	//Go lo asigna automáticamente
	paises := make(map[string]string, 5)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"	
	fmt.Println(paises)

	//Otra forma de declarar mapas y asignar directamente
	campeonato := map[string]int {
		"Barcelona": 39,
		"Real Madrid": 38,
		"Chivas": 37,
		"Boca Juniors": 30}
	fmt.Println(campeonato)

	//Agregando elementos al mapa
	campeonato["River Play"] = 25
	//Actualizando un elemento
	campeonato["Chivas"] = 55
	//Borrando
	delete(campeonato,"Real Madrid")

	fmt.Println(campeonato)

	//Recorriendo un mapa
	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s tiene un puntaje: %d \n", equipo,puntaje)
	}

	//Verificando si un elemento del mapa existe
	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe %t \n", puntaje, existe)

	
	puntaje, existe = campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe %t \n", puntaje, existe)
}