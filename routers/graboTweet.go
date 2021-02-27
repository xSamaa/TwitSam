package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xSamaa/twitsam/bd"
	"github.com/xSamaa/twitsam/models"
)

/*GraboTweet permite grabar el tweet en la base de datos */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Error al insertar el registro del tweet"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se inserto el tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
