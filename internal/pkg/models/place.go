package models

type Place struct {
	ID          string
	Address     string
	Coordinates Coordinates
}

type Coordinates struct {
	Latitude  float32
	Longitude float32
}
