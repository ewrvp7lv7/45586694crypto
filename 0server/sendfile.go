package server

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/EwRvp7LV7/45586694crypto/logger"
)

func sendFile(conn net.Conn, name string) {

	inputFile, err := os.Open(ROOT + "/" + name)
	if err != nil {
		logger.Println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}
	defer inputFile.Close()

	stats, _ := inputFile.Stat()

	conn.Write([]byte(fmt.Sprintf("download %s %d\n", name, stats.Size())))

	io.Copy(conn, inputFile)

	logger.Println("File " + name + " Send successfully")
}
