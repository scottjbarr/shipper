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
		panic("Please set SERVER address e.g. SERVER=127.0.0.1:50000")
	}

	conn, err := net.Dial("udp", server)
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

		// unable to write the message, but not much we can do.
		//
		// Saving the message(s) for later is dangerous because what if you
		// have a ton of processes saving messages for later? i.e. many
		// containers running on a host, all backing up memory.
		if err != nil {
			fmt.Printf("ERROR Cannot write to server %v\n", server)
		}
	}
}
