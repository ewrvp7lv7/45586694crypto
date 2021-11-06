package server

import (
	"bufio"
	"errors"
	"net"

	"github.com/EwRvp7LV7/45586694crypto/logger"
)

func AuthenticateClient(conn net.Conn) error {

	conn.Write([]byte("Server:Connection Established"))
	creds := GetCred()

	reader := bufio.NewScanner(conn)
	//validate user
	reader.Scan()
	uname := reader.Text()
	reader.Scan()
	passwd := reader.Text()

	for _, cred := range *creds {

		if cred.Username == uname && cred.Password == passwd {
			logger.Println("Server:Client " + uname + " Validated")
			conn.Write([]byte("1"))
			return nil
		}
	}

	conn.Write([]byte("0"))
	return errors.New("invalid credentials: " + uname)

}
