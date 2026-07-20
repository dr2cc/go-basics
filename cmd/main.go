package main

import (
	"errors"
	"fmt"
	"sync"
)

type employee struct {
	id     int
	name   string
	age    int
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

type memoryStorage struct {
	mu sync.RWMutex
	db map[int]employee
}

func (d *memoryStorage) insert(e employee) error {
	d.db[e.id] = e
	return nil
}

func (d *memoryStorage) get(id int) (employee, error) {
	e, ok := d.db[id]
	if !ok {
		return employee{}, errors.New("employee with such id doesen`t exist")
	}
	return e, nil
}

func (d *memoryStorage) delete(id int) error {
	delete(d.db, id)
	return nil
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		db: make(map[int]employee),
	}
}

type DumbStorage struct{}

func newDumbStorage() *DumbStorage {
	return &DumbStorage{}
}

func (d *DumbStorage) insert(e employee) error {
	fmt.Printf("вставка пользователя с id: %d прошла успешно\n", e.id)
	return nil
}

func (d *DumbStorage) get(id int) (employee, error) {
	e := employee{
		id: id,
	}
	return e, nil
}

func (d *DumbStorage) delete(id int) error {
	fmt.Printf("удаление пользователя с id: %d прошла успешно\n", id)
	return nil
}

func main() {
	ms := newMemoryStorage()
	ds := newDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.get(3))

	spawnEmployees(ds)
}

func spawnEmployees(s storage) {
	for i := 1; i <= 10; i++ {
		s.insert(employee{id: i})
	}
}
