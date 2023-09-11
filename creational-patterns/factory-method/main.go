package main

import (
	"fmt"
)

// IProduct
type ICar interface {
	getName() string
	setName(name string)
	drive()
}

// BaseProduct
type Car struct {
	name string
}

func (c *Car) getName() string {
	return c.name
}

func (c *Car) setName(name string) {
	c.name = name
}

func (s *Car) drive() {
	fmt.Println("Driving Car at 20 kph")
}

// ConcreteProduct
type Sedan struct {
	Car
}

func (s *Sedan) drive() {
	fmt.Println("Driving Sedan at 60 kph")
}

// ConcreteProduct
type SUV struct {
	Car
}

func (s *SUV) drive() {
	fmt.Println("Driving SUV at 30 kph")
}

// ProductFactory
type CarFactory struct{}

func (cf *CarFactory) createCar(name string) ICar {
	switch name {
	case "sedan":
		return &Sedan{
			Car{
				name: "Brand A Sedan",
			},
		}
	case "suv":
		return &SUV{
			Car{
				name: "Brand A SUV",
			},
		}
	default:
		return nil
	}
}

// Client
func main() {

	cf := &CarFactory{}

	sedan := cf.createCar("sedan")
	printDetails(sedan)
	sedan.drive()

	suv := cf.createCar("suv")
	printDetails(suv)
	suv.drive()
}

func printDetails(c ICar) {
	fmt.Printf("Brand: %s", c.getName())
	fmt.Println()
}
