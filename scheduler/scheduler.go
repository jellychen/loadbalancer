package scheduler

import (
	"errors"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type Scheduler interface {
	Next() string
}

type scheduler struct {
	next <-chan string
}

func NewScheduler(nodes []string) (*scheduler, error) {
	l := len(nodes)
	if l == 0 {
		return nil, errors.New("at leat one node needed!")
	}
	s := new(scheduler)
	i := 0
	c := make(chan string)
	go func() {
		for {
			logger.Debugf("next node pos: %d", i)
			c <- nodes[i]
			i += 1
			if i == l {
				i = 0
			}
		}
	}()
	s.next = c
	return s, nil
}

func (s *scheduler) Next() string {
	return <-s.next
}
