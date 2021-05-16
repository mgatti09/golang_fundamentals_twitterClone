package main

/* INTERFACES
Definir conductas, operaciones, comportamientos
En la interfaz se definen los métodos que se van usar para implementar esa interfaz
Una interfaz me permite agrupar en un tipo determinado de comportamiento a muchos objetos que pueden o no
tener relación entre si
*/

import "fmt"

type SerVivo interface{
	estaVivo() bool
}
type humano interface{
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool //Al colocar esto entonces humano pertenece a SerVivo
}

type animal interface{
	respirar()
	comer()
	esCarnivoro() bool
	estaVivo() bool  //Al colocar esto entonces animal pertenece a SerVivo
}

type vegetal interface{
	ClasificacionVegetal() string
}

/* Genero humano */
type hombre struct{
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	esHombre bool
	vivo bool
}

type mujer struct{
	hombre
}

//Hombre implementado la inferfaz humano
func (h *hombre) respirar() {h.respirando = true}
func (h *hombre) comer() {h.comiendo = true}
func (h *hombre) pensar() {h.pensando = true}
func (h *hombre) sexo() string {
	if h.esHombre==true{
		return "Hombre"
	}else{
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool {return h.vivo}

func HumanosRespirando(hu humano){
	hu.respirar()
	fmt.Printf("Soy un(a) %s, y ya estoy respirando \n", hu.sexo())
}

/* ---------------------------- */
/* Genero Animal */
type perro struct{
	respirando bool
	comiendo bool
	carnivoro bool
	vivo bool
}

func (p *perro) respirar() {p.respirando = true}
func (p *perro) comer() {p.comiendo = true}
func (p *perro) esCarnivoro() bool {return p.carnivoro}
func (p *perro) estaVivo() bool {return p.vivo}

func AnimalesRespirar(an animal){
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal) int {
	if an.esCarnivoro() {
		return 1
	}
	return 0
}

/* Ser Vivo */
func estoyVivo(v SerVivo) bool{
	return v.estaVivo()
}

func main(){
	Pedro := new(hombre)
	Pedro.esHombre = true
	Pedro.vivo = true
	Pedro.vivo=true
	//Aunque recibe una clase hombre y la funcion espera clase humano no falla porque hombre 
	//implememto la interfaz de humano
	HumanosRespirando(Pedro)

	Maria := new(mujer)
	HumanosRespirando(Maria)
	fmt.Printf("Estoy vivo = %t\n", estoyVivo(Maria))

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro=true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros=+ AnimalesCarnivoros(Dogo)
	fmt.Printf("Total carnivoros %d \n", totalCarnivoros)

	fmt.Printf("Estoy vivo = %t\n", estoyVivo(Dogo))
}