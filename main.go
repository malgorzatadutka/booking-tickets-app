package main

import (
	e "booking-app/env"
	f "booking-app/functions"
	u "booking-app/user"
	"fmt"
)

func main() {

	f.WelcomeUsers()
	f.DateAndPlace()

	for {
		userName, userLastName, email, phone, userTickets, sector := u.GetUserInput()

		isValidName, isValidEmail, isValidTicket, isValidSector := f.ValidateUserInput(userName, userLastName, email, userTickets, sector, e.AvailableTickets)

		if isValidEmail && isValidName && isValidTicket && isValidSector {

			u.BookTicket(userTickets, userName, userLastName, email, phone, sector)
			fmt.Printf("The first names of people booking tickets: %v\n", u.UsersFirstNames())
			// starts new goroutine
			e.Wg.Add(1)
			go u.SendTicket(userTickets, sector, userName, userLastName, email)

			if e.AvailableTickets == 0 {
				//program ending
				fmt.Println("Concert is soled out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Input error: First or last name too short")
			} else if !isValidEmail {
				fmt.Println("Input error: Email doesn't contain @ sign")
			} else if !isValidTicket {
				fmt.Println("Input error: Invalid number of tickets")
			} else if isValidSector {
				fmt.Println("Input error: Invalid sector number")
			} else {
				fmt.Println("Please try again!")
			}
		}

	}
	e.Wg.Wait()
}
