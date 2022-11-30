package main

import (
	"fmt"
	//"strconv"  needed when we were using maps
	"strings"
)

// package level variables
const totalTickets int = 50
var conferenceName = "go conference"
var remainingTickets uint= 50
//var bookings = make([]map[string]string, 0)
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	emailId string
	userTicket uint
}

func main() {
	

	greetingUser()
	
	for{
		//get user input
		firstName, lastName, emailId, userTicket :=  getUserInput()

		// we need to check if user is entering right input or not
		isValidName,isValidEmail,isValidTickets := userInputValidate(firstName , lastName, emailId, userTicket)

		if isValidEmail && isValidName && isValidTickets{
			remainingTickets = remainingTickets - userTicket

			// using maps
			// var userData = make(map[string]string)
			// userData["firstName"] = firstName
			// userData["lastName"] = lastName
			// userData["emailId"] = emailId
			// userData["userTicket"] = strconv.FormatUint(uint64(userTicket),10)

			// using struct
			var userData = userData{
				firstName: firstName,
				lastName: lastName,
				emailId: emailId,
				userTicket: userTicket,
			}

			bookings = append( bookings , userData )
			fmt.Printf("List of booking:- %v \n", bookings)
			

			fmt.Printf("Thank you %v %v for booking tickets, You have booked %v Tickets \n", firstName, lastName, userTicket)
			fmt.Printf("You will recieve confirmation on %v \n",emailId)
			fmt.Printf("We are left with %v Tickets \n",remainingTickets)

			firstNames := printFirstName()
			fmt.Printf("The first name of all booked members are: %v\n", firstNames)

			if remainingTickets==0 {
				fmt.Println("Tickets are sold out this year, see you next year")
				break
			}

		}else{
			if !isValidName{
				fmt.Println("You have entered first name or last name incorrectly")
			}
			if !isValidEmail{
				fmt.Println("You have entered wrong email address")
			}
			if !isValidTickets{
				fmt.Println("You have entered wrong number of tickets")
			}
		}
	}
}

func greetingUser(){
	fmt.Printf("Welcome to this %v Booking Application\n", conferenceName)
	fmt.Println("Get your tickets booked here for " + conferenceName)
}

func printFirstName() []string{
	firstNames := []string{}
	for _,booking := range bookings{
		firstNames = append(firstNames , booking.firstName)
	}
	//fmt.Printf("The first name of all booked members are: %v\n", firstNames)
	return firstNames
}


func getUserInput()(string, string, string, uint){
	var firstName string
	var lastName string
	var userTicket uint
	var emailId string

	fmt.Print("Enter your first name:- ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name:- ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email:- ")
	fmt.Scan(&emailId)

	fmt.Print("Enter number of tickets you want to buy:- ")
	fmt.Scan(&userTicket)

	return firstName, lastName, emailId, userTicket
}

func userInputValidate(firstName string, lastName string, emailId string, userTicket uint)(bool,bool,bool){
	isValidName := len(firstName)>2 && len(lastName)>2
	isValidEmail := strings.Contains(emailId , "@")
	isValidTickets := userTicket>0 && userTicket <= remainingTickets
	return isValidName,isValidEmail,isValidTickets
}