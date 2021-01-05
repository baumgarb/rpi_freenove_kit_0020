package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
	}
	defer rpio.Close()
	
	pin := rpio.Pin(17)
	pin.Output()
	
	var input string
	fmt.Printf("What do you wanna do next?\n\t0. Exit\n\t1. Turn on LED\n\t2. Turn off LED\n")
	keepGoing := true
	for keepGoing {
		fmt.Printf("\tChoose option number: ")
		fmt.Scanf("%s", &input)
		
		switch(input) {
			case "0":
				keepGoing = false
			case "1":
				pin.High()
			case "2":
				pin.Low()
			default:
				continue
		}
	}
}
