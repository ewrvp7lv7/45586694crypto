package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func sendFile(conn net.Conn, name string, myFPass string) {

	content, err := ioutil.ReadFile(ROOT + "/" + name)
	if err != nil {
		log.Println(err)
		return
	}

	arrEnc, err := CBCEncrypter(myFPass, content)
	if err != nil {
		log.Println("EOF", err) //remove
		return
	}

	conn.Write([]byte(fmt.Sprintf("upload %s %d\n", name, len(arrEnc))))

	io.Copy(conn, bytes.NewReader(arrEnc))

	buffer := make([]byte, BUFFERSIZE)
	n, _ := conn.Read(buffer)
	log.Println(string(buffer[:n]))

	checkFileMD5Hash(ROOT + "/" + name)
}
