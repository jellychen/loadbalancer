package main

import (
	"flag"
	"github.com/bborbe/loadbalancer/server"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	log.Print("loadbalancer started")
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
		log.Print("create server failed, %v", err)
		return
	}
	{
		err := srv.Start()
		if err != nil {
			log.Print("start server failed, %v", err)
			return
		}
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	for {
		sig := <-ch
		log.Printf("receive signal %v", sig)

		err := srv.Stop()
		if err != nil {
			log.Print("stop server failed, %v", err)
			return
		}
		log.Print("loadbalancer finished")
		return
	}
}

func splitNodes(content string) []string {
	return strings.Split(content, ",")
}
