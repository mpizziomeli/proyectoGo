package commons

import (
	"encoding/json"
	"github.com/mpizziomeli/proyectoGo/proyectoGo/models"
	"log"
	"net/http"
)

func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al convertir el mensaje: %s", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}
