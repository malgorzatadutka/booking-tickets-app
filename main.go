package main

import (
	"booking-app/env"
	"booking-app/functions"
	"booking-app/user"
	"fmt"
)

func main() {

	functions.WelcomeUsers()
	functions.DateAndPlace()

	for {
		userName, userLastName, email, phone, userTickets, sector := user.GetUserInput()

		isValidName, isValidEmail, isValidTicket, isValidSector := functions.ValidateUserInput(userName, userLastName, email, userTickets, sector, env.AvailableTickets)

		if isValidEmail && isValidName && isValidTicket && isValidSector {

			user.BookTicket(userTickets, userName, userLastName, email, phone, sector)
			fmt.Printf("The first names of people booking tickets: %v\n", user.UsersFirstNames())
			// starts new goroutine
			env.Wg.Add(1)
			go user.SendTicket(userTickets, sector, userName, userLastName, email)

			if env.AvailableTickets == 0 {
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
	env.Wg.Wait()
}
