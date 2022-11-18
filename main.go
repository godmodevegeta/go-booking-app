package main

import "fmt"

func main() {

	var name = "Middle-Earth"
	fmt.Printf("Hey! Welcome to the %v!\n", name)
	fmt.Println("The journey is about to start!")

	// fmt.Println(name)
	const totalSeats uint = 60
	var remainingSeats int = 60

	fmt.Println("Total available seats:", totalSeats)
	fmt.Printf("But there are only %v seats available to be booked!\n", remainingSeats)

	// var bookings [20]string
	// bookings[0] = userName

	// var firstName string
	// var lastName string
	// var userTicket int

	// fmt.Println("Please enter first name:")
	// fmt.Scan(&firstName)
	// fmt.Println("Please enter last name:")
	// fmt.Scan(&lastName)

	// fmt.Printf("User %v %v has booked %v tickets.\n", firstName, lastName, userTicket)
	// fmt.Printf("Also userName is %T and userTicket is %T\n", userName, userTicket)
	sum := 0
	for i := 0; i <= 5; i++ {
		sum += i
	}
	fmt.Printf("The sum till 5 is %v.\n", sum)

}
