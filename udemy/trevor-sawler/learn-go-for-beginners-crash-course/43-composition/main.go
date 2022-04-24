package main

import (
	"fmt"
)

type Vehicle struct {
	NrOfWheels     int
	NrOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of passengers: ", v.NrOfPassengers)
	fmt.Println("Number of wheels:", v.NrOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("Is Electric:", c.IsElectric)
	fmt.Println("Is Hybric:", c.IsHybrid)
	c.Vehicle.showDetails()
}

func main() {
	suv := Vehicle{
		NrOfWheels:     4,
		NrOfPassengers: 6,
	}

	volvoXC90 := Car{
		Make:       "Volvo",
		Model:      "XC90 T8",
		Year:       2021,
		IsElectric: false,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	volvoXC90.show()

	teslaModelX := Car{
		Make:       "Tesla",
		Model:      "Model X",
		Year:       2021,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}
	teslaModelX.show()
}
