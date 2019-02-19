package routes

import (
	"github.com/gorilla/mux"
	"github.com/mpizziomeli/proyectoGo/proyectoGo/controllers"
	"github.com/urfave/negroni"
)

func SetUserRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(false)
	subRouter.HandleFunc("/", controllers.UserCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
