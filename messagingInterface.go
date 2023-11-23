package main

import (
	"bufio"
	"fmt"
	"os"
)

func messageingInterface(ml *MessageList, ms *MessageStack) {
	// handling panic incase stack/list is empty and can't pop anymore
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()
	for {
		var userChoice int
		var senderName string
		var userMessage string
		var searchWord string
		fmt.Println("")
		fmt.Println("Please choose what you would like to do")
		fmt.Println("1: See Messages")
		fmt.Println("2: Send Message")
		// fmt.Println("3: Unsend last Message")
		fmt.Println("3: Unsend last message")
		fmt.Println("4: Search messages of a single user")
		fmt.Println("5: Search through the chat for a word")
		fmt.Println("6: Exit")
		fmt.Scanln(&userChoice)
		switch userChoice {
		case 1:
			err := ml.seeMessages()
			if err != nil {
				fmt.Print("Error in See Messages:", err)
			} else {
				fmt.Println("Printed messages succesfully")
			}

		case 2:
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Println("Please key in the sender Name")
			if scanner.Scan() {
				senderName = scanner.Text()
			} else {
				fmt.Println("There was an error reading your message:", scanner.Err())
			}
			fmt.Println("Please key in what you would like to say")
			// create an instance of scanner
			// check if the scan was succesful
			if scanner.Scan() {
				userMessage = scanner.Text()
			} else {
				fmt.Println("There was an error reading your message:", scanner.Err())
			}
			wg.Add(1)
			go func() {
				addtoListAndStack(senderName, userMessage, ml, ms)
			}()
		case 3:
			unsendMessage(ms, ml)
		case 4:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Please key in the sender name")
			if scanner.Scan() {
				senderName = scanner.Text()
			} else {
				fmt.Println("There was an error reading your message:", scanner.Err())
			}
			err := ml.searchingByName(senderName)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 5:
			fmt.Println("Which word would you like to search for?")
			fmt.Scanln(&searchWord)
			err := ml.searchByWord(searchWord)
			if err != nil {
				fmt.Println("There was an error searching for the word:", err)
			}
		case 6:
			return
		default:
			fmt.Println("Please key in a number from 1-6")
		}
	}
}
