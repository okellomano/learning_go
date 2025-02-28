package main

import (
	"booking-app/shared"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go"
var remainingTickets uint = 50

// initializing a list of maps (for same data types)
var bookings = make([]UserData, 0)

// use struct to save different data types
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidUserTickets := shared.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidUserTickets {
		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// print first names
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// no more tickets.
			fmt.Println("The conference is booked out!")
			//break
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v conference booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are remaining\n\n", conferenceTickets, remainingTickets)
	fmt.Println("RSVP to save your seat!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// Once a user books a ticket, send it to them via email
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#########################################################")
	fmt.Printf("Senging ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#########################################################")
	wg.Done()
}
