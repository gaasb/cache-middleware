package cache_middleware

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
)

var acceptedConn = map[string]net.Conn{}

func NewServer() net.Listener {
	server, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		Close(err)
	}
	return server

}
func Serve(server net.Listener) {
	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(server)
	for {
		Handle(server)

	}
}

func Handle(server net.Listener) {
	clientConn, err := server.Accept()
	logIf(err)
	err = AuthClient(clientConn)
	logIf(err)
}

func Close(err error) {
	fmt.Println("Closed" + err.Error())
	os.Exit(1)
}

func AuthClient(conn net.Conn) error {
	address := conn.RemoteAddr().String()
	if acceptedConn[address] != nil {
		return nil
	} else {
		conn.Close()
		return errors.New(address + " : Address not accepted")
	}
}
func logIf(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
