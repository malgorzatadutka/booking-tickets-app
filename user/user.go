package user

import "fmt"

var (
	ticketsBuyers = make([]User, 1)
)

type User struct {
	firstName   string
	lastName    string
	email       string
	phone       string
	userTickets uint
	sector      uint
}

func UsersFirstNames() []string {
	firstNames := []string{}
	for _, ticketsBuyer := range ticketsBuyers {
		firstNames = append(firstNames, ticketsBuyer.firstName)
	}
	return firstNames
}

func GetUserInput() (string, string, string, string, uint, uint) {
	var userName string
	var userLastName string
	var email string
	var phone string
	var userTickets uint
	var sector uint

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

	fmt.Println("Enter preferable tickets' sector: ")
	fmt.Scan(&sector)

	return userName, userLastName, email, phone, userTickets, sector
}
