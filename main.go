package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings = make([]map [string]string, 0)


	

func main() {

	greetUsers()



	for {
		
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserinput(firstName, lastName, email, userTickets, remainingTickets)


		if isValidName && isValidEmail && isValidUserTickets { 


			//function to book ticket
	
			bookTicket(userTickets , firstName, lastName , email)

			//function print first names
			firstNames := getFirstNames()
			fmt.Printf("the first names of out bookings are: %v\n", firstNames)


			if remainingTickets == 0{
				//end program

				fmt.Printf("Our %v is totally booked! Come back next year\n", conferenceName)
				break	
			}

		
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Email address is incorrect")
			}

			if !isValidUserTickets {
				fmt.Println("Number of tickets is invalid")
			}
			fmt.Println("Try again!!!")
		}
	}

}



func greetUsers() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}


func getFirstNames( ) []string {

	firstNames:= []string{}
	for _, booking := range bookings{
		firstNames = append(firstNames , booking["firstName"])
	} 
	
	return firstNames

}


func getUserInput() (string, string, string, uint){
	var firstName string
		var lastName string
		var email string
		var userTickets uint
	

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets you need: ")
		fmt.Scan(&userTickets)


		return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for the user
	var userData = make(map [string]string) 
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	strconv.FormatUint(uint64(userTickets), 10)


	bookings = append(bookings, userData)
	fmt.Printf("List of bookings are %v\n", bookings)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice Type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))


	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}