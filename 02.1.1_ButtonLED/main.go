package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	ledPinNr := 17
	buttonPinNr := 18

	err := rpio.Open()
	defer rpio.Close()

	if err != nil {
		fmt.Println(err)
	}

	ledPin := rpio.Pin(ledPinNr)
	ledPin.Output()
	buttonPin := rpio.Pin(buttonPinNr)
	buttonPin.Input()
	buttonPin.PullUp()

	for {
		input := buttonPin.Read()
		if input == rpio.Low {
			ledPin.High()
		} else {
			ledPin.Low()
		}

		time.Sleep(time.Second)
		fmt.Println("Input: ", input)
	}
}
