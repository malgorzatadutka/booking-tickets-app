package functions

import (
	"strconv"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, tickets uint, sector string, availableTickets uint) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := tickets > 0 && tickets <= availableTickets
	sect, err := strconv.Atoi(sector)
	if err != nil {
		// ... handle error
		panic(err)
	}
	isValidSector := sect >= 1 && sect <= 13
	//return multiple values
	return isValidName, isValidEmail, isValidTicket, isValidSector
}
