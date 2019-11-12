package main

import (
	"fmt")

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934


type car struct {
	gasPedal uint16
	brakePedal uint16
	steeringWheel uint16
	topSpeedKmh float64

}

func newerTopSpeed(c car, speed float64) car {
	c.topSpeedKmh = speed
	return c
}

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)

}

func (c *car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax / kmh_multiple)

}

func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}


func main() {
	a_car := car {
		gasPedal: 65000, 
		brakePedal: 0,
		steeringWheel: 12561,
		topSpeedKmh: 225.0}

	fmt.Println(a_car.gasPedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car = newerTopSpeed(a_car, 500)

	//a_car.newTopSpeed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

}
