package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

func leoArchivo(){
	//ioutil.ReadFile se encarga de abrir el archivo, leer el buffer y cerrar el archivo.
	//Lee todo el archivo
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	}else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2(){
	// Variante usando os para recorrer linea por linea del archivo txt
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	}else {
		scanner := bufio.NewScanner(archivo)
		//Leyendo linea a linea hasta el EOF
		for scanner.Scan() {
			linea := scanner.Text()
			fmt.Printf("Linea > "+linea+"\n")
		}
	}
	archivo.Close()
}

func graboArchivo(){
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	//Print ln sobre un archivo, comienza con una F
	fmt.Fprintln(archivo, "Esta es una línea nueva")
	archivo.Close()
}

func graboArchivo2(){
	//Adicionando texto a un archivo ya existente
	fileName := "./nuevoArchivo.txt"
	if !Append(fileName, "\nuna segunda línea") {		
		fmt.Println("Error en la segunda línea")
		return
	}
}

func Append(archivo string, texto string) bool {
	//os.O_WRONLY indica que es para lectura y escritura
	//os.O_APPEND indica que se escribira al final del archivo
	// 0644 permisos de lectura escritura
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error en la escritura del archivo")
		return false
	}
	return true
}