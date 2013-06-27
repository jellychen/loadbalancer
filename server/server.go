package server

import (
	"github.com/bborbe/loadbalancer/connectionhandler"
	"errors"
	"log"
	"net"
	"github.com/bborbe/loadbalancer/scheduler"
)

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
					log.Print(err)
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
