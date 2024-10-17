package main

import (
	"bufio"
	"filesrv/logger"
	"filesrv/try"
	"io"
	"net"
	"os"
	"strings"
)

func main() {

	server, e := net.Listen("tcp", "0.0.0.0:1919")
	try.Catch(e)
	logger.Info("Listening at %v.", server.Addr().String())

	stream, e := server.Accept()
	try.Catch(e)
	logger.Info("Accepted connection from %v.", stream.RemoteAddr())

	br := bufio.NewReader(stream)
	filenameBytes, _, e := br.ReadLine()
	try.Catch(e)
	filename := strings.TrimSpace(string(filenameBytes))
	logger.Info("Received the filename %v.", filename)

	file, e := os.Create(filename)
	try.Catch(e)
	logger.Info("File %v created.", filename)

	_, e = stream.Write([]byte("OK\n"))
	try.Catch(e)
	logger.Info("Confirmation sent.")

	bytes, e := io.Copy(file, stream)
	try.Catch(e)
	e = file.Close()
	try.Catch(e)
	logger.Info("File closed, received %v bytes.", bytes)

	e = stream.Close()
	try.Catch(e)
	logger.Info("Stream closed.")

}
