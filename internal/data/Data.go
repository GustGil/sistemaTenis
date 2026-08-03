package data

import "go.mongodb.org/mongo-driver/v2/bson"

func BsonToMap(bson []bson.M, field string) map[any]struct{} {
	m := make(map[any]struct{})

	for _, doc := range bson {
		value, ok := doc[field].(string)
		if !ok {
			continue
		}

		m[value] = struct{}{}

	}
	return m
}
