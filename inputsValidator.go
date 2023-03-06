package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)
type Participant struct {
	Name string
	Email string
	Tickets uint64
}
	
func readInputs(avaliableTickets uint64) (Participant,error){

	var participant Participant;
	scanner := bufio.NewScanner(os.Stdin);

	fmt.Print("Enter your name:")
	scanner.Scan()
	participant.Name = scanner.Text();
	if emptyInput(participant.Name){
		return Participant{}, fmt.Errorf("name " + EMPTY_INPUT);
	}

	fmt.Print("Enter number of tickets:")
	scanner.Scan()
	tickets,invalidNumber := strconv.ParseUint(scanner.Text(),10,64);
	if invalidNumber != nil ||  tickets > avaliableTickets{
		return Participant{}, fmt.Errorf("tickets " + INVALID_NUMBER);
	}
	participant.Tickets = tickets;

	fmt.Print("Enter your email address:")
	scanner.Scan()
	participant.Email = scanner.Text();
	if !validEMail(participant.Email) {
		return Participant{}, fmt.Errorf(INVALID_EMAIL);
	}
	return participant,nil	
}

func emptyInput(input string)bool{
	if len(input) == 0{
		return true;
	}
	return false;
}

func validEMail(email string) bool{
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
