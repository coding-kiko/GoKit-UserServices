package utils

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// converts struct to bson document
func StructtoBson(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}

func GetCorrectFilterId(id string) bson.D {
	if strings.Contains(id, "@") {
		return bson.D{{"email", id}}
	} else {
		return bson.D{{"_id", id}}
	}
}
