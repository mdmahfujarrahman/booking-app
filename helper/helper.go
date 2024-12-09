package helper

import "strings"

func ValidateUserInput(firstName string, secondName string, email string, numberOfTickets uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(secondName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isTicketsValid := numberOfTickets > 0 && numberOfTickets <= remainingTicket
	return isValidName, isEmailValid, isTicketsValid
}
