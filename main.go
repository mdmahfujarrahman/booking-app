package main

import (
	"fmt"
	"strings"

)

var conferenceName = "Go Conference"

const conferenceTicket = 50

var remainingTicket uint = 50
var bookings []string

func main() {

	greetUsers()

	for len(bookings) < 50 {

		firstName, secondName, email, numberOfTickets := getUserInput()
		isValidName, isEmailValid, isTicketsValid := validateUserInput(firstName, secondName, email, numberOfTickets)

		if !isValidName || !isEmailValid || !isTicketsValid {
			errorMessage := getInputValidationError(isValidName, isEmailValid, isTicketsValid)
			fmt.Println(errorMessage)
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
			bookings, remainingTicket := bookTicket(numberOfTickets, firstName, secondName, email)

			firstNames := getBookingUserFirstName(bookings)

			fmt.Printf("The first names of booking are %v \n", firstNames)

			notTicketsRemaining := remainingTicket == 0
			if notTicketsRemaining {
				fmt.Printf("Our %v tickets is booked out. Come back next year", conferenceName)
				break
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking Application \n", conferenceName)
	fmt.Printf("Now We have %v tickets \n", remainingTicket)
	fmt.Printf("Get Your ticket Now %v \n", conferenceTicket)
}

func getBookingUserFirstName(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		name := strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}

	return firstNames

}

func validateUserInput(firstName string, secondName string, email string, numberOfTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(secondName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isTicketsValid := numberOfTickets > 0 && numberOfTickets <= remainingTicket
	return isValidName, isEmailValid, isTicketsValid
}

func getInputValidationError(isValidName bool, isEmailValid bool, isTicketsValid bool) string {
	var errorMessage string

	if !isValidName {
		errorMessage = "Please Enter Valid Name"
	}

	if !isEmailValid {
		errorMessage = "Please Enter valid Email Address"
	}

	if !isTicketsValid {
		errorMessage = "Please Enter valid tickets number"
	}

	return errorMessage
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, secondName, email, numberOfTickets

}

func bookTicket(numberOfTickets uint, firstName string, secondName string, email string) ([]string, uint) {
	remainingTicket = remainingTicket - numberOfTickets
	bookings = append(bookings, firstName+" "+secondName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, secondName, numberOfTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTicket, conferenceName)
	return bookings, remainingTicket
}
