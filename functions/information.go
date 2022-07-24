package functions

import (
	e "booking-app/env"
	"fmt"
	"time"
)

func WelcomeUsers() {
	fmt.Printf("Welcom in %s concert's ticket reservation system\nGrab Your ticket!\n", e.BandName)
	fmt.Printf("We have total of %d tickets, but only %d tickets are still avaliable for sale!\n", e.ConcertTickets, e.AvailableTickets)
}

func DateAndPlace() {
	fmt.Printf("Date: %v,\nLocation: %s\n", e.ConcertTime.Format(time.ANSIC), e.ConcertLocation)
}
