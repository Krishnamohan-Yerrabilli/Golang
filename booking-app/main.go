package main

import "fmt"

func main() {
	const Tickets_count = 50
	Tickets_available := 50
	conferenceName := "Go_conference"
	var booking_list [50]string

	fmt.Printf("\nTickets_count is a type of %T, and conferenceName is a type of %T\n", Tickets_count, conferenceName)

	fmt.Printf("\nWelcome to  our %v we are so happy this site is for you\n", conferenceName)
	fmt.Printf("Overall We have %v tickets are opened for the concert, and we have %v tickets are still available\n", Tickets_count, Tickets_available)
	fmt.Println("\nGet your tickets for the", conferenceName)

	var first_name string

	fmt.Println("Enter your first_name: ")
	fmt.Scan(&first_name)

	var last_name string
	fmt.Println("Enter your last_name: ")
	fmt.Scan(&last_name)

	var email_name string
	fmt.Println("Enter your email: ")
	fmt.Scan(&email_name)

	var tickets_booked int
	fmt.Println("Please enter how many tickets you want to book: ")
	fmt.Scan(&tickets_booked)

	fmt.Printf("Congratulations %v %v you booked %v tickets, you will get a mail at %v Thankyou", first_name, last_name, tickets_booked, email_name)

	Tickets_available = Tickets_available - tickets_booked
	booking_list[0] = first_name + " " + last_name

	fmt.Printf("\nWhole array list: %v\n", booking_list)
	fmt.Printf("The first value: %v\n", booking_list[0])
	fmt.Printf("The Type : %T\n", booking_list)
	fmt.Printf("The length of the array: %v\n", len(booking_list))

	fmt.Printf("\nNow only %v tickets has been available for %v, have a great day!", Tickets_available, conferenceName)

}
