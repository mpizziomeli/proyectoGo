package main

import (
	"flag"
	"github.com/mpizziomeli/proyectoGo/proyectoGo/migration"
	"log"
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
}
