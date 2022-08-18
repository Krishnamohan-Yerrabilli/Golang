package main

import "fmt"

func main() {
	const Tickets_count = 50
	var Tickets_available = 50
	var conferenceTickets = "Go_Conference"

	fmt.Println("Welcome to our", conferenceTickets, "tickes we are so happy this site is for you")
	fmt.Println("Overall We had ", Tickets_count, "for the concert, and we have", Tickets_available, "tickets are still available")
	fmt.Println("Get your tickets for the ", conferenceTickets)
}
