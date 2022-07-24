package user

import (
	e "booking-app/env"
	"fmt"
	"time"
)

func BookTicket(userTickets uint, userName string, userLastName string, email string, phone string, sector uint) {
	e.AvailableTickets = e.AvailableTickets - userTickets

	//map for a user
	var userData = User{
		firstName:   userName,
		lastName:    userLastName,
		email:       email,
		phone:       phone,
		userTickets: userTickets,
		sector:      sector,
	}

	ticketsBuyers = append(ticketsBuyers, userData)
	fmt.Printf("List of ticket buyers: %v\n", ticketsBuyers)

	switch sector {
	case 1, 3, 5, 7, 9:
		fmt.Println("Right sector")
	case 2, 4, 6, 8, 10:
		fmt.Println("Left sector")
	case 11, 12, 13:
		fmt.Println("Vip sector")

	}

	fmt.Printf("Thank you %s %s for booking %d tickets.\nYou will recieve confirmation email at adress: %s and SMS at phone: %s\n", userName, userLastName, userTickets, email, phone)
	fmt.Printf("For %s concert %v tickets remaining\n", e.BandName, e.AvailableTickets)
}

func SendTicket(userTickets uint, sector uint, userName string, userLastName string, email string) {
	time.Sleep(20 * time.Second)
	//ticket generation
	var ticket = fmt.Sprintf("%v tickets in sector %v for user %v %v", userTickets, sector, userName, userLastName)
	//sending ticket
	fmt.Println("************")
	fmt.Printf("Sending ticket %v to email: %v", ticket, email)
	fmt.Println("************")
	e.Wg.Done()
}
