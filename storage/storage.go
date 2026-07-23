package storage

import (
	"errors"
	"fmt"
	"sync"
)

type Employee struct {
	ID     int
	name   string
	age    int
	salary int
}

type Storage interface {
	Insert(e Employee) error
	Get(id int) (Employee, error)
	Delete(id int) error
}

type memoryStorage struct {
	mu sync.RWMutex
	db map[int]Employee
}

func (d *memoryStorage) Insert(e Employee) error {
	d.db[e.ID] = e
	return nil
}

func (d *memoryStorage) Get(id int) (Employee, error) {
	e, ok := d.db[id]
	if !ok {
		return Employee{}, errors.New("employee with such id doesen`t exist")
	}
	return e, nil
}

func (d *memoryStorage) Delete(id int) error {
	delete(d.db, id)
	return nil
}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{
		db: make(map[int]Employee),
	}
}

type DumbStorage struct{}

func NewDumbStorage() *DumbStorage {
	return &DumbStorage{}
}

func (d *DumbStorage) Insert(e Employee) error {
	fmt.Printf("вставка пользователя с id: %d прошла успешно\n", e.ID)
	return nil
}

func (d *DumbStorage) Get(id int) (Employee, error) {
	e := Employee{
		ID: id,
	}
	return e, nil
}

func (d *DumbStorage) Delete(id int) error {
	fmt.Printf("удаление пользователя с id: %d прошла успешно\n", id)
	return nil
}
