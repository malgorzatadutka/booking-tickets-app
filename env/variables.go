package env

import (
	"sync"
	"time"
)

// Booking app configuration
var (
	AvailableTickets uint = 5000
	ConcertTime           = time.Date(2022, time.November, 10, 21, 0, 0, 0, time.UTC)
	Wg                    = sync.WaitGroup{}
)
