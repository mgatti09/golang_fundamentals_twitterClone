package main

/*
En Go solo existe For
*/
import "fmt"

func main(){
	
	for i := 1; i < 10; i++{
		fmt.Println(i)
	}

	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	var i = 0
	for i < 10 {
		fmt.Printf("\n Valor de i: %d",i)
		if i == 5 {
			fmt.Printf(" multiplicamos por 2 \n")
			i = i * 2 
			continue
		}
		fmt.Printf("      Paso por aqui \n")
		i++
	}

	// Ejemplo de usar Go to
	var j int = 0

	RUTINA: 
		for j < 10 {
			if j == 4 {
				j=j+2
				fmt.Println("Voy a RUTINA sumando 2 a j")
				goto RUTINA
			}
			fmt.Printf("Valor de j: %d\n",j)
			j++
		}
}