package routers

import (
	"net/http"

	"github.com/xSamaa/twitsam/bd"
	"github.com/xSamaa/twitsam/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Error al borrar una relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se borro la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
