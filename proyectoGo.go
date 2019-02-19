package main

import (
	"flag"
	"github.com/mpizziomeli/proyectoGo/proyectoGo/migration"
	"github.com/mpizziomeli/proyectoGo/proyectoGo/routes"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Creación de tablas")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Empezó la creación de las tablas")
		migration.Migrate()
		log.Println("Finalizó la creación de las tablas")
	}

	//Iniciar rutas
	router := routes.InitRoutes()

	//Iniciar middleware
	n:= negroni.Classic()
	n.UseHandler(router)

	//Iniciar servidor
	server := &http.Server{
		Addr: ":8080",
		Handler:n,
	}

	log.Println("Iniciando el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución")
}
