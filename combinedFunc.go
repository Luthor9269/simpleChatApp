package main

import (
	"fmt"

	"github.com/google/uuid"
)

// creates a message of type Message
// function returns that message
func createMessage(sender string, message string, id [16]byte) *Message {
	newMessage := &Message{
		id:      id,
		sender:  sender,
		message: message,
	}
	return newMessage
}

// takes in these parameters to contruct a message
// and puts it into a Node
// Returns a Node
func createNode(sender string, message string, id [16]byte) (*Node, error) {
	newMessage := createMessage(sender, message, id)
	newNode := &Node{
		message: newMessage,
		next:    nil,
	}
	return newNode, nil
}

func addtoListAndStack(sender string, message string, ml *MessageList, ms *MessageStack) error {
	defer wg.Done()
	mutex.Lock()
	newId := uuid.New()
	ml.addMessageToList(sender, message, newId)
	ms.addMessageToStack(sender, message, newId)
	mutex.Unlock()
	fmt.Println("Message added Succesfully")
	return nil
}

// function to undo the last message from the chat
// have to remove from both stack and list

// ///This function takes the top from the STACK and gets the ID
// Using this ID it iterates through the linkedList and then removes that element
func unsendMessage(ms *MessageStack, ml *MessageList) error {
	// var messageId [16]byte
	poppedNode, err := ms.popFromStack()
	if err != nil {
		fmt.Println(err)
	}
	messageId := poppedNode.message.id
	ml.removeNodeFromList(messageId)
	fmt.Println("Message removed succesfully")
	return nil
}
