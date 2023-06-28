package route

import (
	"fmt"
	"go_course/hw6/task2/transport"
)

type Route struct {
	vehicles []transport.PassengerTransport
}

func (r *Route) AddVehicle(t transport.PassengerTransport) {
	r.vehicles = append(r.vehicles, t)
}

func (r *Route) ShowVehicles() {
	fmt.Println("Транспортні засоби на маршруті:")
	for _, vehicle := range r.vehicles {
		fmt.Printf("- %T\n", vehicle)
	}
}

func (r *Route) Travel() {
	fmt.Println("Подорож розпочалась!")

	for _, vehicle := range r.vehicles {
		vehicle.Move()
		vehicle.ChangeSpeed(100)
		vehicle.BoardPassengers(10)
		vehicle.Stop()
		err := vehicle.DisembarkPassengers(5)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			break
		}
	}

	fmt.Println("Подорож завершилась!")
}
