package main

import (
	"fmt"
	"net"

	client "github.com/EwRvp7LV7/45586694crypto/1client"
)

const (
	PORT = "2121"
	HOST = "localhost"
)

func run() (err error) {

	connect, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		return err
	}
	fmt.Println("TCP server is Connected @ ", HOST, ":", PORT)

	defer connect.Close()

	if err := client.AuthenticateClient(connect); err != nil {
		return err
	}

	client.HandleClient(connect)

	return nil
}

func main() {

	// flag.Parse()

	if err := run(); err != nil {
		fmt.Print(err)
	}

}
