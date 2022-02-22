package entity

type Store struct {
	ID       string   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string   `bson:"name" json:"name"`
	Gps      Gps      `bson:"gps" json:"gps"`
	Location Location `bson:"location" json:"location"`
	Open     string   `bson:"open" json:"open"`
	Products []string `bson:"products" json:"products"`
	Source   string   `bson:"source" json:"source"`
}
type Gps struct {
	Latitude  float64 `bson:"latitude" json:"latitude"`
	Longitude float64 `bson:"longitude" json:"longitude"`
}
type Location struct {
	Street  string `bson:"street" json:"street"`
	Zip     string `bson:"zip" json:"zip"`
	City    string `bson:"city" json:"city"`
	Country string `bson:"country" json:"country"`
}
