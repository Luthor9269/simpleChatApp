package main

// individual messages
// this is a node
type Message struct {
	id      [16]byte
	sender  string
	message string
}

type Node struct {
	message *Message
	next    *Node
}

// List of messages
// passing in the type which is of &message/ pointer(MESSAGES)
type MessageList struct {
	head *Node
	size int
}

// a stack struct
// going to hold all messages
// should I just implement 5 or have the whole list of messages held inside?
type MessageStack struct {
	top  *Node
	size int
}
