package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/xSamaa/twitsam/bd"
)

/*ObtenerBanner envia el avatar al HTTP  */
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado"+err.Error(), http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploards/banner" + perfil.Banner)

	if err != nil {
		http.Error(w, "Imagen banner no encontrada"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen banner"+err.Error(), http.StatusBadRequest)
		return
	}

}
