package main

import (
	"errors"
	"fmt"
)

var ErrNotAllowed = errors.New("Hata!")

func entryInParty(age int) (string, error) {
	if age < 18 {
		return "", ErrNotAllowed
	}

	return "TICKET", nil
}

func main() {
	ticket, err := entryInParty(17)

	if err != nil {
		fmt.Println("Hata:", err)
	}

	ticket, err = entryInParty(19)

	if err == ErrNotAllowed {
		fmt.Println("ErrNotAllowed")
	}

	fmt.Println("Ticket:", ticket)
}
