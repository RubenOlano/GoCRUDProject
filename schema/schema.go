package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id     primitive.ObjectID `json:"_id,omitempty"`
	Title  string             `json:"title,omitempty"`
	Year   int32              `json:"year,omitempty"`
	Poster string             `json:"poster,omitempty"`
}
