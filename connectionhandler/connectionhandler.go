package connectionhandler

import (
	"io"
	"net"

	"github.com/bborbe/loadbalancer/scheduler"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type ConnectionHandler interface {
	HandleConnection(conn net.Conn)
}

type connectionhandler struct {
	scheduler scheduler.Scheduler
}

func NewConnectionHandler(scheduler scheduler.Scheduler) *connectionhandler {
	c := new(connectionhandler)
	c.scheduler = scheduler
	return c
}

func (c *connectionhandler) HandleConnection(clientConn net.Conn) {
	logger.Debug("process connection started")
	defer clientConn.Close()

	serverConn, err := net.Dial("tcp", c.scheduler.Next())
	if err != nil {
		logger.Debug(err)
		return
	}
	defer serverConn.Close()

	done := make(chan bool, 2)
	go copyChan("client->server", clientConn, serverConn, done)
	go copyChan("server->client", serverConn, clientConn, done)
	<-done
	<-done
	logger.Debug("process connection finished")
}

func copyChan(name string, input net.Conn, output net.Conn, done chan bool) {
	logger.Debugf("%s copyChan started", name)
	io.Copy(input, output)
	logger.Debugf("%s copyChan finished", name)
	done <- true
}
