package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {

			bookTickets(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			// print first names
			firstNames := getFirstNames(bookings)
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

func greetUsers(name string, totalTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v conference booking application!\n", name)
	fmt.Printf("We have a total of %v tickets, and %v are remaining\n\n", totalTickets, remainingTickets)
	fmt.Println("RSVP to save your seat!")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
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

func bookTickets(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
