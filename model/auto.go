package model

type Car struct {
	ID        int64  `jsonapi:"primary,cars"`
	BrandName string `jsonapi:"attr,brand"`
	Model     string `jsonapi:"attr,model"`
	Price     int    `jsonapi:"attr,price"`
	Status    int
	Mileage   int
}
