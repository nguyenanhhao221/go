package elon

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", car.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	if c.speed == c.batteryDrain {
		return trackDistance <= c.battery
	}
	return (trackDistance / c.speed * c.batteryDrain) <= c.battery
}
