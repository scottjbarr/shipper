package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	server := os.Getenv("SERVER")

	if server == "" {
		panic("Please set SERVER address e.g. SERVER=127.0.0.0:5000")
	}

	// setup the server address
	serverAddr, err := net.ResolveUDPAddr("udp", server)
	if err != nil {
		panic(err)
	}

	// setup the client address to send from
	localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	// build the connection
	conn, err := net.DialUDP("udp", localAddr, serverAddr)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// channel to receive bytes on STDIN
	c := make(chan []byte)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)

		// block until input is received
		for scanner.Scan() {
			// write the bytes to the channel
			c <- scanner.Bytes()
		}
	}()

	// forever loop...
	for {
		// wait for messages on channel, blocking
		b := <-c

		// message received, write it to the connection
		_, err := conn.Write(b)

		// unable to write the message. Not much we can do. We don't want to
		// save the messages for later, because what if you have a ton of
		// processes saving messages for later (i.e. many containers running
		// on a host).
		if err != nil {
			fmt.Printf("ERROR Cannot write to server %v\n", server)
		}
	}
}
