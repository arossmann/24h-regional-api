package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Store struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string             `json:"name"`
	Gps      Gps                `json:"gps"`
	Location Location           `json:"location"`
	Open     string             `json:"open"`
	Products []string           `json:"products"`
	Source   string             `json:"source"`
}
type Gps struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Location struct {
	Street  string `json:"street"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Country string `json:"country"`
}
