package main

import (
	"log"
	"net"
	"time"

	server "github.com/EwRvp7LV7/45586694crypto/0server"
	"github.com/EwRvp7LV7/45586694crypto/logger"
)

const (
	PORT = "2121"
)
//tsl
func run() (err error) {

	serv, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		return err
	}
	logger.Println("TCP server is UP @ localhost: " + PORT)

	defer serv.Close()

	for {
		connection, err := serv.Accept()
		connection.SetDeadline(time.Now().Add(time.Minute * 2))
		if err != nil {
			logger.Println("Client Connection failed")
			continue
		}
		
		go server.HandleServer(connection)
	}

	// return nil
}

func main() {

	// flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}

}
