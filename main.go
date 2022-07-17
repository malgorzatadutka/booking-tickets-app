package main

import (
	f "booking-app/functions"
	"fmt"
	"strconv"
	"time"
)

//package level variables
var bandName string = "Queen"
var concertLocation string = "London"
var concertTime = time.Date(2022, time.November, 10, 21, 0, 0, 0, time.UTC)

const concertTickets uint = 5000

var availableTickets uint = 5000
var ticketsBuyers = make([]map[string]string, 1)

func main() {

	welcomeUsers()
	dateAndPlace()

	for {
		userName, userLastName, email, phone, userTickets, sector := getUserInput()

		isValidName, isValidEmail, isValidTicket, isValidSector := f.ValidateUserInput(userName, userLastName, email, userTickets, sector, availableTickets)

		if isValidEmail && isValidName && isValidTicket && isValidSector {

			bookTicket(userTickets, userName, userLastName, email, phone, sector)
			fmt.Printf("These are all of our selling tickets: %v\n", usersFirstNames())

			if availableTickets == 0 {
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

}

func welcomeUsers() {
	fmt.Printf("Welcom in %s concert's ticket reservation system\nGrab Your ticket!\n", bandName)
	fmt.Printf("We have total of %d tickets, but only %d tickets are still avaliable for sale!\n", concertTickets, availableTickets)
}

func dateAndPlace() {
	fmt.Printf("Date: %v,\nLocation: %s\n", concertTime.Format(time.ANSIC), concertLocation)
}

func usersFirstNames() []string {
	firstNames := []string{}
	for _, ticketsBuyer := range ticketsBuyers {
		firstNames = append(firstNames, ticketsBuyer["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, string, uint, string) {
	var userName string
	var userLastName string
	var email string
	var phone string
	var userTickets uint
	var sector string

	fmt.Println("Enter your first name: ")
	// pointer to userName
	fmt.Scan(&userName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&userLastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your phone: ")
	fmt.Scan(&phone)

	fmt.Println("Enter tickets' number: ")
	fmt.Scan(&userTickets)

	fmt.Println("Enter tickets' sector: ")
	fmt.Scan(&sector)

	return userName, userLastName, email, phone, userTickets, sector
}

func bookTicket(userTickets uint, userName string, userLastName string, email string, phone string, sector string) {
	availableTickets = availableTickets - userTickets

	//map for a user
	var userData = make(map[string]string)
	userData["firstName"] = userName
	userData["lastName"] = userLastName
	userData["email"] = email
	userData["phone"] = phone
	userData["sector"] = sector
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	ticketsBuyers = append(ticketsBuyers, userData)
	fmt.Printf("List of ticket buyers: %v\n", ticketsBuyers)

	switch sector {
	case "1", "3", "5", "7", "9":
		fmt.Println("Right sector")
	case "2", "4", "6", "8", "10":
		fmt.Println("Left sector")
	case "11", "12", "13":
		fmt.Println("Vip sector")

	}

	fmt.Printf("Thank you %s %s for booking %d tickets.\nYou will recieve confirmation email at adress: %s and SMS at phone: %s\n", userName, userLastName, userTickets, email, phone)
	fmt.Printf("For %s concert %v tickets remaining\n", bandName, availableTickets)
}
