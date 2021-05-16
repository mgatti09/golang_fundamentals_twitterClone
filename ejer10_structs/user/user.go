package user

import "time"

type Usuario struct{
	Id 			int
	Nombre 		string
	FechaAlta 	time.Time
	Status 		bool
}

//Función que va a operar como método de asignación de valores
//La función impacta al struct Usuario por tanto para hacer referencia a ese Struct se coloca 
//(this *Usuario) this viene siendo un puntero a Usuario
func (this *Usuario) AltaUsuario(id int, nombre string, fechaAlta time.Time, status bool){
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaAlta
	this.Status = status
}