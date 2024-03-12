package helper

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func UserDataRetriever() (string, string, string, uint, string) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var errorMessage string	
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)	

	if (len(firstName) < 2 || len(lastName) < 2){
		errorMessage = "First and Last names must not be less than 2 chars"
	} else if !strings.Contains(email, "@"){
		errorMessage = "Invalid Email"
	}
	return firstName, lastName, email, userTickets, errorMessage
}

func TicketValidator(userTickets uint, remainingTickets uint) bool {
	if userTickets > remainingTickets || remainingTickets == 0{
		return false
	} else {
		return true
	}
}


func SendEmail(wg *sync.WaitGroup, firstName string, lastName string, email string, userTickets uint){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v", userTickets, firstName + " " + lastName)
	fmt.Println("***********************")
	fmt.Printf("Sending ticket: \n%v \nto email address %v\n", ticket, email)
	fmt.Println("***********************")
	wg.Done()
}