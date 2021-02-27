package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xSamaa/twitsam/bd"
	"github.com/xSamaa/twitsam/jwt"
	"github.com/xSamaa/twitsam/models"
)

/*Login realiza el login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email es obligatorio", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if existe == false {
		http.Error(w, "Usuario y/o contraeña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Error al generar el token "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// Guardar en la cockie desde back.
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
