package main

import (
	"fmt"
	"strings"
	"strconv"
)

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint= 50
	var bookings []string

func main() {
		greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			processOrder(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address entered does not contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid.")
			}
		}
	}
}

func greetUsers()  {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//Ask user for their info
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter how many tickets you'd like to purchase: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func processOrder(userTickets uint, firstName string, lastName string, email string) (uint, uint, []string, string, string, string, string) {
	//Subtract number of tickets purchased from the remaining availble tickets
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you, %v %v, for booking %v tickets.\nYou will receive a confirmation email at %v\n",firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	return remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName
}