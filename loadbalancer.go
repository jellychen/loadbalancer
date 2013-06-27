package main

import (
	"github.com/bborbe/loadbalancer/server"
	"log"
	"flag"
	"strings"
	"strconv"
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
	srv, err := server.NewServer(":" + strconv.Itoa(port), nodes)
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
	srv.Wait()
	{
		err := srv.Stop()
		if err != nil {
			log.Print("stop server failed, %v", err)
			return
		}
	}
	log.Print("loadbalancer finished")
}

func splitNodes(content string) []string {
	return strings.Split(content, ",")
}
