package main

import "net/http" //Todo para construir un servidor web

func main(){
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000",nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"./index.html")
}