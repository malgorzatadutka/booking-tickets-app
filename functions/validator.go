package functions

import (
	"strings"
)

//ValidateUserInput validate format of passing user's information
func ValidateUserInput(firstName string, lastName string, email string, tickets uint, sector uint, availableTickets uint) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidSector := sector >= 1 && sector <= 13
	isValidTicket := tickets > 0 && tickets <= availableTickets

	//return multiple values
	return isValidName, isValidEmail, isValidTicket, isValidSector
}
