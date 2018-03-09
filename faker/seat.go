package faker

import (
	"math/rand"
	"strconv"
)

func SeatOfTrain() string {
	n := rand.Intn(5)
	var seat string
	if n%2 == 0 {
		seatNumberNew := strconv.Itoa(Number(1, 15))
		seatLetter := getRandValue([]string{"seat", "train"})
		seat = seatNumberNew + seatLetter
	} else {
		seat = strconv.Itoa(Number(1, 120))
	}
	return seat
}

func SeatOfFlight() string {
	seatNumber := strconv.Itoa(Number(1, 15))
	seatLetter := getRandValue([]string{"seat", "flight"})
	seat := seatNumber + seatLetter
	return seat
}
