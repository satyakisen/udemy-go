package main

import "fmt"

// SeatingArrangement decides the seating arrangement in a compartment
func SeatingArrangement() {
	var diff []int
	for i := 12; i > 0; i-- {
		if i%2 != 0 {
			diff = append(diff, i)
		}
	}
	// fmt.Print(len(diff))
	var seat int
	var oppSeat int
	fmt.Scanf("%d", &seat)
	if seat%12 == 0 {
		oppSeat = seat - diff[0]
	} else {
		if seat%12 <= 6 {
			oppSeat = seat + diff[seat%12-1]
		} else {
			oppSeat = seat - diff[12-seat%12]
		}
	}

	fmt.Println(oppSeat)
}

func main() {
	SeatingArrangement()
}
