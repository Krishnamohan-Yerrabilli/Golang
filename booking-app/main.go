package main

import "fmt"

func main() {
	const Tickets_count = 50
	Tickets_availabe := 45
	conferenceTickets := "Go_conference"

	fmt.Printf("\nTickets_count is a type of %T, and conferenceTickets is a type of %T\n", Tickets_count, conferenceTickets)

	fmt.Printf("\nWelcome to  our %v we are so happy this site is for you\n", conferenceTickets)
	fmt.Printf("Overall We have %v tickets are opened for the concert, and we have %v tickets are still available\n", Tickets_count, Tickets_availabe)
	fmt.Println("\nGet your tickets for the", conferenceTickets)

	var username string

	var user_booked int
	//username of a user

	username = "mohan"
	user_booked = 5

	fmt.Printf("\n%v booked %v tickets for the concert. \n", username, user_booked)

}
