package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Resume struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FileName   string             `bson:"fileName" json:"fileName"`
	RawText    string             `bson:"rawText" json:"rawText"`
	ParsedName string             `bson:"parsedName" json:"parsedName"`
	Email      string             `bson:"email" json:"email"`
	Phone      string             `bson:"phone" json:"phone"`
	Skills     []string           `bson:"skills" json:"skills"`
}
