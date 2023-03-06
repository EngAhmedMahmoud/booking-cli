package main

import (
	"fmt"
	"time"
)

func sendTicketByEmail(user Participant,remainingTickets uint64) {
	time.Sleep(time.Second*10)
	fmt.Println("#########################################")
	fmt.Println("Hello! ",user.Name)
	fmt.Println("We have sent you ",user.Tickets)
	fmt.Println("please check you email address ",user.Email)
	fmt.Println("The remaining tickets: ",remainingTickets)
	fmt.Println("#########################################")
}