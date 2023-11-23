package main

import (
	"errors"
)

// function to add to stack
func (ms *MessageStack) addMessageToStack(sender string, message string, id [16]byte) error {
	newNode, err := createNode(sender, message, id)
	//Add to the top of the stack
	if err != nil {
		return err
	} else {
		if ms.top == nil {
			ms.top = newNode
		} else {
			newNode.next = ms.top
			ms.top = newNode
		}
		ms.size++
		return nil
	}
}

func (ms *MessageStack) popFromStack() (*Node, error) {

	var poppedNode *Node
	if ms.top == nil {
		panic(errors.New("the chat is empty, cannot undo last message"))
	} else {
		poppedNode = ms.top
		ms.top = ms.top.next
	}
	ms.size--
	return poppedNode, nil
}
