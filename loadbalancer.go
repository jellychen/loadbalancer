package main

import (
	"flag"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bborbe/loadbalancer/server"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

func main() {
	defer logger.Close()
	logger.Debug("loadbalancer started")
	portPtr := flag.Int("port", 8081, "ListenPort")
	nodesPtr := flag.String("nodes", "", "NodeList")
	flag.Parse()
	port := *portPtr
	nodes := splitNodes(*nodesPtr)
	if port <= 0x0 || port >= 0xFFFF {
		flag.Usage()
		return
	}
	srv, err := server.NewServer(":"+strconv.Itoa(port), nodes)
	if err != nil {
		logger.Debugf("create server failed, %v", err)
		return
	}
	{
		err := srv.Start()
		if err != nil {
			logger.Debugf("start server failed, %v", err)
			return
		}
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	for {
		sig := <-ch
		logger.Debugf("receive signal %v", sig)

		err := srv.Stop()
		if err != nil {
			logger.Debugf("stop server failed, %v", err)
			return
		}
		logger.Debug("loadbalancer finished")
		return
	}
}

func splitNodes(content string) []string {
	return strings.Split(content, ",")
}
