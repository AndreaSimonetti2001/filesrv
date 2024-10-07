package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {

	server, _ := net.Listen("tcp", "0.0.0.0:1919")
	fmt.Printf("LISTENING AT %v\n", server.Addr().String())

	stream, _ := server.Accept()
	fmt.Printf("ACCEPTED CONNECTION FROM %v\n", stream.RemoteAddr())

	br := bufio.NewReader(stream)
	filenameBytes, _, _ := br.ReadLine()
	filename := strings.TrimSpace(string(filenameBytes))
	fmt.Printf("RECEIVED THE FILENAME %v\n", filename)

	file, _ := os.Create(filename)
	fmt.Printf("FILE %v CREATED\n", filename)

	stream.Write([]byte("OK\n"))
	fmt.Printf("CONFIRMATION SENT\n")

	bytes, _ := io.Copy(file, stream)
	file.Close()
	fmt.Printf("RECEIVED %v BYTES\n", bytes)
	stream.Close()
	fmt.Printf("DONE\n")

}
