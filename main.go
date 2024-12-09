package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"

)

var conferenceName = "Go Conference"

const conferenceTicket = 50

var remainingTicket uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for len(bookings) < 50 {

		firstName, secondName, email, numberOfTickets := getUserInput()
		isValidName, isEmailValid, isTicketsValid := helper.ValidateUserInput(firstName, secondName, email, numberOfTickets, remainingTicket)

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

func getBookingUserFirstName(bookings []map[string]string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
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

func bookTicket(numberOfTickets uint, firstName string, secondName string, email string) ([]map[string]string, uint) {
	remainingTicket = remainingTicket - numberOfTickets

	newBooking := make(map[string]string)

	newBooking["firstName"] = firstName
	newBooking["lastName"] = secondName
	newBooking["email"] = email
	newBooking["numberOfTickets"] = strconv.FormatUint(uint64(numberOfTickets), 10)

	bookings = append(bookings, newBooking)

	fmt.Printf("List of the bookings is %v \n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, secondName, numberOfTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTicket, conferenceName)
	return bookings, remainingTicket
}
