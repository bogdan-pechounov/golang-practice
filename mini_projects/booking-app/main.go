package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 1)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {
	fmt.Printf("Welcome to %v!\n", conferenceName)
	fmt.Println("Remaining tickets:", remainingTickets)

	// fmt.Println("test", bookings, "test")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int
		// fmt.Println(firstName == "", &firstName)

		//get input
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		//update info
		if userTickets > int(remainingTickets) {
			fmt.Printf("We only have %v tickets remaining.\n", remainingTickets)
			continue
		}
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// a := -1
		// fmt.Println(uint(a))

		fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)

		if remainingTickets == 0 {
			break
		}
	}

	fmt.Printf("Bookings: %v\n", bookings)
	fmt.Printf("First names: %v\n", getFirstNames())

	wg.Wait()
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		numberOfTickets: userTickets,
		email:           email,
	}

	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, userData)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("#####")
	wg.Done()
}
