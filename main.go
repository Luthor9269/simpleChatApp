package main

import (
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	// creating and using two data structures
	// an arrayList to save the
	// Initialise an instance of MessageList and MessageStack
	//empty message list
	// Getting a preset Message

	myMessageList := &MessageList{
		head: nil,
		size: 0,
	}
	myMessageStack := &MessageStack{
		top:  nil,
		size: 0,
	}
	wg.Add(7)
	// Adding messages directly into list and stack
	go addtoListAndStack("Mayank", "Hello Everyone", myMessageList, myMessageStack)
	go addtoListAndStack("Emil", "Hello Mayank. Whats up?", myMessageList, myMessageStack)
	go addtoListAndStack("Aesther", "I think its time to get started on the project", myMessageList, myMessageStack)
	go addtoListAndStack("James", "Which project are you refering to?", myMessageList, myMessageStack)
	go addtoListAndStack("Mayank", "The Subscription Playlist", myMessageList, myMessageStack)
	go addtoListAndStack("Emil", "Havent we started on the subscription playlist?", myMessageList, myMessageStack)
	go addtoListAndStack("Aesther", "We have started the brainstorming phase for the subscription playlist", myMessageList, myMessageStack)
	wg.Wait()

	messageingInterface(myMessageList, myMessageStack)
}
