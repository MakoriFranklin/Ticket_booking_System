package main

import (
	"fmt"
	"strings"
)

func main() {
	ConferenceName := "Go Confernce"
	const ConferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}
	fmt.Printf("Welcome to %v booking application \n", ConferenceName)
	fmt.Printf("We have  total of %v tickets and %v still available! \n", ConferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// var userTickets int
		fmt.Println("Please enter your username:")
		fmt.Scan(&firstName)
		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email:")
		fmt.Scan(&email)
		fmt.Println("Please enter number of tickets:")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining. You need to book %v tickets or less", remainingTickets, remainingTickets)
			break
		}

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v remaining tickets for %v \n", remainingTickets, ConferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Println("These are all our bookings %v", firstNames)
		if remainingTickets == 0 {
			fmt.Println("Our conference is sold out. Try again nex year")
			break
		}
	}

}
