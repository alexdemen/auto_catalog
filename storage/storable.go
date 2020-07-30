package storage

import "github.com/alexdemen/auto_catalog/model"

type Storable interface {
	GetCars() []*model.Car
	AddCar(car model.Car) error
}
