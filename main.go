package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are remaining\n", conferenceTickets, remainingTickets)

	fmt.Println("RSVP to save your seat!")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	bookings := []string{}

	// get the user input
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tckets:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, userTickets, email)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("%v remaining tickets for %v \n", remainingTickets, conferenceName)

	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Bookings: %v\n", bookings)
}
