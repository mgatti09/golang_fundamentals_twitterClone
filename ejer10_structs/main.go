package main

/* STRUCTS
En GO no existe la POO perse. No hay declaraciones de clases, objetos, herencia, polimorfismo, etc.
Go implementa una nueva forma de hacer POO, se basa en los STRUCTS.
*/

import (
	"fmt"
	"time" //para manejo tipos de datos fecha y fechahora

	us "go-multitenancy/user" //importando nuestro paquete
)

//Emulando herencia con Go
type pepe struct{
	us.Usuario
}

func main(){
	u := new(pepe)
	u.AltaUsuario(1,"Marcos Gatti", time.Now(), true)
	fmt.Println(u.Usuario)
}