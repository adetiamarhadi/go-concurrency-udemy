package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		fmt.Println(">>> waiting data from ping")

		s := <-ping

		fmt.Println(">>> data from ping received")

		fmt.Println(">>> sending data to pong")

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))

		fmt.Println(">>> send data to pong done")
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		fmt.Print("-> ")

		var userInput string
		fmt.Scanln(&userInput)
		if strings.EqualFold("q", userInput) {
			break
		}

		fmt.Println("/// userInput received")

		fmt.Println("/// sending userInput to ping")

		ping <- userInput

		fmt.Println("/// ping received the data")

		fmt.Println("/// waiting data from pong")

		response := <-pong
		fmt.Println("Response:", response)
	}

	fmt.Println("All done. Closing channels.")
	close(ping)
	close(pong)
}
