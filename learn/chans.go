package learn

func send(c chan string, s string) {
	c <- s
}

func SimpleChannels() string {
	messages := make(chan string)
	go send(messages, "ping")
	// inline func as well: go func() { messages <- "ping" }()
	msg := <-messages
	return msg
}
