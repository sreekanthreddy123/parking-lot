package ParkingSpot

import "testing"

func TestParkingSpot(t *testing.T) {
	spot := ParkingSpot{1}
	got := spot.park()
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
