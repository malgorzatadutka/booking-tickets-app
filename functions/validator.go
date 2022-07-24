package functions

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, tickets uint, sector uint, availableTickets uint) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := tickets > 0 && tickets <= availableTickets
	isValidSector := sector >= 1 && sector <= 13
	//return multiple values
	return isValidName, isValidEmail, isValidTicket, isValidSector
}
