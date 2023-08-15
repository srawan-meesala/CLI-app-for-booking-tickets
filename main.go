package main

import (
	"fmt"
	"strings"
)

func main(){
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application.\n",conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here.")

	var bookings []string

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter you first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter you last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter you Email Address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only %v tickets remaining, so you can't book %v tickets\n",remainingTickets,userTickets)
		} else {
			bookings = append(bookings, firstName + " " + lastName)
			remainingTickets -= userTickets

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
			fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)


			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n\n",firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		}
	}
}