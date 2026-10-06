package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Tenis struct {
	ID       bson.ObjectID `bson:"_id, omitempty"`
	Name     string        `bson:"name"`
	Price    string        `bson:"price"`
	Category string        `bson:"category"`
	Url      string        `bson:"url"`
	Status   string        `bson:"status"`
	Prompt   string        `bson:"prompt"`
}

type CrushionTech struct {
	ID       int
	Name     string
	Category []string
}
