package bd

import (
	"context"
	"time"

	"github.com/xSamaa/twitsam/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores lee los tweets de los seguidores */
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitsam")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condicion := make([]bson.M, 0)
	condicion = append(condicion, bson.M{"$match": bson.M{"usuarioid": ID}})
	condicion = append(condicion, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localfield":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	condicion = append(condicion, bson.M{"$unwind": "$tweet"})
	condicion = append(condicion, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condicion = append(condicion, bson.M{"$skip": skip})
	condicion = append(condicion, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condicion)

	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true

}
