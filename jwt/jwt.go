package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/xSamaa/twitsam/models"
)

/*GeneroJWT genera el JWT*/
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("ClaveParaGenerarJWT")
	payLoad := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payLoad)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
