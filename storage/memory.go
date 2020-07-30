package storage

import "github.com/alexdemen/auto_catalog/model"

type MemoryStorage struct{}

func (m MemoryStorage) GetCars() []*model.Car {
	return []*model.Car{
		{ID: 1, BrandName: "BMW", Model: "X5", Price: 450000},
		{ID: 1, BrandName: "BMW", Model: "X5", Price: 451000},
		{ID: 1, BrandName: "BMW", Model: "X5", Price: 452000},
		{ID: 1, BrandName: "BMW", Model: "X5", Price: 453000},
		{ID: 1, BrandName: "Volvo", Model: "model 1", Price: 454000},
	}
}
