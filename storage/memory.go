package storage

import (
	"errors"
	"fmt"
	"github.com/alexdemen/auto_catalog/model"
	"sync"
)

type MemoryStorage struct {
	data      map[int]*model.Car
	idCounter int
	mutex     sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	store := MemoryStorage{data: make(map[int]*model.Car)}
	return &store
}

func (m *MemoryStorage) GetCars() ([]*model.Car, error) {
	result := make([]*model.Car, 0, len(m.data))
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, val := range m.data {
		result = append(result, val)
	}

	return result, nil
}

func (m *MemoryStorage) AddCar(car *model.Car) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.idCounter++
	car.ID = m.idCounter
	m.data[car.ID] = car

	return nil
}

func (m *MemoryStorage) GetCar(id int) (*model.Car, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if car, exist := m.data[id]; exist {
		return car, nil
	}

	return nil, errors.New(fmt.Sprintf("car with id '%d' not found", id))
}

func (m *MemoryStorage) DeleteCar(id int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.data, id)

	return nil
}

func (m *MemoryStorage) Update(car *model.Car) error {
	if _, exist := m.data[car.ID]; !exist {
		return errors.New("car with id '%d' not exist")
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data[car.ID] = car

	return nil
}
