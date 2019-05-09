package ParkingSpot

import (
	"fmt"
)

type ParkingSpot struct {
	slotNumber int
}

func (e ParkingSpot) park() int {
	fmt.Printf("Park")
	e.slotNumber++
	fmt.Printf("%d", e.slotNumber)
	return e.slotNumber
}
