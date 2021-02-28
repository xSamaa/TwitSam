package bd

import (
	"context"
	"time"

	"github.com/xSamaa/twitsam/models"
)

/*InsertoRelacion para insertar una relacion en la bd */
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitsam")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
