package ftp

import (
	"fmt"
	"net"
)

// Conn represents a connection to the FTP server
type Conn struct {
	conn     net.Conn
	dataType dataType
	dataPort *dataPort
	rootDir  string
	workDir  string
}

// NewConn returns a new FTP connection
func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}

// EOL returns the line terminator matching the FTP standard for the datatype.
func (c *Conn) EOL() string {
	switch c.dataType {
	case ascii:
		return "\r\n"
	case binary:
		return "\n"
	default:
		return "\n"
	}
}

type dataPort struct {
	h1, h2, h3, h4 int // host
	p1, p2         int // port
}

func dataPortFromHostPort(hostPort string) (*dataPort, error) {
	var dp dataPort
	_, err := fmt.Sscanf(hostPort, "%d,%d,%d,%d,%d,%d",
		&dp.h1, &dp.h2, &dp.h3, &dp.h4, &dp.p1, &dp.p2)
	if err != nil {
		return nil, err
	}
	return &dp, nil
}

//func dataPortToHostPort(hostPort string) (*dataPort, error) {
//	return
//}

func (d *dataPort) toAddress() string {
	if d == nil {
		return ""
	}
	// convert hex port bytes to decimal port
	port := d.p1<<8 + d.p2
	return fmt.Sprintf("%d.%d.%d.%d:%d", d.h1, d.h2, d.h3, d.h4, port)
}