package transport

import "fmt"

type Transport interface {
	Move()
	Stop()
	ChangeSpeed(speed int)
}

type PassengerTransport interface {
	Transport
	BoardPassengers(num int)
	DisembarkPassengers(num int) error
}

type Car struct {
	speed      int
	passengers int
}

func (c *Car) Move() {
	fmt.Println("Автомобіль рухається")
}

func (c *Car) Stop() {
	fmt.Println("Автомобіль зупиняється")
}

func (c *Car) ChangeSpeed(speed int) {
	c.speed = speed
	fmt.Printf("Автомобіль змінює швидкість на %d км/год\n", speed)
}

func (c *Car) BoardPassengers(num int) {
	c.passengers += num
	fmt.Printf("В автомобіль сідає %d пасажирів\n", num)
}

func (c *Car) DisembarkPassengers(num int) error {
	var err error

	if num > c.passengers {
		err = fmt.Errorf("перевищена кількість пасажирів")
	}

	c.passengers -= num
	fmt.Printf("З автомобіля висаджується %d пасажирів\n", num)

	return err
}

type Train struct {
	passengers int
	speed      int
}

func (t *Train) Move() {
	fmt.Println("Потяг рухається")
}

func (t *Train) Stop() {
	fmt.Println("Потяг зупиняється")
}

func (t *Train) ChangeSpeed(speed int) {
	t.speed = speed
	fmt.Printf("Потяг змінює швидкість на %d км/год\n", speed)
}

func (t *Train) BoardPassengers(num int) {
	t.passengers += num
	fmt.Printf("В потяг сідає %d пасажирів\n", num)
}

func (t *Train) DisembarkPassengers(num int) error {
	var err error

	if num > t.passengers {
		err = fmt.Errorf("перевищена кількість пасажирів")
	}

	t.passengers -= num
	fmt.Printf("З потяга висаджується %d пасажирів\n", num)

	return err
}

type Airplane struct {
	passengers int
	speed      int
}

func (a *Airplane) Move() {
	fmt.Println("Літак рухається")
}

func (a *Airplane) Stop() {
	fmt.Println("Літак сідає")
}

func (a *Airplane) ChangeSpeed(speed int) {
	a.speed = speed
	fmt.Printf("Літак змінює швидкість на %d км/год\n", speed)
}

func (a *Airplane) BoardPassengers(num int) {
	a.passengers += num
	fmt.Printf("В літак сідає %d пасажирів\n", num)
}

func (a *Airplane) DisembarkPassengers(num int) error {
	var err error

	if num > a.passengers {
		err = fmt.Errorf("перевищена кількість пасажирів")
	}

	a.passengers -= num

	fmt.Printf("З літака висаджується %d пасажирів\n", num)

	return err
}
