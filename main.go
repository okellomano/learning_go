package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidUserTickets {

			bookTickets(userTickets, firstName, lastName, email)

			// print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// no more tickets.
				fmt.Println("The conference is booked out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name of last name enterd is too short")
			}
			if !isValidEmail {
				fmt.Println("Enter a proper email address")
			}
			if !isValidUserTickets {
				fmt.Println("The number of tickets is invalid!")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v conference booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are remaining\n\n", conferenceTickets, remainingTickets)
	fmt.Println("RSVP to save your seat!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTickets
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// get the user input
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tckets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
