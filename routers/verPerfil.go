package routers

import (
	"encoding/json"
	"net/http"

	"github.com/xSamaa/twitsam/bd"
)

/*VerPerfil permite extraer los valores del perfil */
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el regsitro"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "applucation/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
