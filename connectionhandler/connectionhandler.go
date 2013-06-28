package connectionhandler

import (
	"net"
	"log"
	"io"
	"github.com/bborbe/loadbalancer/scheduler"
)

type ConnectionHandler interface {
	HandleConnection(conn net.Conn)
}

type connectionhandler struct {
	scheduler scheduler.Scheduler
}

func NewConnectionHandler( scheduler scheduler.Scheduler ) *connectionhandler {
	c := new(connectionhandler)
	c.scheduler = scheduler
	return c
}

func (c *connectionhandler) HandleConnection(clientConn net.Conn) {
	log.Print("process connection started")
	defer clientConn.Close()

	serverConn, err := net.Dial("tcp", c.scheduler.Next())
	if err != nil {
		log.Print(err)
		return
	}
	defer serverConn.Close()

	done := make(chan bool, 2)
	go copyChan("client->server",clientConn, serverConn, done)
	go copyChan("server->client",serverConn, clientConn, done)
	<-done
	<-done
	log.Print("process connection finished")
}

func copyChan(name string, input net.Conn, output net.Conn, done chan bool) {
	log.Printf("%s copyChan started", name)
	io.Copy(input, output)
	log.Printf("%s copyChan finished", name)
	done <- true
}

