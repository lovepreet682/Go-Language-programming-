package main

import "fmt"

//create the interface
type vehicle interface {
	Start()
	Stop()
}

// Define a Car type that implements the Vehicle interface
type Car struct {
	Name string
}

//crate the method
func (c Car) Start() {
	fmt.Println(c.Name, "is Starting...")
}
func (c Car) Stop() {
	fmt.Println(c.Name, "is Stop")
}

func main() {
	// Create instances of Car
	var car = Car{Name: "Swift"}
	car.Start()
	car.Stop()
}
