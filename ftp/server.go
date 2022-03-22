package ftp

import (
	"fmt"
	"log"
	"net"
	"path/filepath"
)

type Addr struct {
	Network string // name of the network (for example, "tcp", "udp")
	URI string // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
	Port int
}

type Server struct {
	LocalAddr Addr
	RemoteAddr Addr
}


func (c *Server) Listen (address Addr)  {

	// TODO make config
	listener, err := net.Listen("tcp", address.URI)
	if err != nil {
		fmt.Println("ERROR::", err)
		return
	}

	fmt.Println("Server start listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ERROR::",err)
		}
		fmt.Println("New connection:", conn.RemoteAddr().Network(), conn.RemoteAddr().String())
		go c.handleConnection(conn)

	}
}

func (c *Server) handleConnection(conn net.Conn)  {
	absPath, err := filepath.Abs("D:/Humble App")
	if err != nil {
		log.Fatal(err)
	}
	Route(NewConn(conn, absPath))

	// Close connection
	err = conn.Close()
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}