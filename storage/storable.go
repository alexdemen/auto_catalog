package storage

import "github.com/alexdemen/auto_catalog/model"

type Storable interface {
	GetCars() ([]*model.Car, error)
	AddCar(car *model.Car) error
	GetCar(id int) (*model.Car, error)
	DeleteCar(id int) error
	Update(car *model.Car) error
}
