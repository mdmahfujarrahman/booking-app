package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTicket uint = 50
	var bookings []string

	greetUsers(conferenceName, conferenceTicket, remainingTicket)

	for len(bookings) < 50 {
		var firstName string
		var secondName string
		var email string
		var numberOfTickets uint

		fmt.Println("Please Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Please Enter your last name: ")
		fmt.Scan(&secondName)
		fmt.Println("Please Enter your Email: ")
		fmt.Scan(&email)
		fmt.Printf("Number of tickets you want but,left only %v \n", remainingTicket)
		fmt.Scan(&numberOfTickets)

		isValidName := len(firstName) >= 2 && len(secondName) >= 2
		isEmailValid := strings.Contains(email, "@")
		isTicketsValid := numberOfTickets > 0 && numberOfTickets <= remainingTicket

		if !isValidName || !isEmailValid || !isTicketsValid {
			fmt.Println("Your input data was incorrect")
		} else {

			if numberOfTickets > 45 {
				fmt.Println("Sorry! You can't buy more then 45 tickets")
				fmt.Scan(&numberOfTickets)
			}

			if numberOfTickets > remainingTicket {
				fmt.Printf("You can't buy %v number of tickets, because we have left only %v \n", numberOfTickets, remainingTicket)
				fmt.Printf("Number of tickets you want but,left only %v \n", remainingTicket)
				continue
			}

			remainingTicket = remainingTicket - numberOfTickets
			bookings = append(bookings, firstName+" "+secondName)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, secondName, numberOfTickets, email)
			fmt.Printf("%v tickets remaining for %v \n", remainingTicket, conferenceName)

			getBookingUserFirstName(bookings)

			notTicketsRemaining := remainingTicket == 0

			if notTicketsRemaining {
				fmt.Printf("Our %v tickets is booked out. Come back next year", conferenceName)
				break
			}
		}

	}

}

func greetUsers(conferenceName string, conferenceTicket int, remainingTicket uint) {
	fmt.Printf("Welcome to %v booking Application \n", conferenceName)
	fmt.Printf("Now We have %v tickets \n", remainingTicket)
	fmt.Printf("Get Your ticket Now %v \n", conferenceTicket)
}

func getBookingUserFirstName(bookings []string) {

	firstNames := []string{}

	for _, booking := range bookings {
		name := strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}

	fmt.Printf("The first names of booking are %v \n", firstNames)

}

func getUserInput() {

}
