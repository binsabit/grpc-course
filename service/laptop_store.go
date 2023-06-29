package service

import (
	"errors"
	"sync"

	"github.com/binsabit/grpc-course/protogen/pb"
)

var ErrorAlreadyExist = errors.New("laptop already exist")
var ErrorDoesNotExist = errors.New("laptop does not exist")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{data: make(map[string]*pb.Laptop)}
}

func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, ErrorDoesNotExist
	}
	copyLaptop := &pb.Laptop{
		Id:          laptop.Id,
		Brand:       laptop.Brand,
		Name:        laptop.Name,
		Cpu:         laptop.Cpu,
		Gpus:        laptop.Gpus,
		Ram:         laptop.Ram,
		Keyboard:    laptop.Keyboard,
		Screen:      laptop.Screen,
		Storages:    laptop.Storages,
		Weight:      laptop.Weight,
		PriceUsd:    laptop.PriceUsd,
		ReleaseYear: laptop.ReleaseYear,
		UpdatedAt:   laptop.UpdatedAt,
	}
	return copyLaptop, nil
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if _, ok := store.data[laptop.Id]; ok {
		return ErrorAlreadyExist
	}

	copyLaptop := &pb.Laptop{
		Id:          laptop.Id,
		Brand:       laptop.Brand,
		Name:        laptop.Name,
		Cpu:         laptop.Cpu,
		Gpus:        laptop.Gpus,
		Ram:         laptop.Ram,
		Keyboard:    laptop.Keyboard,
		Screen:      laptop.Screen,
		Storages:    laptop.Storages,
		Weight:      laptop.Weight,
		PriceUsd:    laptop.PriceUsd,
		ReleaseYear: laptop.ReleaseYear,
		UpdatedAt:   laptop.UpdatedAt,
	}
	store.data[laptop.Id] = copyLaptop
	return nil

}
