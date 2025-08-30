package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	counter := 0
	randomNumber := rand.IntN(100)
	for x := 0; x < 10; x++ {
		fmt.Print("Enter a number ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess == randomNumber {
			fmt.Println("You win!")
			os.Exit(0)
		} else if guess < randomNumber {
			fmt.Println("Guess too low")
		} else if guess > randomNumber {
			fmt.Println("Guess to high")
		}
		counter = counter + 1

	}
	fmt.Println("sorry you didnt guess my number. It was ", randomNumber)
}
