package routes

import (
	"github.com/gorilla/mux"
	"github.com/mpizziomeli/proyectoGo/proyectoGo/controllers"
)

func SetLoginRouter(router *mux.Router) {
	prefix := "/api/login"
	router.HandleFunc(prefix, controllers.Login).Methods("POST")

}

