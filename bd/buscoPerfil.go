package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/xSamaa/twitsam/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil busca el perfil en la BD */
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twitsam")
	col := db.Collection("users")

	var perfil models.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}
	return perfil, nil
}
