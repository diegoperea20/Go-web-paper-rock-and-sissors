package main

import (
	"log"
	"net/http"
	"webgopaper/handlers"
)


func main() {
	
	//Crear enrutador 
	router :=http.NewServeMux()

	//manejador de archivos estaticos
	fs := http.FileServer(http.Dir("./static"))
	//para acceder a los archivos estaticos
	router.Handle("/static/", http.StripPrefix("/static/", fs))




	//Registrar rutas
	router.HandleFunc("/",handlers.Index)
	router.HandleFunc("/new",handlers.NewGame)
	router.HandleFunc("/game",handlers.Game)
	router.HandleFunc("/play",handlers.Play)
	router.HandleFunc("/about",handlers.About)


	port:=":8080"
	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))

}