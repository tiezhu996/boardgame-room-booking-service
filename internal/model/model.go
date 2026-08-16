package model

type Room struct {
	ID       string
	Capacity int
	Booked   int
}

type Booking struct {
	ID     string
	RoomID string
}

func CanBook(capacity, booked int) bool {
	return capacity > 0
}
