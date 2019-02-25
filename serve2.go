package main

import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_petal      uint16
	brake_petal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func (c car) kmh() float64 {
	return float64(c.gas_petal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car{23000, 0, 15000, 190.0}
	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println("<head></head>")
	fmt.Println("<body>")
	fmt.Println(a_car.gas_petal)
	fmt.Println(a_car.kmh())
	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println("</body>")
	fmt.Println("</html>")
}
