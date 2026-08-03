package product

import (
	"sistemaTenis/internal/data"
	"sistemaTenis/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func SaveBasic(shoe *Tenis) {
	repository.InsertAny("SneakerDB", "shoe", shoe)
}

func GetShoesByParam(param *bson.M) map[any]struct{} {

	resp := repository.ProjectionAny("SneakerDB", "shoe", param)
	ShoeUrlList := data.BsonToMap(resp, "url")

	return ShoeUrlList
}
