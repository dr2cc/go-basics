package main

import (
	"fmt"
	"go_basics/storage"
)

func main() {
	ms := storage.NewMemoryStorage()
	ds := storage.NewDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.Get(3))

	spawnEmployees(ds)
}

func spawnEmployees(s storage.Storage) {
	for i := 1; i <= 10; i++ {
		s.Insert(storage.Employee{ID: i})
	}
}
