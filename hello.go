package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := [] string{} // slices allow custom size of the list

	fmt.Printf("conferenceTickets is %d, remainingTickets is %d, conferenceName is %s\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//var bookings = [50]string{"Nana", "Nicole", "Peter"} -> with default items

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for his name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if remainingTickets == 0 {
			fmt.Println("No enough tickets")
			return
		}

		if userTickets > remainingTickets {
			fmt.Println("No enough tickets to buy")
			return
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thanks you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		
		firstNames := []string{}
		
		for _, booking := range bookings {
			var names = strings.Fields(booking) // input: "Nicole Smith" output: ["Nicole", "Smith"]
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}
}