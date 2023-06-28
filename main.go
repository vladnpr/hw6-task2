package main

import (
	"go_course/hw6/task2/route"
	"go_course/hw6/task2/transport"
)

func main() {
	car := transport.Car{}
	train := transport.Train{}
	airplane := transport.Airplane{}

	r := route.Route{}

	r.AddVehicle(&car)
	r.AddVehicle(&train)
	r.AddVehicle(&airplane)
	r.ShowVehicles()
	r.Travel()
}
