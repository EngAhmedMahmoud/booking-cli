package main

import (
	"fmt"
	"os"
	"strconv"
)
func main() {
	eventName := os.Getenv("EVENT_NAME");
	greeting(eventName)
	eventTiclets := os.Getenv("EVENT_TICKETS");
	var remainingTickets,_  = strconv.ParseUint(eventTiclets,10,64);

	for remainingTickets > 0{
		user,err :=readInputs(remainingTickets);
		if err != nil {
			fmt.Println(err);
		}else{
			remainingTickets = remainingTickets - user.Tickets;
			go sendTicketByEmail(user,remainingTickets)
		}
	}
}