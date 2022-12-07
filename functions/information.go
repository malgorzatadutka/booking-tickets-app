package functions

import (
	"booking-app/env"
	"fmt"
	"time"
)

// WelcomeUsers greets user and inform about concert details
func WelcomeUsers() {
	fmt.Printf("Welcom in %s concert's ticket reservation system\nGrab Your ticket!\n", env.BandName)
	fmt.Printf("We have total of %d tickets, but only %d tickets are still avaliable for sale!\n", env.ConcertTickets, env.AvailableTickets)
}

// DateAndPlace inform user about concert's date and location
func DateAndPlace() {
	fmt.Printf("Date: %v,\nLocation: %s\n", env.ConcertTime.Format(time.ANSIC), env.ConcertLocation)
}
