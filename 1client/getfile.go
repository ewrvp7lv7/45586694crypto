package client

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	"log"
	"net"
	"os"
)

func getFile(conn net.Conn, myFPass string) {
	buf := new(bytes.Buffer)

	for {
		io.CopyN(buf, conn, 1)
		if buf.Bytes()[len(buf.Bytes())-1] == '\n' {
			break
		}
	}

	commandArr := strings.Fields(strings.Trim(buf.String(), "\n"))
	buf.Reset()

	name := commandArr[1]

	fileSize, err := strconv.ParseInt(commandArr[2], 10, 64)
	if err != nil || fileSize == -1 {
		log.Println("file size error", err)
		conn.Write([]byte("file size error"))
		return
	}

	io.Copy(io.Writer(buf), io.LimitReader(conn, fileSize))

	arrDec, err := CBCDecrypter(myFPass, buf.Bytes())
	if err != nil {
		log.Println(err)
		return
	}

	outputFile, err := os.Create(ROOT + "/" + name)
	if err != nil {
		log.Println(err)
	}
	io.Copy(outputFile, bytes.NewReader(arrDec)) //Для кратности 16 в исходный файл добавляем
	defer outputFile.Close()

	// conn.Write([]byte("File Downloaded successfully"))
	log.Println("File Downloaded successfully")

	checkFileMD5Hash(ROOT + "/" + name)
}
