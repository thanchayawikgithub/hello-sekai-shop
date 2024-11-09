package utils

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ConvertToObject(id string) (bson.ObjectID, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return bson.NilObjectID, err
	}
	return objectID, nil
}
