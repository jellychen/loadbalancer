package server

import (
	"errors"
	"net"

	"github.com/bborbe/loadbalancer/connectionhandler"
	"github.com/bborbe/loadbalancer/scheduler"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type Server interface {
	Start() error
	Stop() error
	Wait()
}

type server struct {
	addr              string
	listener          net.Listener
	connectionHandler connectionhandler.ConnectionHandler
	done              chan bool
}

func NewServer(addr string, nodes []string) (*server, error) {
	s := new(server)
	s.addr = addr
	scheduler, err := scheduler.NewScheduler(nodes)
	if err != nil {
		return nil, err
	}
	s.connectionHandler = connectionhandler.NewConnectionHandler(scheduler)
	return s, nil
}

func (s *server) Start() error {
	if s.listener != nil {
		return errors.New("already started")
	}
	l, e := net.Listen("tcp", s.addr)
	if e != nil {
		return e
	}
	s.listener = l
	go func() {
		for {
			if s.listener != nil {
				conn, err := l.Accept()
				if err != nil {
					logger.Debug(err)
					continue
				}
				go s.connectionHandler.HandleConnection(conn)
			}
		}
		s.done <- true
	}()
	return nil
}

func (s *server) Wait() {
	<-s.done
}

func (s *server) Stop() error {
	if s.listener == nil {
		return errors.New("already stopped")
	}
	l := s.listener
	s.listener = nil
	return l.Close()
}
