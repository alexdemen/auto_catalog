package storage

import "github.com/alexdemen/auto_catalog/model"

type MemoryStorage struct {
	data map[int64]*model.Car
}

func NewMemoryStorage() *MemoryStorage {
	store := MemoryStorage{data: make(map[int64]*model.Car)}
	store.data[1] = &model.Car{ID: 1, BrandName: "BMW", Model: "X5", Price: 450000, Status: model.IN_STOCK_STATUS}
	store.data[2] = &model.Car{ID: 2, BrandName: "BMW", Model: "X5", Price: 450000, Status: model.OUT_OF_SALES_STATUS}
	store.data[3] = &model.Car{ID: 3, BrandName: "BMW", Model: "X5", Price: 450000, Status: model.SOLD_STATUS}
	store.data[4] = &model.Car{ID: 4, BrandName: "BMW", Model: "X5", Price: 450000, Status: model.TRANSIT_STATUS}
	store.data[5] = &model.Car{ID: 5, BrandName: "BMW", Model: "X5", Price: 450000, Status: model.SOLD_STATUS}

	return &store
}

func (m MemoryStorage) GetCars() []*model.Car {
	result := make([]*model.Car, 0, len(m.data))
	for _, val := range m.data {
		result = append(result, val)
	}

	return result
}

func (m MemoryStorage) AddCar(car model.Car) error {
	panic("implement me")
}
