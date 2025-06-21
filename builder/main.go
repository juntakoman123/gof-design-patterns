package main

import "fmt"

type Car struct {
	Engine string
	Seats  int
	GPS    bool
	Color  string
}

// CarBuilder は Car を段階的に構築する
type CarBuilder struct {
	engine string
	seats  int
	gps    bool
	color  string
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (b *CarBuilder) SetEngine(e string) *CarBuilder {
	b.engine = e
	return b
}

func (b *CarBuilder) SetSeats(n int) *CarBuilder {
	b.seats = n
	return b
}

func (b *CarBuilder) SetGPS(enabled bool) *CarBuilder {
	b.gps = enabled
	return b
}

func (b *CarBuilder) SetColor(c string) *CarBuilder {
	b.color = c
	return b
}

func (b *CarBuilder) Build() *Car {
	return &Car{
		Engine: b.engine,
		Seats:  b.seats,
		GPS:    b.gps,
		Color:  b.color,
	}
}

func main() {
	car := NewCarBuilder().
		SetEngine("V8").
		SetSeats(4).
		SetGPS(true).
		SetColor("Red").
		Build()

	fmt.Println(car.Engine)
}
