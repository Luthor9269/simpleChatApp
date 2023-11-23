package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

func (ml *MessageList) addMessageToList(sender string, message string, id [16]byte) error {
	// adding to stack
	newNode, err := createNode(sender, message, id)
	if err != nil {
		return err
	} else {
		if ml.head == nil {
			ml.head = newNode
		} else {
			tempNode := ml.head
			// ranging through the nexts until we reach the end
			for tempNode.next != nil {
				// fmt.Println("Hello", newNode.message)
				tempNode = tempNode.next
			}
			// once we reach the end of the linkedList
			// insert our newNode into next
			tempNode.next = newNode
		}
		ml.size++
		return nil
	}
}

// function to see all the messages
func (ml *MessageList) seeMessages() error {
	mutex.Lock()
	defer mutex.Unlock()
	// iterate through the linkedList and print the messages
	currentNode := ml.head
	if currentNode == nil {
		panic(errors.New("there are no Messages in this chat, Unable to see messages"))
	}
	// as long as current node is not nil
	// iterate through and print
	for currentNode != nil {
		fmt.Printf("%v: %v\n", currentNode.message.sender, currentNode.message.message)
		currentNode = currentNode.next
	}
	return nil
}

// Removing the node from both stack and list
func (ml *MessageList) removeNodeFromList(id [16]byte) error {
	if ml.head == nil {
		fmt.Println("The list is empty, Cant remove given id.")
		return errors.New("list is empty")
	}

	// if the id to be removed is at the first element
	if bytes.Equal(ml.head.message.id[:], id[:]) {
		ml.head = ml.head.next
		ml.size--
		return nil
	}
	// Find the node with the specified ID
	// Going through the list keeping track of a previous and a current
	var previousNode *Node
	currentNode := ml.head
	for currentNode != nil && !bytes.Equal(currentNode.message.id[:], id[:]) {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	// Check if the node with the given ID was found
	if currentNode == nil {
		fmt.Println("Message with the ID was not found")
		return nil
	}
	// Remove the node from the linked list
	previousNode.next = currentNode.next
	ml.size--
	return nil
}

// searching functionalities

func (ml *MessageList) searchingByName(senderName string) error {
	senderName = strings.ToLower(senderName)
	// iterate through the linkedList and print the messages
	currentNode := ml.head
	if currentNode == nil {
		panic(errors.New("there are no Messages in this chat, Unable to see messages"))
	}
	var hasUser = false
	// as long as current node is not nil
	// iterate through and print
	for currentNode != nil {
		currSender := strings.ToLower(currentNode.message.sender)
		if currSender == senderName {
			hasUser = true
			fmt.Printf("%v: %v\n", currentNode.message.sender, currentNode.message.message)
		}
		currentNode = currentNode.next
	}
	if !hasUser {
		err := errors.New("no such user exists. Please try again")
		return err
	}
	return nil
}

// searching for specific message using a keyword or phrase?
func (ml *MessageList) searchByWord(word string) error {
	var hasWord bool
	// finding each message to compare in
	// iterate through the linkedList and print the messages
	currentNode := ml.head
	if currentNode == nil {
		panic(errors.New("there are no Messages in this chat, Unable to search for word in messages"))
	}
	// as long as current node is not nil
	// iterate through and print
	for currentNode != nil {
		// implement logic
		wordFound := checkForString(currentNode.message.message, word)
		if wordFound {
			fmt.Printf("%v: %v\n", currentNode.message.sender, currentNode.message.message)
			hasWord = true
			// also push this into a list
		}
		currentNode = currentNode.next
	}
	if !hasWord {
		err := errors.New("this word doesnt exist in the chat. Please try again")
		return err
	}
	return nil
}

// helper function
func checkForString(message string, word string) bool {
	// need to change the strings to lower case
	message = strings.ToLower(message)
	word = strings.ToLower(word)
	// creating lens of the two strings
	lenMessage := len(message)
	lenWord := len(word)

	// Loop through to start comparing
	for i := 0; i <= lenMessage-lenWord; i++ {
		//a boolean to check if the words are currently matching
		//this will be true until there is a break in the matching sequence
		matching := true
		for k := 0; k < lenWord; k++ {
			//check for matching words from the starting point of i
			// keep looping through the substring until it dosnt match
			// when it doesnt match immediately break out of the for loop and go to the next iteration of I
			if message[i+k] != word[k] {
				matching = false
				break
			}
		}
		// Check if there is a match in the string
		if matching {
			return true
		}
	}
	return false
}
