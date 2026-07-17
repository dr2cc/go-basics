package main

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

func main() {

}
